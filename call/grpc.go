package call

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"google.golang.org/grpc"
)

var (
	Image image.ImageClient
	City  city.CityClient
)

func init() {
	connImage, err := grpc.Dial(config.Env.Service.Image, grpc.WithInsecure())
	logger.Must(err)

	connCity, err := grpc.Dial(config.Env.Service.City, grpc.WithInsecure())
	logger.Must(err)

	Image = image.NewImageClient(connImage)
	City = city.NewCityClient(connCity)
}
