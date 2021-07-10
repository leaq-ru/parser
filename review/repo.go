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
	Positive  bool               `bson:"p,omitempty"`
	Status    status             `bson:"s,omitempty"`
}

type status uint8

const (
	moderation status = iota
	ok
)

func Create(
	ctx context.Context,
	companyID,
	userID primitive.ObjectID,
	text string,
	positive bool,
) (
	Review,
	error,
) {
	r := Review{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		UserID:    userID,
		Text:      text,
		Positive:  positive,
		Status:    moderation,
	}
	_, err := mongo.Reviews.InsertOne(ctx, r)
	if err != nil {
		return Review{}, err
	}

	return r, nil
}

func Get(ctx context.Context, companyID primitive.ObjectID, skip, limit int64) ([]Review, error) {
	cur, err := mongo.Reviews.Find(ctx, Review{
		CompanyID: companyID,
		Status:    ok,
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

func Delete(ctx context.Context, reviewID, userID primitive.ObjectID, forceUserID bool) error {
	q := Review{
		ID: reviewID,
	}
	if forceUserID {
		q.UserID = userID
	}

	_, err := mongo.Reviews.DeleteOne(ctx, q)
	return err
}

func DeleteAll(ctx context.Context, userID primitive.ObjectID) error {
	_, err := mongo.Reviews.DeleteMany(ctx, Review{
		UserID: userID,
	})
	return err
}

func CountModeration(ctx context.Context, userID primitive.ObjectID) (int64, error) {
	return mongo.Reviews.CountDocuments(ctx, bson.M{
		"u": userID,
		"s": nil,
	})
}

func SetOK(ctx context.Context, reviewID primitive.ObjectID) error {
	_, err := mongo.Reviews.UpdateOne(ctx, Review{
		ID: reviewID,
	}, bson.M{
		"$set": Review{
			Status: ok,
		},
	})
	return err
}
