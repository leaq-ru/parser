package companyimpl

import (
	"context"
	"errors"
	"github.com/SevereCloud/vksdk/v2/api"
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
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"strconv"
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

	var setMu sync.Mutex
	set := bson.M{
		"v":  true,
		"ua": time.Now().UTC(),
	}
	unset := bson.M{}

	var categoryOIDToValidate primitive.ObjectID
	if req.GetCategoryId() != nil {
		if req.GetCategoryId().GetValue() != "" {
			categoryOIDToValidate, err = primitive.ObjectIDFromHex(req.GetCategoryId().GetValue())
			if err != nil {
				logger.Log.Error().Err(err).Send()
				return
			}
		} else {
			unset["c"] = ""
		}
	}

	var cityOIDToValidate primitive.ObjectID
	if req.GetCityId() != nil {
		if req.GetCityId().GetValue() != "" {
			cityOIDToValidate, err = primitive.ObjectIDFromHex(req.GetCityId().GetValue())
			if err != nil {
				logger.Log.Error().Err(err).Send()
				return
			}
		} else {
			unset["l.c"] = ""
		}
	}

	if req.GetTitle() != nil {
		if req.GetTitle().GetValue() != "" {
			if len([]rune(req.GetTitle().GetValue())) > 50 {
				err = errors.New("title too long. max 50 symbols")
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
			if len(email) > 50 {
				err = errors.New("email too long. max 50 symbols")
				return
			}

			if email != "" {
				set["e"] = email
			}
		} else {
			unset["e"] = ""
		}
	}
	if req.GetPhone() != nil {
		if req.GetPhone().GetValue() != 0 {
			runePhone := []rune(strconv.Itoa(int(req.GetPhone().GetValue())))
			validPhone := len(runePhone) == 11 || runePhone[0] == []rune("7")[0]
			if !validPhone {
				err = errors.New("phone invalid")
				return
			}

			set["p"] = req.GetPhone().GetValue()
		} else {
			unset["p"] = ""
		}
	}
	if req.GetAddressStreet() != nil {
		if req.GetAddressStreet().GetValue() != "" {
			if len([]rune(req.GetAddressStreet().GetValue())) > 50 {
				err = errors.New("address too long. max 50 symbols")
				return
			}
			set["l.a"] = strings.TrimSpace(req.GetAddressStreet().GetValue())
		} else {
			unset["l.a"] = ""
		}
	}
	if req.GetAddressHouse() != nil {
		if req.GetAddressHouse().GetValue() != "" {
			if len([]rune(req.GetAddressHouse().GetValue())) > 50 {
				err = errors.New("address too long. max 50 symbols")
				return
			}
			set["l.at"] = strings.TrimSpace(req.GetAddressHouse().GetValue())
		} else {
			unset["l.at"] = ""
		}
	}
	if req.GetInstagramUrl() != nil {
		err = setURLQuery(req.GetInstagramUrl(), "https://www.instagram.com/", "so.i.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetYoutubeUrl() != nil {
		err = setURLQuery(req.GetYoutubeUrl(), "https://www.youtube.com/", "so.y.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetFacebookUrl() != nil {
		err = setURLQuery(req.GetFacebookUrl(), "https://www.facebook.com/", "so.f.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetTwitterUrl() != nil {
		err = setURLQuery(req.GetTwitterUrl(), "https://twitter.com/", "so.t.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetAppStoreUrl() != nil {
		err = setURLQuery(req.GetAppStoreUrl(), "https://apps.apple.com/", "ap.a.u", set, unset)
		if err != nil {
			return
		}
	}
	if req.GetGooglePlayUrl() != nil {
		err = setURLQuery(req.GetGooglePlayUrl(), "https://play.google.com/", "ap.g.u", set, unset)
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

	var (
		needDeleteCompPosts bool
		compBodyForVk       company.Company
	)
	if req.GetVkUrl() != nil {
		if req.GetVkUrl().GetValue() != "" {
			compBodyForVk.DigVk(ctx, req.GetVkUrl().GetValue())
		} else {
			needDeleteCompPosts = true
			unset["so.v"] = ""
			unset["pe"] = ""
		}
	}

	if compBodyForVk.GetSocial().GetVk().GetGroupId() != 0 {
		set["so.v"] = compBodyForVk.GetSocial().GetVk()

		if len(compBodyForVk.People) != 0 {
			set["pe"] = compBodyForVk.People
		}
	}

	var eg errgroup.Group
	if !cityOIDToValidate.IsZero() {
		eg.Go(func() (e error) {
			cityItem, e := call.City.GetById(ctx, &city.GetByIdRequest{
				CityId: req.GetCityId().GetValue(),
			})
			if e != nil {
				return
			}

			if cityItem.GetId() == "" {
				e = errors.New("cityId invalid")
				return
			}

			setMu.Lock()
			set["l.c"] = cityOIDToValidate
			setMu.Unlock()
			return
		})
	}

	if !categoryOIDToValidate.IsZero() {
		eg.Go(func() (e error) {
			categoryItem, e := call.Category.GetById(ctx, &category.GetByIdRequest{
				CategoryId: req.GetCategoryId().GetValue(),
			})
			if e != nil {
				return
			}

			if categoryItem.GetId() == "" {
				e = errors.New("categoryId invalid")
				return
			}

			setMu.Lock()
			set["c"] = categoryOIDToValidate
			setMu.Unlock()
			return
		})
	}

	if compBodyForVk.GetSocial().GetVk().GetGroupId() != 0 {
		eg.Go(func() (e error) {
			return post.ReplaceMany(
				ctx,
				compOID,
				compBodyForVk.GetSocial().GetVk().GetGroupId(),
				true,
			)
		})
	}

	var needDeleteOldAvatar bool
	if req.GetAvatarBase64().GetValue() != "" {
		eg.Go(func() (e error) {
			newAvatar, e := call.Image.PutBase64(ctx, &image.PutBase64Request{
				Base64: req.GetAvatarBase64().GetValue(),
			})
			if e != nil {
				return
			}

			if newAvatar.GetS3Url() != "" {
				needDeleteOldAvatar = true
				setMu.Lock()
				set["a"] = newAvatar.GetS3Url()
				setMu.Unlock()
			}
			return
		})
	}
	err = eg.Wait()
	if err != nil {
		if errors.Is(err, api.ErrAccess) {
			err = errors.New("vk group access denied")
			logger.Log.Error().
				Int("vkGroupID", compBodyForVk.GetSocial().GetVk().GetGroupId()).
				Err(err).
				Send()
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	query := bson.M{
		"$set": set,
	}
	if len(unset) != 0 {
		query["$unset"] = unset
	}

	sess, err := mongo.Client.StartSession()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	defer sess.EndSession(ctx)

	err = sess.StartTransaction()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	sc := m.NewSessionContext(ctx, sess)

	var oldComp company.Company
	err = mongo.Companies.FindOneAndUpdate(sc, company.Company{
		ID: compOID,
	}, query, options.FindOneAndUpdate().SetReturnDocument(options.Before)).Decode(&oldComp)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if needDeleteCompPosts {
		_, err = mongo.Posts.DeleteMany(sc, post.Post{
			CompanyID: compOID,
		})
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	err = sess.CommitTransaction(sc)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if needDeleteOldAvatar && oldComp.Avatar != "" {
		_, err = call.Image.Remove(ctx, &image.RemoveRequest{
			S3Url: string(oldComp.Avatar),
		})
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	res = &empty.Empty{}
	return
}

type valuer interface {
	GetValue() string
}

func setURLQuery(val valuer, expectedPrefix, queryKey string, set, unset bson.M) (err error) {
	if len(val.GetValue()) > 250 {
		err = errors.New("url too long. max 250 symbols")
	}

	if val.GetValue() != "" {
		if !strings.HasPrefix(val.GetValue(), expectedPrefix) {
			err = errors.New("url invalid")
			logger.Log.Error().
				Str("val", val.GetValue()).
				Str("expectedPrefix", expectedPrefix).
				Err(err).
				Send()
			return
		}

		set[queryKey] = val.GetValue()
	} else {
		unset[queryKey] = ""
	}
	return
}
