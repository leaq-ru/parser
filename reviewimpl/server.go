package reviewimpl

import (
	"context"
	"errors"
	"fmt"
	safeerr "github.com/leaq-ru/lib-safeerr"
	"github.com/leaq-ru/parser/call"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/md"
	"github.com/leaq-ru/parser/pagination"
	"github.com/leaq-ru/parser/review"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"github.com/leaq-ru/proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
	"unicode/utf8"
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

	if len(reviews) == 0 {
		return &parser.GetReviewsResponse{}, nil
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
	if utf8.RuneCountInString(req.GetText()) > maxLen {
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

	var userData *user.ShortUser
	var eg errgroup.Group
	eg.Go(func() error {
		u, e := call.User.GetById(ctx, &user.GetByIdRequest{
			UserId: userID.Hex(),
		})
		if e != nil {
			logger.Log.Error().Err(e).Send()
			return safeerr.InternalServerError
		}
		userData = u
		if u.GetBanReview() {
			return errors.New("you not allowed to post reviews")
		}
		return nil
	})
	eg.Go(func() error {
		countMod, e := review.CountModeration(ctx, userID)
		if e != nil {
			logger.Log.Error().Err(e).Send()
			return safeerr.InternalServerError
		}
		if countMod >= maxModeration {
			return errors.New("too many reviews in moderation, try later")
		}
		return nil
	})
	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	r, err := review.Create(ctx, compID, userID, req.GetText(), req.GetPositive())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}

	err = review.ProduceModeration(ctx, &parser.ReviewItem{
		Id:        r.ID.Hex(),
		Text:      r.Text,
		User:      userData,
		CreatedAt: r.ID.Timestamp().String(),
		Positive:  r.Positive,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return nil, safeerr.InternalServerError
	}

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
