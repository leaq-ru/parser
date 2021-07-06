package reviewimpl

import (
	"context"
	"errors"
	"fmt"
	safeerr "github.com/nnqq/scr-lib-safeerr"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/md"
	"github.com/nnqq/scr-parser/pagination"
	"github.com/nnqq/scr-parser/review"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		return nil, safeerr.InvalidParam("companyId")
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
		mUserIDs[r.UserID.Hex()] = struct{}{}
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
		return nil, safeerr.InternalServerError
	}
	mUsers := make(map[string]*user.ShortUser, len(users.GetUsers()))
	for _, u := range users.GetUsers() {
		mUsers[u.GetId()] = u
	}

	res := &parser.GetReviewsResponse{}
	res.Reviews = make([]*parser.ReviewItem, len(reviews))
	for i, r := range reviews {
		res.Reviews[i] = &parser.ReviewItem{
			Id:        r.ID.Hex(),
			Text:      r.Text,
			User:      mUsers[r.UserID.Hex()],
			CreatedAt: r.ID.Timestamp().String(),
			Positive:  r.Positive,
		}
	}

	return res, nil
}

func (s *server) Create(ctx context.Context, req *parser.CreateRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	const (
		maxLen        = 3000
		maxModeration = 10
	)
	if len(req.GetText()) > maxLen {
		return nil, fmt.Errorf("text too long, max length %d", maxLen)
	}

	compID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		return nil, safeerr.InvalidParam("companyId")
	}
	userID, err := md.GetUserOID(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, err
	}

	userData, err := call.User.GetById(ctx, &user.GetByIdRequest{
		UserId: userID.Hex(),
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}
	if userData.GetBanReview() {
		return nil, errors.New("you not allowed to post reviews")
	}

	countMod, err := review.CountModeration(ctx, userID)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}
	if countMod >= maxModeration {
		return nil, errors.New("too many reviews in moderation, try later")
	}

	err = review.Create(ctx, compID, userID, req.GetText(), req.GetPositive())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}

	// TODO: STAN PRODUCE MSG

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *parser.DeleteRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	reviewID, err := primitive.ObjectIDFromHex(req.GetReviewId())
	if err != nil {
		return nil, safeerr.InvalidParam("reviewId")
	}
	userID, err := md.GetUserOID(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, err
	}

	err = review.Delete(ctx, reviewID, userID, true)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}

	return &emptypb.Empty{}, nil
}

func (s *server) ChangeStatus(ctx context.Context, req *parser.ChangeStatusRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	reviewID, err := primitive.ObjectIDFromHex(req.GetReviewId())
	if err != nil {
		return nil, safeerr.InvalidParam("reviewId")
	}

	switch req.GetStatus() {
	case parser.ReviewStatus_OK:
		err = review.SetOK(ctx, reviewID)
	case parser.ReviewStatus_DELETE:
		err = review.Delete(ctx, reviewID, primitive.ObjectID{}, false)
	default:
		err = safeerr.InvalidParam("status")
	}
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DisallowAuthorDeleteAll(ctx context.Context, req *parser.DisallowAuthorDeleteAllRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, safeerr.InvalidParam("userId")
	}

	_, err = call.User.ModifyRights(ctx, &user.ModifyRightsRequest{
		UserId:    req.GetUserId(),
		BanReview: wrapperspb.Bool(true),
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, err
	}

	err = review.DeleteAll(ctx, userID)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
