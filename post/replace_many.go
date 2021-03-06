package post

import (
	"context"
	"errors"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/vk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ReplaceMany(
	ctx context.Context,
	companyID primitive.ObjectID,
	vkGroupID int,
	replaceIfNewLessThanOld bool,
) (
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if companyID.IsZero() {
		err = errors.New("empty companyID")
		logger.Log.Error().Err(err).Send()
		return
	}
	if vkGroupID == 0 {
		err = errors.New("empty vkGroupID")
		logger.Log.Error().Err(err).Send()
		return
	}

	wall, err := vk.UserApi.WallGet(api.Params{
		"owner_id": -vkGroupID,
		"count":    20,
		"filter":   "owner",
	})
	if err != nil {
		logger.Log.Error().Int("vkGroupID", vkGroupID).Err(err).Send()
		return
	}

	var newDocs []Post
	for _, item := range wall.Items {
		if item.PostType != "post" {
			continue
		}

		doc := Post{
			CompanyID: companyID,
			Date:      time.Unix(int64(item.Date), 0),
			Text:      item.Text,
		}

		for _, attach := range item.Attachments {
			if attach.Type != "photo" {
				continue
			}

			ph := Photo{}
			for _, size := range attach.Photo.Sizes {
				if ph.URLm != "" && ph.URLr != "" {
					break
				}

				switch size.Type {
				case "m":
					ph.URLm = size.URL
				case "r":
					ph.URLr = size.URL
				}
			}
			doc.Photos = append(doc.Photos, ph)
		}

		if doc.Text == "" && len(doc.Photos) == 0 {
			continue
		}

		newDocs = append(newDocs, doc)
	}

	if !replaceIfNewLessThanOld {
		count, e := mongo.Posts.CountDocuments(ctx, Post{
			CompanyID: companyID,
		}, options.Count().SetLimit(20))
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}

		if int(count) > len(newDocs) {
			logger.Log.Debug().Msg("posts count in mongo > newDocs to replace. Skip inserting")
			return nil
		}
	}

	sess, err := mongo.Client.StartSession()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	defer sess.EndSession(ctx)

	err = sess.StartTransaction()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	sc := m.NewSessionContext(ctx, sess)

	_, err = mongo.Posts.DeleteMany(sc, Post{
		CompanyID: companyID,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var docsToInsert []interface{}
	for _, doc := range newDocs {
		docsToInsert = append(docsToInsert, doc)
	}
	if len(docsToInsert) != 0 {
		_, err = mongo.Posts.InsertMany(sc, docsToInsert)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	err = sess.CommitTransaction(sc)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}
