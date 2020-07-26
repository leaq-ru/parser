package companyimpl

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"
)

func makeProtoFullCompany(in model.Company, inCity *city.CityItem,
	inCategory *category.CategoryItem) *parser.FullCompany {
	domain := &parser.Domain{}
	if in.Domain != nil {
		domain.RegistrationDate = in.Domain.RegistrationDate.String()
		domain.Registrar = in.Domain.Registrar
		domain.Address = in.Domain.Address
	}

	address := ""
	addressTitle := ""
	if in.Location != nil {
		address = in.Location.Address
		addressTitle = in.Location.AddressTitle
	}

	var app *parser.App
	if in.App != nil {
		app = &parser.App{}
		if in.App.GooglePlay != nil {
			app.GooglePlay = &parser.UrlItem{
				Url: in.App.GooglePlay.URL,
			}
		}
		if in.App.AppStore != nil {
			app.AppStore = &parser.UrlItem{
				Url: in.App.AppStore.URL,
			}
		}
	}

	var social *parser.Social
	if in.Social != nil {
		social = &parser.Social{}
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
		if in.Social.Instagram != nil {
			social.Instagram = &parser.UrlItem{Url: in.Social.Instagram.URL}
		}
		if in.Social.Facebook != nil {
			social.Instagram = &parser.UrlItem{Url: in.Social.Facebook.URL}
		}
		if in.Social.Twitter != nil {
			social.Instagram = &parser.UrlItem{Url: in.Social.Twitter.URL}
		}
		if in.Social.Youtube != nil {
			social.Instagram = &parser.UrlItem{Url: in.Social.Youtube.URL}
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

	return &parser.FullCompany{
		Id:          in.ID.Hex(),
		Category:    inCategory,
		Url:         in.URL,
		Slug:        in.Slug,
		Title:       in.Title,
		Email:       in.Email,
		Description: in.Description,
		Online:      in.Online,
		Phone:       float64(in.Phone),
		Inn:         float64(in.INN),
		Kpp:         float64(in.KPP),
		Ogrn:        float64(in.OGRN),
		Domain:      domain,
		Avatar:      string(in.Avatar),
		Location: &parser.FullLocation{
			City:         inCity,
			Address:      address,
			AddressTitle: addressTitle,
		},
		App:       app,
		Social:    social,
		People:    people,
		UpdatedAt: in.UpdatedAt.String(),
	}
}

func (s *server) GetBySlug(ctx context.Context, req *parser.GetBySlugRequest) (res *parser.FullCompany, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	comp := model.Company{}
	err = mongo.Companies.FindOne(ctx, bson.M{
		"s": req.GetSlug(),
	}).Decode(&comp)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	wg := sync.WaitGroup{}
	var (
		resCity *city.CityItem
		errCity error
	)
	if comp.Location != nil && !comp.Location.CityID.IsZero() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resCity, errCity = call.City.GetById(ctx, &city.GetByIdRequest{
				CityId: comp.Location.CityID.Hex(),
			})
			if errCity != nil {
				logger.Log.Error().Err(errCity).Send()
			}
		}()
	}

	var (
		resCategory *category.CategoryItem
		errCategory error
	)
	if !comp.CategoryID.IsZero() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resCategory, errCategory = call.Category.GetById(ctx, &category.GetByIdRequest{
				CategoryId: comp.CategoryID.Hex(),
			})
			if errCategory != nil {
				logger.Log.Error().Err(errCategory).Send()
			}
		}()
	}
	wg.Wait()

	if errCity != nil {
		err = errCity
		return
	}
	if errCategory != nil {
		err = errCategory
		return
	}

	res = makeProtoFullCompany(comp, resCity, resCategory)
	return
}
