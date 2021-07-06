package reviewimpl

import (
	"context"
	safeerr "github.com/nnqq/scr-lib-safeerr"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/pagination"
	"github.com/nnqq/scr-parser/review"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type server struct {
	parser.UnimplementedReviewServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetReviews(ctx context.Context, req *parser.GetReviewsRequest) (*parser.GetReviewsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	compID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		return nil, safeerr.BadRequest
	}

	limit, err := pagination.ApplyDefaultLimit(req)
	if err != nil {
		return nil, err
	}

	reviews, err := review.Get(ctx, compID, int64(req.GetOpts().GetSkip()), int64(limit))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}

	mUserIDs := make(map[string]struct{}, len(reviews))
	for _, r := range reviews {
		mUserIDs[r.UserID.String()] = struct{}{}
	}
	userIDs := make([]string, len(mUserIDs))
	ind := 0
	for userID := range mUserIDs {
		userIDs[ind] = userID
		ind += 1
	}

	users, err := call.User.GetByIds(ctx, &user.GetByIdsRequest{
		UserIds: userIDs,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.BadRequest
	}
	mUsers := make(map[string]*user.ShortUser, len(users.GetUsers()))
	for _, u := range users.GetUsers() {
		mUsers[u.GetId()] = u
	}

	res := &parser.GetReviewsResponse{}
	res.Reviews = make([]*parser.ReviewItem, len(reviews))
	for i, r := range reviews {
		res.Reviews[i] = &parser.ReviewItem{
			Id:   r.ID.String(),
			Text: r.Text,
			User: mUsers[r.UserID.String()],
		}
	}

	return res, nil
}

func (s *server) Create(ctx context.Context, req *parser.CreateRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *server) Delete(ctx context.Context, req *parser.DeleteRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *server) ChangeStatus(ctx context.Context, req *parser.ChangeStatusRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *server) DisallowAuthorDeleteAll(ctx context.Context, req *parser.DisallowAuthorDeleteAllRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
