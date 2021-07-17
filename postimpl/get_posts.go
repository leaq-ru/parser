package postimpl

import (
	"context"
	"errors"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/post"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (*server) GetPosts(ctx context.Context, req *parser.GetPostsRequest) (res *parser.GetPostsResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	companyOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		return
	}

	var excludeOIDs []primitive.ObjectID
	for _, id := range req.GetOpts().GetExcludeIds() {
		oID, e := primitive.ObjectIDFromHex(id)
		if e != nil {
			err = e
			return
		}
		excludeOIDs = append(excludeOIDs, oID)
	}

	limit := uint32(20)
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 || req.GetOpts().GetLimit() < 0 {
			err = errors.New("limit out of 1-100")
			return
		} else if req.GetOpts().GetLimit() != 0 {
			limit = req.GetOpts().GetLimit()
		}
	}

	posts, err := post.Get(ctx, companyOID, req.GetOpts().GetSkip(), limit, excludeOIDs)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetPostsResponse{}
	for _, item := range posts {
		resPost := &parser.PostItem{
			Id:        item.ID.Hex(),
			CompanyId: item.CompanyID.Hex(),
			Date:      item.Date.String(),
			Text:      item.Text,
		}

		for _, photo := range item.Photos {
			resPost.Photos = append(resPost.Photos, &parser.PhotoItem{
				UrlM: photo.URLm,
				UrlR: photo.URLr,
			})
		}

		res.Posts = append(res.Posts, resPost)
	}
	return
}
