package companyimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type cityCat interface {
	GetCityId() string
	GetCategoryId() string
}

func makeQueryCityCat(req cityCat) (query bson.D, err error) {
	query = bson.D{}
	if req.GetCityId() != "" {
		oID, errOID := primitive.ObjectIDFromHex(req.GetCityId())
		if errOID != nil {
			err = errOID
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key:   "l.c",
			Value: oID,
		})
	}
	if req.GetCategoryId() != "" {
		oID, errOID := primitive.ObjectIDFromHex(req.GetCategoryId())
		if errOID != nil {
			err = errOID
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key:   "c",
			Value: oID,
		})
	}
	return
}

func withSelect(query bson.D, sel parser.Select, key string) bson.D {
	switch sel {
	case parser.Select_YES:
		return append(query, bson.E{
			Key: key,
			Value: bson.M{
				"$ne": nil,
			},
		})
	case parser.Select_NO:
		return append(query, bson.E{
			Key:   key,
			Value: nil,
		})
	default:
		return query
	}
}

func toFullCompany(
	inCompany model.Company,
	inCity *city.CityItem,
	inCategory *category.CategoryItem,
) (
	out *parser.FullCompany,
) {
	domain := &parser.Domain{}
	if inCompany.Domain != nil {
		domain.RegistrationDate = inCompany.Domain.RegistrationDate.String()
		domain.Registrar = inCompany.Domain.Registrar
		domain.Address = inCompany.Domain.Address
	}

	var app *parser.App
	if inCompany.App != nil {
		app = &parser.App{}
		if inCompany.App.GooglePlay != nil {
			app.GooglePlay = &parser.UrlItem{
				Url: inCompany.App.GooglePlay.URL,
			}
		}
		if inCompany.App.AppStore != nil {
			app.AppStore = &parser.UrlItem{
				Url: inCompany.App.AppStore.URL,
			}
		}
	}

	var social *parser.Social
	if inCompany.Social != nil {
		social = &parser.Social{}
		if inCompany.Social.Vk != nil {
			social.Vk = &parser.Vk{
				GroupId:      float64(inCompany.Social.Vk.GroupID),
				Name:         inCompany.Social.Vk.Name,
				ScreenName:   inCompany.Social.Vk.ScreenName,
				IsClosed:     parser.IsClosed(inCompany.Social.Vk.IsClosed),
				Description:  inCompany.Social.Vk.Description,
				MembersCount: float64(inCompany.Social.Vk.MembersCount),
				Photo_200:    string(inCompany.Social.Vk.Photo200),
			}
		}
		if inCompany.Social.Instagram != nil {
			social.Instagram = &parser.UrlItem{Url: inCompany.Social.Instagram.URL}
		}
		if inCompany.Social.Facebook != nil {
			social.Facebook = &parser.UrlItem{Url: inCompany.Social.Facebook.URL}
		}
		if inCompany.Social.Twitter != nil {
			social.Twitter = &parser.UrlItem{Url: inCompany.Social.Twitter.URL}
		}
		if inCompany.Social.Youtube != nil {
			social.Youtube = &parser.UrlItem{Url: inCompany.Social.Youtube.URL}
		}
	}

	var people []*parser.People
	for _, p := range inCompany.People {
		people = append(people, &parser.People{
			VkId:        float64(p.VkID),
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			VkIsClosed:  p.VkIsClosed,
			Sex:         parser.Sex(p.Sex),
			Photo_200:   string(p.Photo200),
			Phone:       float64(p.Phone),
			Email:       p.Email,
			Description: p.Description,
		})
	}

	var location *parser.FullLocation
	if inCompany.Location != nil {
		location = &parser.FullLocation{
			City:         inCity,
			Address:      inCompany.Location.Address,
			AddressTitle: inCompany.Location.AddressTitle,
		}
	}

	return &parser.FullCompany{
		Id:          inCompany.ID.Hex(),
		Category:    inCategory,
		Url:         inCompany.URL,
		Slug:        inCompany.Slug,
		Title:       inCompany.Title,
		Email:       inCompany.Email,
		Description: inCompany.Description,
		Online:      inCompany.Online,
		Phone:       float64(inCompany.Phone),
		Inn:         float64(inCompany.INN),
		Kpp:         float64(inCompany.KPP),
		Ogrn:        float64(inCompany.OGRN),
		Domain:      domain,
		Avatar:      string(inCompany.Avatar),
		Location:    location,
		App:         app,
		Social:      social,
		People:      people,
		UpdatedAt:   inCompany.UpdatedAt.String(),
	}
}

func toFullCompanies(
	inCompanies []model.Company,
	inCities *city.CitiesResponse,
	inCategories *category.CategoriesResponse,
) (
	out []*parser.FullCompany,
	err error,
) {
	mCity := map[string]*city.CityItem{}
	for _, c := range inCities.GetCities() {
		mCity[c.GetId()] = c
	}

	mCategory := map[string]*category.CategoryItem{}
	for _, c := range inCategories.GetCategories() {
		mCategory[c.GetId()] = c
	}

	for _, c := range inCompanies {
		var fullCity *city.CityItem
		if c.Location != nil && !c.Location.CityID.IsZero() {
			fc, ok := mCity[c.Location.CityID.Hex()]
			if !ok {
				err = errors.New(fmt.Sprintf("unexpected empty city id=%s", c.Location.CityID.Hex()))
				logger.Log.Error().Err(err).Send()
				return
			}
			fullCity = fc
		}

		var fullCategory *category.CategoryItem
		if !c.CategoryID.IsZero() {
			fc, ok := mCategory[c.CategoryID.Hex()]
			if !ok {
				err = errors.New(fmt.Sprintf("unexpected empty category id=%s", c.CategoryID.Hex()))
				logger.Log.Error().Err(err).Send()
				return
			}
			fullCategory = fc
		}

		out = append(out, toFullCompany(c, fullCity, fullCategory))
	}
	return
}

type GetQuerier interface {
	GetCityIds() []string
	GetCategoryIds() []string
	GetHasEmail() parser.Select
	GetHasPhone() parser.Select
	GetHasOnline() parser.Select
	GetHasInn() parser.Select
	GetHasKpp() parser.Select
	GetHasOgrn() parser.Select
	GetHasAppStore() parser.Select
	GetHasGooglePlay() parser.Select
	GetHasVk() parser.Select
	GetHasInstagram() parser.Select
	GetHasTwitter() parser.Select
	GetHasYoutube() parser.Select
	GetHasFacebook() parser.Select
	GetVkMembersCount() *parser.VkMembersCount
}

func makeGetQuery(req GetQuerier) (query bson.D, err error) {
	query = bson.D{}

	if len(req.GetCityIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, c := range req.GetCityIds() {
			oID, errOID := primitive.ObjectIDFromHex(c)
			if errOID != nil {
				err = errOID
				return
			}
			oIDs = append(oIDs, oID)
		}

		query = append(query, bson.E{
			Key: "l.c",
			Value: bson.M{
				"$in": oIDs,
			},
		})
	}
	if len(req.GetCategoryIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, c := range req.GetCategoryIds() {
			oID, errOID := primitive.ObjectIDFromHex(c)
			if errOID != nil {
				err = errOID
				return
			}
			oIDs = append(oIDs, oID)
		}

		query = append(query, bson.E{
			Key: "c",
			Value: bson.M{
				"$in": oIDs,
			},
		})
	}
	query = withSelect(query, req.GetHasEmail(), "e")
	query = withSelect(query, req.GetHasPhone(), "p")
	query = withSelect(query, req.GetHasOnline(), "o")
	query = withSelect(query, req.GetHasInn(), "i")
	query = withSelect(query, req.GetHasKpp(), "k")
	query = withSelect(query, req.GetHasOgrn(), "og")
	query = withSelect(query, req.GetHasAppStore(), "ap.a.u")
	query = withSelect(query, req.GetHasGooglePlay(), "ap.g.u")
	query = withSelect(query, req.GetHasVk(), "so.v.g")
	query = withSelect(query, req.GetHasInstagram(), "so.i.u")
	query = withSelect(query, req.GetHasTwitter(), "so.t.u")
	query = withSelect(query, req.GetHasYoutube(), "so.y.u")
	query = withSelect(query, req.GetHasFacebook(), "so.f.u")
	if req.GetVkMembersCount() != nil {
		value := bson.M{}
		if req.GetVkMembersCount().GetFrom() != 0 {
			value["$gt"] = req.GetVkMembersCount().GetFrom()
		}
		if req.GetVkMembersCount().GetTo() != 0 {
			value["$lt"] = req.GetVkMembersCount().GetTo()
		}

		if len(value) != 0 {
			query = append(query, bson.E{
				Key:   "so.v.m",
				Value: value,
			})
		}
	}
	return
}

func (s *server) Get(ctx context.Context, req *parser.GetRequest) (res *parser.GetResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(20)
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 || req.GetOpts().GetLimit() < 0 {
			err = errors.New("limit out of 1-100")
			return
		} else if req.GetOpts().GetLimit() != 0 {
			limit = int64(req.GetOpts().GetLimit())
		}
	}

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	// only documents with email or phone
	query = append(query, bson.E{
		Key: "$or",
		Value: []bson.E{
			makeNotNil(email),
			makeNotNil(phone),
		},
	})

	wgGet := sync.WaitGroup{}
	wgGet.Add(2)
	var (
		companies []model.Company
		errFind   error
	)
	go func() {
		defer wgGet.Done()

		opts := options.Find()
		opts.SetLimit(limit)
		opts.SetSort(bson.M{
			"_id": -1,
		})

		queryWithRangeID := query
		if req.GetOpts() != nil && req.GetOpts().GetFromId() != "" {
			oID, errOID := primitive.ObjectIDFromHex(req.GetOpts().GetFromId())
			if errOID != nil {
				err = errOID
				return
			}

			queryWithRangeID = append(query, bson.E{
				Key: "_id",
				Value: bson.M{
					"$lt": oID,
				},
			})
		}

		var cur *m.Cursor
		cur, errFind = mongo.Companies.Find(ctx, queryWithRangeID, opts)
		if errFind != nil {
			logger.Log.Error().Err(errFind).Send()
			return
		}
		logger.Err(cur.All(ctx, &companies))
	}()

	var (
		resTotalCount int64
		errTotalCount error
	)
	go func() {
		defer wgGet.Done()
		resTotalCount, errTotalCount = mongo.Companies.CountDocuments(
			ctx,
			query,
			options.Count().SetLimit(100000),
		)
		logger.Err(errTotalCount)
	}()
	wgGet.Wait()

	if errFind != nil {
		err = errFind
		return
	}
	if errTotalCount != nil {
		err = errTotalCount
		return
	}

	var cityIDs, categoryIDs []string
	for _, c := range companies {
		if c.Location != nil && !c.Location.CityID.IsZero() {
			cityIDs = append(cityIDs, c.Location.CityID.Hex())
		}
		if !c.CategoryID.IsZero() {
			categoryIDs = append(categoryIDs, c.CategoryID.Hex())
		}
	}

	wgFullDocs := sync.WaitGroup{}
	var (
		cities    *city.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			cities, errCities = call.City.GetByIds(ctx, &city.GetByIdsRequest{
				CityIds: cityIDs,
			})
			logger.Err(errCities)
		}()
	}

	var (
		categories    *category.CategoriesResponse
		errCategories error
	)
	if len(categoryIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			categories, errCategories = call.Category.GetByIds(ctx, &category.GetByIdsRequest{
				CategoryIds: categoryIDs,
			})
			logger.Err(errCategories)
		}()
	}
	wgFullDocs.Wait()

	if errCities != nil {
		err = errCities
		return
	}
	if errCategories != nil {
		err = errCategories
		return
	}

	res = &parser.GetResponse{
		TotalCount: uint32(resTotalCount),
	}
	res.Companies, err = toFullCompanies(companies, cities, categories)
	return
}
