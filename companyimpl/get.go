package companyimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"sync"
	"time"
)

type cityCat interface {
	GetCityId() string
	GetCategoryId() string
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

func toHex(in []primitive.ObjectID) (out []string) {
	for _, oID := range in {
		out = append(out, oID.Hex())
	}
	return
}

func toDNSItems(in []*parser.DnsItem) (out []*parser.DnsItem) {
	for _, item := range in {
		out = append(out, &parser.DnsItem{
			Id:   item.GetId(),
			Name: item.GetName(),
		})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return
}

func toFullCompany(
	inCompany company.Company,
	inCity *parser.CityItem,
	inCategory *parser.CategoryItem,
) (
	out *parser.FullCompany,
) {
	domain := &parser.Domain{}
	if inCompany.Domain != nil {
		if !inCompany.Domain.RegistrationDate.IsZero() {
			domain.RegistrationDate = inCompany.Domain.RegistrationDate.String()
		}

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
			Photo_200:   p.Photo200,
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
		Avatar:      inCompany.Avatar,
		Location:    location,
		App:         app,
		Social:      social,
		People:      people,
		UpdatedAt:   inCompany.UpdatedAt.String(),
	}
}

func toFullCompanies(
	inCompanies []company.Company,
	inCities *parser.CitiesResponse,
	inCategories *parser.CategoriesResponse,
) (
	out []*parser.FullCompany,
	err error,
) {
	mCity := map[string]*parser.CityItem{}
	for _, c := range inCities.GetCities() {
		mCity[c.GetId()] = c
	}

	mCategory := map[string]*parser.CategoryItem{}
	for _, c := range inCategories.GetCategories() {
		mCategory[c.GetId()] = c
	}

	for _, c := range inCompanies {
		var fullCity *parser.CityItem
		if c.Location != nil && !c.Location.CityID.IsZero() {
			fc, ok := mCity[c.Location.CityID.Hex()]
			if !ok {
				err = errors.New(fmt.Sprintf("unexpected empty city id=%s", c.Location.CityID.Hex()))
				logger.Log.Error().Err(err).Send()
				return
			}
			fullCity = fc
		}

		var fullCategory *parser.CategoryItem
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

type GetQuerierV1 interface {
	GetCityIds() []string
	GetCategoryIds() []string
	GetTechnologyIds() []string
	GetTechnologyFindRule() parser.FindRule
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

type GetQuerierV2 interface {
	GetCityIds() []string
	GetCategoryIds() []string
	GetTechnologyIds() []string
	GetTechnologyFindRule() parser.FindRule
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
	GetDnsIds() []string
}

func appendOIDs(inQuery bson.D, key, op string, ids []string) (outQuery bson.D, err error) {
	outQuery = inQuery
	if len(ids) == 0 {
		return
	}

	var oIDs []primitive.ObjectID
	for _, id := range ids {
		oID, errOID := primitive.ObjectIDFromHex(id)
		if errOID != nil {
			err = errOID
			return
		}
		oIDs = append(oIDs, oID)
	}

	outQuery = append(outQuery, bson.E{
		Key: key,
		Value: bson.M{
			op: oIDs,
		},
	})
	return
}

func makeGetQueryV1(req GetQuerierV1) (query bson.D, err error) {
	query = bson.D{}

	query, err = appendOIDs(query, "l.c", "$in", req.GetCityIds())
	if err != nil {
		return
	}

	query, err = appendOIDs(query, "c", "$in", req.GetCategoryIds())
	if err != nil {
		return
	}

	techOp := "$in"
	if req.GetTechnologyFindRule() == parser.FindRule_ALL {
		techOp = "$all"
	}
	query, err = appendOIDs(query, "ti", techOp, req.GetTechnologyIds())
	if err != nil {
		return
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

	query = append(query, bson.E{
		Key:   "h",
		Value: nil,
	})
	return
}

func makeGetQueryV2(req GetQuerierV2) (query bson.D, err error) {
	query = bson.D{}

	query, err = appendOIDs(query, "l.c", "$in", req.GetCityIds())
	if err != nil {
		return
	}

	query, err = appendOIDs(query, "c", "$in", req.GetCategoryIds())
	if err != nil {
		return
	}

	query, err = appendOIDs(query, "dn", "$in", req.GetDnsIds())
	if err != nil {
		return
	}

	techOp := "$in"
	if req.GetTechnologyFindRule() == parser.FindRule_ALL {
		techOp = "$all"
	}
	query, err = appendOIDs(query, "ti", techOp, req.GetTechnologyIds())
	if err != nil {
		return
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

	query = append(query, bson.E{
		Key:   "h",
		Value: nil,
	})
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

	query, err := makeGetQueryV1(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	if len(req.GetExcludeIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, id := range req.GetExcludeIds() {
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

	cur, err := mongo.Companies.Find(ctx, queryWithRangeID, opts)
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
		cities    *parser.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			cities, errCities = cityimpl.NewServer().GetCityByIds(ctx, &parser.GetCityByIdsRequest{
				CityIds: cityIDs,
			})
			logger.Err(errCities)
		}()
	}

	var (
		categories    *parser.CategoriesResponse
		errCategories error
	)
	if len(categoryIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			categories, errCategories = categoryimpl.NewServer().GetCategoryByIds(ctx, &parser.GetCategoryByIdsRequest{
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

	res = &parser.GetResponse{}
	res.Companies, err = toFullCompanies(companies, cities, categories)
	return
}
