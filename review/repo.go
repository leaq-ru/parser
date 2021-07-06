package review

import (
	"context"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CompanyID primitive.ObjectID `bson:"c,omitempty"`
	UserID    primitive.ObjectID `bson:"u,omitempty"`
	Text      string             `bson:"t,omitempty"`
	Status    status             `bson:"s,omitempty"`
}

type status uint8

const (
	MODERATION status = iota
	OK
)

func Create(ctx context.Context, companyID, userID primitive.ObjectID, text string) error {
	_, err := mongo.Reviews.InsertOne(ctx, Review{
		CompanyID: companyID,
		UserID:    userID,
		Text:      text,
		Status:    MODERATION,
	})
	return err
}

func Get(ctx context.Context, companyID primitive.ObjectID, skip, limit int64) ([]Review, error) {
	cur, err := mongo.Reviews.Find(ctx, Review{
		CompanyID: companyID,
		Status:    OK,
	}, options.Find().
		SetSort(bson.M{
			"_id": -1,
		}).
		SetSkip(skip).
		SetLimit(limit))
	if err != nil {
		return nil, err
	}

	var res []Review
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Delete(ctx context.Context, reviewID, userID primitive.ObjectID) error {
	_, err := mongo.Reviews.DeleteOne(ctx, Review{
		ID:     reviewID,
		UserID: userID,
	})
	return err
}

func DeleteAll(ctx context.Context, userID primitive.ObjectID) error {
	_, err := mongo.Reviews.DeleteMany(ctx, Review{
		UserID: userID,
	})
	return err
}

func CountModeration(ctx context.Context, userID primitive.ObjectID) (int64, error) {
	return mongo.Reviews.CountDocuments(ctx, Review{
		UserID: userID,
		Status: MODERATION,
	})
}

func SetOK(ctx context.Context, reviewID primitive.ObjectID) error {
	_, err := mongo.Reviews.UpdateOne(ctx, Review{
		ID: reviewID,
	}, bson.M{
		"$set": Review{
			Status: OK,
		},
	})
	return err
}
