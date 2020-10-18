package call

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"google.golang.org/grpc"
)

var (
	Image      image.ImageClient
	City       city.CityClient
	Category   category.CategoryClient
	Technology technology.TechnologyClient
)

func init() {
	connImage, err := grpc.Dial(config.Env.Service.Image, grpc.WithInsecure())
	logger.Must(err)
	Image = image.NewImageClient(connImage)

	connCity, err := grpc.Dial(config.Env.Service.City, grpc.WithInsecure())
	logger.Must(err)
	City = city.NewCityClient(connCity)

	connCategory, err := grpc.Dial(config.Env.Service.Category, grpc.WithInsecure())
	logger.Must(err)
	Category = category.NewCategoryClient(connCategory)

	connTech, err := grpc.Dial(config.Env.Service.Technology, grpc.WithInsecure())
	logger.Must(err)
	Technology = technology.NewTechnologyClient(connTech)
}
