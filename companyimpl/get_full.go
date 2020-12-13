package companyimpl

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/technology"
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
			eg.Go(func() (e error) {
				cityItem, e := call.City.GetById(ctx, &city.GetByIdRequest{
					CityId: in.Location.CityID.Hex(),
				})
				if e != nil {
					return
				}

				location.City = cityItem
				return
			})
		}
	}

	var cat *category.CategoryItem
	if !in.CategoryID.IsZero() {
		eg.Go(func() (e error) {
			cat, e = call.Category.GetById(ctx, &category.GetByIdRequest{
				CategoryId: in.CategoryID.Hex(),
			})
			return
		})
	}

	var techCats []*parser.TechnologyCategory
	if len(in.TechnologyIDs) != 0 {
		eg.Go(func() (e error) {
			var ids []string
			for _, oID := range in.TechnologyIDs {
				ids = append(ids, oID.Hex())
			}

			techs, e := call.Technology.GetByIds(ctx, &technology.GetByIdsRequest{
				Ids: ids,
			})
			if e != nil {
				return
			}

			techCats, e = toTechnologyCategories(techs.GetTechnologies())
			return
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
		Avatar:               string(in.Avatar),
		Location:             location,
		App:                  app,
		Social:               social,
		People:               people,
		UpdatedAt:            in.UpdatedAt.String(),
		TechnologyCategories: techCats,
		PageSpeed:            in.PageSpeed,
		Verified:             in.Verified,
		Premium:              in.Premium,
	}
	return
}

func (s *server) GetFull(req *parser.GetV2Request, stream parser.Company_GetFullServer) (err error) {
	ctx, cancel := context.WithTimeout(stream.Context(), 3*time.Hour)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var opts *options.FindOptions
	if req.GetOpts().GetLimit() != 0 {
		opts = options.Find().SetLimit(int64(req.GetOpts().GetLimit()))
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
