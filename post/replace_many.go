package post

import (
	"context"
	"errors"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/vk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ReplaceMany(ctx context.Context, companyID primitive.ObjectID, vkGroupID int) (err error) {
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
		"count":    100,
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

	if len(newDocs) == 0 {
		logger.Log.Debug().Msg("no newDocs to replace. Skip inserting")
		return nil
	}

	count, err := mongo.Posts.CountDocuments(ctx, Post{
		CompanyID: companyID,
	}, options.Count().SetLimit(100))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if int(count) > len(newDocs) {
		logger.Log.Debug().Msg("posts count in mongo > newDocs to replace. Skip inserting")
		return nil
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

	err = m.WithSession(ctx, sess, func(sc m.SessionContext) (e error) {
		_, e = mongo.Posts.DeleteMany(sc, Post{
			CompanyID: companyID,
		})
		if e != nil {
			logger.Log.Error().Err(e).Send()
			return
		}

		var docsToInsert []interface{}
		for _, doc := range newDocs {
			docsToInsert = append(docsToInsert, doc)
		}

		_, e = mongo.Posts.InsertMany(sc, docsToInsert)
		if e != nil {
			logger.Log.Error().Err(e).Send()
		}
		return
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()

		errAbort := sess.AbortTransaction(ctx)
		if errAbort != nil {
			err = errAbort
			logger.Log.Error().Err(err).Send()
		}
		return
	}

	err = sess.CommitTransaction(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}
