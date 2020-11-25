package companyimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func toShortCompany(
	inCompany company.Company,
	inCity *city.CityItem,
	inCategory *category.CategoryItem,
) (
	out *parser.ShortCompany,
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

	var shortSocial *parser.ShortSocial
	if inCompany.Social != nil {
		shortSocial = &parser.ShortSocial{}
		if inCompany.Social.Vk != nil {
			shortSocial.Vk = &parser.ShortVk{
				ScreenName:   inCompany.Social.Vk.ScreenName,
				MembersCount: float64(inCompany.Social.Vk.MembersCount),
				GroupId:      float64(inCompany.Social.Vk.GroupID),
			}
		}
		if inCompany.Social.Instagram != nil {
			shortSocial.Instagram = &parser.UrlItem{Url: inCompany.Social.Instagram.URL}
		}
		if inCompany.Social.Facebook != nil {
			shortSocial.Facebook = &parser.UrlItem{Url: inCompany.Social.Facebook.URL}
		}
		if inCompany.Social.Twitter != nil {
			shortSocial.Twitter = &parser.UrlItem{Url: inCompany.Social.Twitter.URL}
		}
		if inCompany.Social.Youtube != nil {
			shortSocial.Youtube = &parser.UrlItem{Url: inCompany.Social.Youtube.URL}
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

	var cityItem *city.CityItem
	if inCompany.Location != nil {
		cityItem = &city.CityItem{
			Id:    inCity.GetId(),
			Title: inCity.GetTitle(),
			Slug:  inCity.GetSlug(),
		}
	}

	return &parser.ShortCompany{
		Id:        inCompany.ID.Hex(),
		City:      cityItem,
		Category:  inCategory,
		Url:       inCompany.URL,
		Slug:      inCompany.Slug,
		Title:     inCompany.Title,
		Email:     inCompany.Email,
		Phone:     float64(inCompany.Phone),
		Avatar:    string(inCompany.Avatar),
		App:       app,
		Social:    shortSocial,
		UpdatedAt: inCompany.UpdatedAt.String(),
		Verified:  inCompany.Verified,
		Premium:   inCompany.Premium,
	}
}

func toMyCompany(
	inCompany company.Company,
	inCity *city.CityItem,
	inCategory *category.CategoryItem,
) (
	out *parser.MyCompany,
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

	var shortSocial *parser.ShortSocial
	if inCompany.Social != nil {
		shortSocial = &parser.ShortSocial{}
		if inCompany.Social.Vk != nil {
			shortSocial.Vk = &parser.ShortVk{
				ScreenName:   inCompany.Social.Vk.ScreenName,
				MembersCount: float64(inCompany.Social.Vk.MembersCount),
				GroupId:      float64(inCompany.Social.Vk.GroupID),
			}
		}
		if inCompany.Social.Instagram != nil {
			shortSocial.Instagram = &parser.UrlItem{Url: inCompany.Social.Instagram.URL}
		}
		if inCompany.Social.Facebook != nil {
			shortSocial.Facebook = &parser.UrlItem{Url: inCompany.Social.Facebook.URL}
		}
		if inCompany.Social.Twitter != nil {
			shortSocial.Twitter = &parser.UrlItem{Url: inCompany.Social.Twitter.URL}
		}
		if inCompany.Social.Youtube != nil {
			shortSocial.Youtube = &parser.UrlItem{Url: inCompany.Social.Youtube.URL}
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

	var cityItem *city.CityItem
	if inCompany.Location != nil {
		cityItem = &city.CityItem{
			Id:    inCity.GetId(),
			Title: inCity.GetTitle(),
			Slug:  inCity.GetSlug(),
		}
	}

	return &parser.MyCompany{
		Id:              inCompany.ID.Hex(),
		City:            cityItem,
		Category:        inCategory,
		Url:             inCompany.URL,
		Slug:            inCompany.Slug,
		Title:           inCompany.Title,
		Email:           inCompany.Email,
		Phone:           float64(inCompany.Phone),
		Avatar:          string(inCompany.Avatar),
		App:             app,
		Social:          shortSocial,
		UpdatedAt:       inCompany.UpdatedAt.String(),
		Verified:        inCompany.Verified,
		Premium:         inCompany.Premium,
		PremiumDeadline: inCompany.PremiumDeadline.String(),
	}
}

func toShortCompanies(
	inCompanies []company.Company,
	inCities *city.CitiesResponse,
	inCategories *category.CategoriesResponse,
) (
	out []*parser.ShortCompany,
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

		out = append(out, toShortCompany(c, fullCity, fullCategory))
	}
	return
}

func toMyCompanies(
	inCompanies []company.Company,
	inCities *city.CitiesResponse,
	inCategories *category.CategoriesResponse,
) (
	out []*parser.MyCompany,
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

		out = append(out, toMyCompany(c, fullCity, fullCategory))
	}
	return
}

func fetchShortCompanies(ctx context.Context, companies []company.Company) (
	res *parser.ShortCompanies,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

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

	res = &parser.ShortCompanies{}
	res.Companies, err = toShortCompanies(companies, cities, categories)
	return
}

func fetchMyCompanies(ctx context.Context, companies []company.Company) (
	res *parser.GetMyResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

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

	res = &parser.GetMyResponse{}
	res.Companies, err = toMyCompanies(companies, cities, categories)
	return
}

func (s *server) GetV2(ctx context.Context, req *parser.GetV2Request) (res *parser.ShortCompanies, err error) {
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
	if len(req.GetOpts().GetExcludeIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, id := range req.GetOpts().GetExcludeIds() {
			oID, e := primitive.ObjectIDFromHex(id)
			if e != nil {
				err = e
				logger.Log.Error().Err(err).Send()
				return
			}

			oIDs = append(oIDs, oID)
		}

		query = append(query, bson.E{
			Key: "_id",
			Value: bson.M{
				"$nin": oIDs,
			},
		})
	}
	if len(req.GetCompanyIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, id := range req.GetCompanyIds() {
			oID, errOID := primitive.ObjectIDFromHex(id)
			if errOID != nil {
				err = errOID
				return
			}
			oIDs = append(oIDs, oID)
		}

		query = append(query, bson.E{
			Key: "_id",
			Value: bson.M{
				"$in": oIDs,
			},
		})
	}

	opts := options.Find()
	opts.SetSkip(int64(req.GetOpts().GetSkip()))
	opts.SetLimit(limit)
	opts.SetSort(bson.M{
		"pr":  -1,
		"_id": -1,
	})

	cur, err := mongo.Companies.Find(ctx, query, opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var companies []company.Company
	err = cur.All(ctx, &companies)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	return fetchShortCompanies(ctx, companies)
}
