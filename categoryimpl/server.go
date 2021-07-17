package categoryimpl

import (
	"context"
	"errors"
	rd "github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jbrukh/bayesian"
	"github.com/leaq-ru/parser/call"
	"github.com/leaq-ru/parser/category"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/redis"
	"github.com/leaq-ru/proto/codegen/go/classifier"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type server struct {
	parser.UnimplementedCategoryServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) FindCategory(ctx context.Context, req *parser.FindCategoryRequest) (res *parser.FindCategoryResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res = &parser.FindCategoryResponse{}

	if req.GetHtml() == "" {
		err = errors.New("html: cannot be blank")
		return
	}

	resCl, err := call.Classifier.Predict(ctx, &classifier.PredictRequest{
		Html: req.GetHtml(),
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	categoryModel := category.Category{}
	dbCategory, err := categoryModel.GetOrCreate(ctx, bayesian.Class(resCl.GetCategoryClass()))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res.CategoryId = dbCategory.ID.Hex()
	return
}

func (s *server) GetAllCategory(ctx context.Context, _ *empty.Empty) (res *parser.CategoriesResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cur, err := mongo.Categories.Find(ctx, bson.D{})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return categoriesCursorToCategoriesResponse(ctx, cur)
}

func (s *server) GetCategoryById(ctx context.Context, req *parser.GetCategoryByIdRequest) (res *parser.CategoryItem, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCategoryId() == "" {
		err = errors.New("categoryId is empty")
		return
	}

	cacheKey := strings.Join([]string{
		"category",
		req.GetCategoryId(),
	}, ":")

	cacheVal, err := redis.Client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		if errors.Is(err, rd.Nil) {
			err = nil
		} else {
			logger.Log.Error().Err(err).Send()
			err = errors.New(http.StatusText(http.StatusInternalServerError))
			return
		}
	} else {
		res = &parser.CategoryItem{}
		err = protojson.Unmarshal(cacheVal, res)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			err = errors.New(http.StatusText(http.StatusInternalServerError))
		}

		logger.Log.Debug().Str("categoryID", req.GetCategoryId()).Msg("full cache")
		return
	}

	oID, err := primitive.ObjectIDFromHex(req.GetCategoryId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	cat := category.Category{}
	err = mongo.Categories.FindOne(ctx, bson.M{
		"_id": oID,
	}).Decode(&cat)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = ErrCategoryNotFound
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CategoryItem{
		Id:    cat.ID.Hex(),
		Title: string(cat.Title),
		Slug:  cat.Slug,
	}

	newCacheVal, err := protojson.Marshal(res)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = errors.New(http.StatusText(http.StatusInternalServerError))
		return
	}

	err = redis.Client.Set(ctx, cacheKey, newCacheVal, 30*time.Second).Err()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = errors.New(http.StatusText(http.StatusInternalServerError))
	}
	return
}

func (s *server) GetCategoryByIds(ctx context.Context, req *parser.GetCategoryByIdsRequest) (res *parser.CategoriesResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var oIDs []primitive.ObjectID
	for _, c := range req.GetCategoryIds() {
		oID, errOID := primitive.ObjectIDFromHex(c)
		if errOID != nil {
			err = errOID
			return
		}
		oIDs = append(oIDs, oID)
	}

	cur, err := mongo.Categories.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return categoriesCursorToCategoriesResponse(ctx, cur)
}

func (s *server) GetCategoryBySlug(ctx context.Context, req *parser.GetCategoryBySlugRequest) (res *parser.CategoryItem, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cat := category.Category{}
	err = mongo.Categories.FindOne(ctx, bson.M{
		"s": req.GetSlug(),
	}).Decode(&cat)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = ErrCategoryNotFound
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CategoryItem{
		Id:    cat.ID.Hex(),
		Title: string(cat.Title),
		Slug:  cat.Slug,
	}
	return
}

func (s *server) GetCategoryHints(ctx context.Context, req *parser.GetCategoryHintsRequest) (
	res *parser.CategoriesResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(20)
	if req.GetLimit() > 100 || req.GetLimit() < 0 {
		err = errors.New("limit out of 1-100")
		return
	} else if req.GetLimit() != 0 {
		limit = int64(req.GetLimit())
	}

	opts := options.Find()
	opts.SetLimit(limit)
	cur, err := mongo.Categories.Find(ctx, bson.M{
		"t": bson.M{
			"$regex":   regexp.QuoteMeta(req.GetTitle()),
			"$options": "$i",
		},
	}, opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return categoriesCursorToCategoriesResponse(ctx, cur)
}
