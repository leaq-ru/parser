package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/md"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/post"
	"github.com/nnqq/scr-parser/rx"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"net/url"
	"strings"
	"sync"
	"time"
)

func (*server) Edit(ctx context.Context, req *parser.EditRequest) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCompanyId() == "" {
		err = errors.New("companyId required")
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		return
	}

	resRole, err := call.Role.CanEditCompany(ctx, &user.CanEditCompanyRequest{
		CompanyId: req.GetCompanyId(),
		UserId:    authUserID,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if !resRole.GetCanEdit() {
		err = errors.New("unauthorized")
		return
	}

	compOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var categoryOID primitive.ObjectID
	if req.GetCategoryId() != nil {
		categoryOID, err = primitive.ObjectIDFromHex(req.GetCategoryId().GetValue())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	var cityOID primitive.ObjectID
	if req.GetCityId() != nil {
		cityOID, err = primitive.ObjectIDFromHex(req.GetCityId().GetValue())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	var setMu sync.Mutex
	set := bson.M{
		"v":  true,
		"ua": time.Now().UTC(),
	}
	unset := bson.M{}

	if req.GetTitle() != nil {
		if req.GetTitle().GetValue() != "" {
			if len([]rune(req.GetTitle().GetValue())) > 48 {
				err = errors.New("title too long. max 48 symbols")
				return
			}
			set["t"] = strings.TrimSpace(req.GetTitle().GetValue())
		} else {
			unset["t"] = ""
		}
	}
	if req.GetDescription() != nil {
		if req.GetDescription().GetValue() != "" {
			if len([]rune(req.GetDescription().GetValue())) > 1500 {
				err = errors.New("description too long. max 1500 symbols")
				return
			}
			set["d"] = strings.TrimSpace(req.GetDescription().GetValue())
		} else {
			unset["d"] = ""
		}
	}
	if req.GetEmail() != nil {
		if req.GetEmail().GetValue() != "" {
			email := rx.Email.FindString(req.GetEmail().GetValue())
			if email != "" {
				set["e"] = email
			}
		} else {
			unset["e"] = ""
		}
	}
	if req.GetAddressStreet() != nil {
		if req.GetAddressStreet().GetValue() != "" {
			if len([]rune(req.GetAddressStreet().GetValue())) > 48 {
				err = errors.New("address too long. max 48 symbols")
				return
			}
			set["l.c.a"] = strings.TrimSpace(req.GetAddressStreet().GetValue())
		} else {
			unset["l.c.a"] = ""
		}
	}
	if req.GetAddressHouse() != nil {
		if req.GetAddressHouse().GetValue() != "" {
			if len([]rune(req.GetAddressHouse().GetValue())) > 48 {
				err = errors.New("address too long. max 48 symbols")
				return
			}
			set["l.c.at"] = strings.TrimSpace(req.GetAddressHouse().GetValue())
		} else {
			unset["l.c.at"] = ""
		}
	}
	if req.GetInstagramUrl() != nil {
		err = setURLQuery(req.GetInstagramUrl(), "instagram.com", "so.i.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetYoutubeUrl() != nil {
		err = setURLQuery(req.GetYoutubeUrl(), "youtube.com", "so.y.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetFacebookUrl() != nil {
		err = setURLQuery(req.GetFacebookUrl(), "facebook.com", "so.f.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetTwitterUrl() != nil {
		err = setURLQuery(req.GetTwitterUrl(), "twitter.com", "so.t.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetAppStoreUrl() != nil {
		err = setURLQuery(req.GetAppStoreUrl(), "apps.apple.com", "ap.a.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetGooglePlayUrl() != nil {
		err = setURLQuery(req.GetGooglePlayUrl(), "play.google.com", "ap.g.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetInn() != nil {
		if req.GetInn().GetValue() != 0 {
			set["i"] = req.GetInn().GetValue()
		} else {
			unset["i"] = ""
		}
	}
	if req.GetKpp() != nil {
		if req.GetKpp().GetValue() != 0 {
			set["k"] = req.GetKpp().GetValue()
		} else {
			unset["k"] = ""
		}
	}
	if req.GetOgrn() != nil {
		if req.GetOgrn().GetValue() != 0 {
			set["og"] = req.GetOgrn().GetValue()
		} else {
			unset["og"] = ""
		}
	}

	var eg errgroup.Group
	var compBodyForVk company.Company
	eg.Go(func() (e error) {
		compBodyForVk.DigVk(ctx, req.GetVkUrl().GetValue())
		return
	})

	if req.GetCityId() != nil {
		eg.Go(func() (e error) {
			cityItem, e := call.City.GetById(ctx, &city.GetByIdRequest{
				CityId: req.GetCityId().GetValue(),
			})
			if e != nil {
				return
			}

			if cityItem.GetId() == "" {
				e = errors.New("cityId invalid")
			}

			setMu.Lock()
			set["l.c"] = cityOID
			setMu.Unlock()
			return
		})
	}

	if req.GetCategoryId() != nil {
		eg.Go(func() (e error) {
			categoryItem, e := call.Category.GetById(ctx, &category.GetByIdRequest{
				CategoryId: req.GetCategoryId().GetValue(),
			})
			if e != nil {
				return
			}

			if categoryItem.GetId() == "" {
				e = errors.New("categoryId invalid")
			}

			setMu.Lock()
			set["c"] = categoryOID
			setMu.Unlock()
			return
		})
	}

	if compBodyForVk.GetSocial().GetVk().GetGroupId() != 0 {
		eg.Go(func() (e error) {
			return post.ReplaceMany(ctx, compOID, compBodyForVk.GetSocial().GetVk().GetGroupId())
		})
	}

	if req.GetAvatarBase64() != nil {
		eg.Go(func() error {
			call.Image.PutBase64()
		})
	}
	err = eg.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if compBodyForVk.GetSocial().GetVk().GetGroupId() != 0 {
		set["so.v"] = compBodyForVk.GetSocial().GetVk()

		if len(compBodyForVk.People) != 0 {
			set["pe"] = compBodyForVk.People
		}
	}

	_, err = mongo.Companies.UpdateOne(ctx, company.Company{
		ID: compOID,
	}, bson.M{
		"$set":   set,
		"$unset": unset,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &empty.Empty{}
	return
}

type valuer interface {
	GetValue() string
}

func setURLQuery(val valuer, expectedHost, queryKey string, set, unset bson.M) (err error) {
	if val.GetValue() != "" {
		invalidSocial := errors.New("url invalid")

		socialURL, e := url.Parse(val.GetValue())
		if e != nil {
			err = invalidSocial
			return
		}

		if socialURL.Host != expectedHost {
			err = invalidSocial
			return
		}

		if socialURL.Scheme == "" {
			socialURL.Scheme = "https://"
		}

		set[queryKey] = socialURL.String()
	} else {
		unset[queryKey] = ""
	}
	return
}
