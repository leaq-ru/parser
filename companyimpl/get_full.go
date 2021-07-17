package companyimpl

import (
	"context"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/dnsimpl"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/technologyimpl"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"time"
)

func fetchFullCompanyV2(ctx context.Context, in company.Company) (out *parser.FullCompanyV2, err error) {
	var eg errgroup.Group
	var location *parser.FullLocation
	if in.Location != nil {
		location = &parser.FullLocation{
			Address:      in.Location.Address,
			AddressTitle: in.Location.AddressTitle,
		}

		if !in.Location.CityID.IsZero() {
			eg.Go(func() error {
				cityItem, e := cityimpl.NewServer().GetCityById(ctx, &parser.GetCityByIdRequest{
					CityId: in.Location.CityID.Hex(),
				})
				if e != nil {
					return e
				}

				location.City = cityItem
				return nil
			})
		}
	}

	var cat *parser.CategoryItem
	if !in.CategoryID.IsZero() {
		eg.Go(func() (e error) {
			cat, e = categoryimpl.NewServer().GetCategoryById(ctx, &parser.GetCategoryByIdRequest{
				CategoryId: in.CategoryID.Hex(),
			})
			return e
		})
	}

	var techCats []*parser.TechCategoryInverted
	if len(in.TechnologyIDs) != 0 {
		eg.Go(func() error {
			techs, e := technologyimpl.NewServer().GetTechByIds(ctx, &parser.GetTechByIdsRequest{
				Ids: toHex(in.TechnologyIDs),
			})
			if e != nil {
				return e
			}

			techCats, e = toTechnologyCategories(techs.GetTechnologies())
			return e
		})
	}

	var dnsItems []*parser.DnsItem
	if len(in.DNSIDs) != 0 {
		eg.Go(func() error {
			resDNS, e := dnsimpl.NewServer().GetDnsByIds(ctx, &parser.GetDnsByIdsRequest{
				Ids: toHex(in.DNSIDs),
			})
			if e != nil {
				return e
			}

			dnsItems = toDNSItems(resDNS.GetDns())
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var domain *parser.Domain
	if in.Domain != nil {
		domain = &parser.Domain{
			Address:          in.Domain.Address,
			Registrar:        in.Domain.Registrar,
			RegistrationDate: in.Domain.RegistrationDate.String(),
		}
	}

	app := &parser.App{
		AppStore:   &parser.UrlItem{},
		GooglePlay: &parser.UrlItem{},
	}
	if in.App != nil {
		if in.App.AppStore != nil {
			app.AppStore.Url = in.App.AppStore.URL
		}
		if in.App.GooglePlay != nil {
			app.GooglePlay.Url = in.App.GooglePlay.URL
		}
	}

	social := &parser.Social{
		Vk:        &parser.Vk{},
		Instagram: &parser.UrlItem{},
		Twitter:   &parser.UrlItem{},
		Youtube:   &parser.UrlItem{},
		Facebook:  &parser.UrlItem{},
	}
	if in.Social != nil {
		if in.Social.Vk != nil {
			social.Vk = &parser.Vk{
				GroupId:      float64(in.Social.Vk.GroupID),
				Name:         in.Social.Vk.Name,
				ScreenName:   in.Social.Vk.ScreenName,
				IsClosed:     parser.IsClosed(in.Social.Vk.IsClosed),
				Description:  in.Social.Vk.Description,
				MembersCount: float64(in.Social.Vk.MembersCount),
				Photo_200:    string(in.Social.Vk.Photo200),
			}
		}
	}

	var people []*parser.People
	for _, p := range in.People {
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

	out = &parser.FullCompanyV2{
		Id:                   in.ID.Hex(),
		Category:             cat,
		Url:                  in.URL,
		Slug:                 in.Slug,
		Title:                in.Title,
		Email:                in.Email,
		Description:          in.Description,
		Online:               in.Online,
		Phone:                float64(in.Phone),
		Inn:                  float64(in.INN),
		Kpp:                  float64(in.KPP),
		Ogrn:                 float64(in.OGRN),
		Domain:               domain,
		Avatar:               in.Avatar,
		Location:             location,
		App:                  app,
		Social:               social,
		People:               people,
		UpdatedAt:            in.UpdatedAt.String(),
		TechnologyCategories: techCats,
		PageSpeed:            in.PageSpeed,
		Verified:             in.Verified,
		Premium:              in.Premium,
		Dns:                  dnsItems,
	}
	return
}

func (s *server) GetFull(req *parser.GetFullRequest, stream parser.Company_GetFullServer) (err error) {
	ctx, cancel := context.WithTimeout(stream.Context(), 10*time.Hour)
	defer cancel()

	query, err := makeGetQueryV2(req.GetQuery())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if req.GetFromId() != "" {
		oID, e := primitive.ObjectIDFromHex(req.GetFromId())
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key: "_id",
			Value: bson.M{
				"$gt": oID,
			},
		})
	}

	opts := options.Find().SetSort(bson.M{
		"_id": 1,
	})
	if req.GetQuery().GetOpts().GetLimit() != 0 {
		opts.SetLimit(int64(req.GetQuery().GetOpts().GetLimit()))
	}

	cur, err := mongo.Companies.Find(ctx, query, opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for cur.Next(ctx) {
		var doc company.Company
		err = cur.Decode(&doc)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		fullComp, e := fetchFullCompanyV2(ctx, doc)
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}

		err = stream.Send(fullComp)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}
	return
}
