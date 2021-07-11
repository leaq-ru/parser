package technologyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/technology"
	"github.com/nnqq/scr-parser/technology_category"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/wappalyzer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"
)

type server struct {
	parser.UnimplementedTechnologyServer
}

func NewServer() *server {
	return &server{}
}

func (*server) FindTech(ctx context.Context, req *parser.FindTechRequest) (res *parser.FindTechResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	waCallStart := time.Now()
	resAnalyze, err := call.Wappalyzer.Analyze(ctx, &wappalyzer.AnalyzeRequest{Url: req.GetUrl()})
	logger.Log.Debug().
		Str("url", req.GetUrl()).
		Dur("ms", time.Since(waCallStart)).
		Msg("got Wappalyzer.Analyze response")
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var (
		wgCats     sync.WaitGroup
		muCats     sync.RWMutex
		cats       = map[string][]technology_category.TechnologyCategory{}
		muErrsCats sync.Mutex
		errsCats   []error
	)
	for _, tech := range resAnalyze.GetTechnologies() {
		for _, cat := range tech.GetCategories() {
			wgCats.Add(1)
			go func(techName string, catID uint32) {
				defer wgCats.Done()

				muCats.RLock()
				_, ok := cats[techName]
				muCats.RUnlock()
				if ok {
					return
				}

				resCategory, e := technology_category.TechnologyCategory{}.Create(ctx, catID)
				if e != nil {
					logger.Log.Error().Err(e).Send()
					muErrsCats.Lock()
					errsCats = append(errsCats, e)
					muErrsCats.Unlock()
					return
				}

				muCats.Lock()
				cats[techName] = append(cats[techName], resCategory)
				muCats.Unlock()
			}(tech.GetName(), cat.GetId())
		}
	}
	wgCats.Wait()

	if len(errsCats) > 0 {
		err = errsCats[0]
		return
	}

	var (
		wgTechs     sync.WaitGroup
		muTechs     sync.Mutex
		techs       = map[string]technology.Technology{}
		muErrsTechs sync.Mutex
		errsTechs   []error
	)
	for _, tech := range resAnalyze.GetTechnologies() {
		cat, ok := cats[tech.GetName()]
		if !ok {
			err = errors.New("expected to get value from map, but nothing found")
			logger.Log.Error().Str("tech.GetName()", tech.GetName()).Err(err).Send()
			return
		}

		var oIDs []primitive.ObjectID
		for _, c := range cat {
			oIDs = append(oIDs, c.ID)
		}

		wgTechs.Add(1)
		go func(techName string) {
			defer wgTechs.Done()

			resTechnology, e := technology.Technology{}.Create(ctx, techName, oIDs)
			if e != nil {
				logger.Log.Error().Err(e).Send()
				muErrsTechs.Lock()
				errsTechs = append(errsTechs, e)
				muErrsTechs.Unlock()
				return
			}

			muTechs.Lock()
			techs[techName] = resTechnology
			muTechs.Unlock()
		}(tech.GetName())
	}
	wgTechs.Wait()

	if len(errsTechs) > 0 {
		err = errsTechs[0]
		return
	}

	res = &parser.FindTechResponse{}
	for _, tech := range techs {
		res.Ids = append(res.Ids, tech.ID.Hex())
	}
	return
}

func (*server) GetTechByIds(ctx context.Context, req *parser.GetTechByIdsRequest) (
	res *parser.GetTechByIdsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	techs, err := technology.Technology{}.GetByIDs(ctx, req.GetIds())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	uniqueCatIDs := map[string]struct{}{}
	for _, tech := range techs {
		for _, catID := range tech.CategoryIDs {
			uniqueCatIDs[catID.Hex()] = struct{}{}
		}
	}

	var catIDs []string
	for oID := range uniqueCatIDs {
		catIDs = append(catIDs, oID)
	}

	cats, err := technology_category.TechnologyCategory{}.GetByIDs(ctx, catIDs)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	mapCats := map[primitive.ObjectID]technology_category.TechnologyCategory{}
	for _, cat := range cats {
		mapCats[cat.ID] = cat
	}

	res = &parser.GetTechByIdsResponse{}
	for _, tech := range techs {
		resTech := &parser.TechItem{
			Id:   tech.ID.Hex(),
			Name: tech.Name,
			Slug: tech.Slug,
		}

		for _, catID := range tech.CategoryIDs {
			cat, ok := mapCats[catID]
			if !ok {
				err = errors.New("expected to get value from map, but nothing found")
				logger.Log.Error().Str("catID", catID.Hex()).Err(err).Send()
				return
			}

			resTech.Categories = append(resTech.Categories, &parser.TechCategoryItem{
				Id:   cat.ID.Hex(),
				Name: cat.Name,
			})
		}

		res.Technologies = append(res.Technologies, resTech)
	}
	return
}

func (*server) GetTechBySlug(ctx context.Context, req *parser.GetTechBySlugRequest) (
	res *parser.GetTechBySlugResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tech, err := technology.Technology{}.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var categoryIDs []string
	for _, oID := range tech.CategoryIDs {
		categoryIDs = append(categoryIDs, oID.Hex())
	}

	cats, err := technology_category.TechnologyCategory{}.GetByIDs(ctx, categoryIDs)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetTechBySlugResponse{
		Technology: &parser.TechItem{
			Id:   tech.ID.Hex(),
			Name: tech.Name,
			Slug: tech.Slug,
		},
	}
	for _, resCatID := range tech.CategoryIDs {
		var resCategory *technology_category.TechnologyCategory
		for _, cat := range cats {
			if cat.ID == resCatID {
				resCategory = &cat
				break
			}
		}
		if resCategory == nil {
			err = errors.New("category not found")
			logger.Log.Error().Str("resCatID", resCatID.Hex()).Err(err).Send()
			return
		}

		res.Technology.Categories = append(res.Technology.Categories, &parser.TechCategoryItem{
			Id:   resCategory.ID.Hex(),
			Name: resCategory.Name,
		})
	}
	return
}

func (*server) GetTechHints(ctx context.Context, req *parser.GetTechHintsRequest) (
	res *parser.GetTechHintsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(20)
	if req.GetLimit() > 100 || req.GetLimit() < 0 {
		err = errors.New("limit out of 1-100")
		return
	} else if req.GetLimit() != 0 {
		limit = int64(req.GetLimit())
	}

	techs, err := technology.Technology{}.GetHints(ctx, req.GetName(), req.GetExcludeIds(), limit)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetTechHintsResponse{}
	for _, tech := range techs {
		res.Technologies = append(res.Technologies, &parser.HintTechItem{
			Id:   tech.ID.Hex(),
			Name: tech.Name,
		})
	}
	return
}
