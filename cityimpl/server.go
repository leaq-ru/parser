package cityimpl

import (
	"context"
	"errors"
	rd "github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/parser/city"
	"github.com/leaq-ru/parser/htmlfinder"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/redis"
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
	parser.UnimplementedCityServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) FindCity(ctx context.Context, req *parser.FindCityRequest) (res *parser.FindCityResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res = &parser.FindCityResponse{}

	if req.GetHtml() == "" {
		err = errors.New("html: cannot be blank")
		return
	}

	foundCity, isFound := htmlfinder.FindCity(req.GetHtml())
	if !isFound {
		return
	}

	cityModel := city.City{}
	dbCity, err := cityModel.GetOrCreate(ctx, foundCity)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res.CityId = dbCity.ID.Hex()
	res.IsFound = true
	return
}

func (s *server) GetAllCity(ctx context.Context, _ *empty.Empty) (res *parser.CitiesResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cur, err := mongo.Cities.Find(ctx, bson.D{})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return citiesCursorToCitiesResponse(ctx, cur)
}

func (s *server) GetCityById(ctx context.Context, req *parser.GetCityByIdRequest) (res *parser.CityItem, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCityId() == "" {
		err = errors.New("cityId is empty")
		return
	}

	cacheKey := strings.Join([]string{
		"city",
		req.GetCityId(),
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
		res = &parser.CityItem{}
		err = protojson.Unmarshal(cacheVal, res)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			err = errors.New(http.StatusText(http.StatusInternalServerError))
		}

		logger.Log.Debug().Str("cityID", req.GetCityId()).Msg("full cache")
		return
	}

	oID, err := primitive.ObjectIDFromHex(req.GetCityId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	cityModel := city.City{}
	err = mongo.Cities.FindOne(ctx, bson.M{
		"_id": oID,
	}).Decode(&cityModel)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = errors.New("city not found")
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CityItem{
		Id:    cityModel.ID.Hex(),
		Title: cityModel.Title,
		Slug:  cityModel.Slug,
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

func (s *server) GetCityByIds(ctx context.Context, req *parser.GetCityByIdsRequest) (res *parser.CitiesResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var oIDs []primitive.ObjectID
	for _, c := range req.GetCityIds() {
		oID, errOID := primitive.ObjectIDFromHex(c)
		if errOID != nil {
			err = errOID
			return
		}
		oIDs = append(oIDs, oID)
	}

	cur, err := mongo.Cities.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return citiesCursorToCitiesResponse(ctx, cur)
}

func (s *server) GetCityBySlug(ctx context.Context, req *parser.GetCityBySlugRequest) (res *parser.CityItem, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cityModel := city.City{}
	err = mongo.Cities.FindOne(ctx, bson.M{
		"s": req.GetSlug(),
	}).Decode(&cityModel)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = errors.New("city not found")
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CityItem{
		Id:    cityModel.ID.Hex(),
		Title: cityModel.Title,
		Slug:  cityModel.Slug,
	}
	return
}

func (s *server) GetCityHints(ctx context.Context, req *parser.GetCityHintsRequest) (res *parser.CitiesResponse, err error) {
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
	cur, err := mongo.Cities.Find(ctx, bson.M{
		"t": bson.M{
			"$regex":   regexp.QuoteMeta(req.GetTitle()),
			"$options": "$i",
		},
	}, opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	return citiesCursorToCitiesResponse(ctx, cur)
}
