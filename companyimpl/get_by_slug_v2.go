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
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"sort"
	"sync"
	"time"
)

// invert technology slice with categories to
// category slice (sorted by name asc) with technologies
func toTechnologyCategories(in []*technology.TechnologyItem) (
	out []*parser.TechnologyCategory,
	err error,
) {
	if len(in) == 0 {
		return
	}

	type catID = string
	cats := map[catID]*technology.CategoryItem{}
	techs := map[catID][]*technology.TechnologyItem{}

	for _, tech := range in {
		for _, cat := range tech.GetCategories() {
			cats[cat.GetId()] = cat
			techs[cat.GetId()] = append(techs[cat.GetId()], tech)
		}
	}

	for catID, cat := range cats {
		outTechs, ok := techs[catID]
		if !ok {
			err = errors.New(fmt.Sprintf("unexpected empty technology id=%s", catID))
			logger.Log.Error().Err(err).Send()
			return
		}

		outCat := &parser.TechnologyCategory{
			Id:   cat.GetId(),
			Name: cat.GetName(),
		}
		for _, tech := range outTechs {
			outCat.Technologies = append(outCat.Technologies, &parser.TechnologyItem{
				Id:      tech.GetId(),
				Name:    tech.GetName(),
				Version: tech.GetVersion(),
				Slug:    tech.GetSlug(),
			})
		}

		out = append(out, outCat)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return
}

func (s *server) GetBySlugV2(ctx context.Context, req *parser.GetBySlugRequest) (
	res *parser.GetBySlugResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	comp := model.Company{}
	err = mongo.Companies.FindOne(ctx, bson.M{
		"s": req.GetSlug(),
	}).Decode(&comp)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = errors.New("company not found")
			return
		}

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

	var (
		resTechs *technology.GetByIdsResponse
		errTechs error
	)
	if len(comp.TechnologyIDs) != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var ids []string
			for _, oID := range comp.TechnologyIDs {
				ids = append(ids, oID.Hex())
			}

			resTechs, errTechs = call.Technology.GetByIds(ctx, &technology.GetByIdsRequest{
				Ids: ids,
			})
			if errTechs != nil {
				logger.Log.Error().Err(errTechs).Send()
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
	if errTechs != nil {
		err = errTechs
		return
	}

	techCats, err := toTechnologyCategories(resTechs.GetTechnologies())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetBySlugResponse{
		FullCompany:          toFullCompany(comp, resCity, resCategory),
		PageSpeed:            comp.PageSpeed,
		TechnologyCategories: techCats,
	}
	return
}