package call

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"google.golang.org/grpc"
)

var (
	Image      image.ImageClient
	City       city.CityClient
	Category   category.CategoryClient
	Technology technology.TechnologyClient
	DNS        technology.DnsClient
	Role       user.RoleClient
	User       user.UserClient
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
	DNS = technology.NewDnsClient(connTech)

	connUser, err := grpc.Dial(config.Env.Service.User, grpc.WithInsecure())
	logger.Must(err)
	Role = user.NewRoleClient(connUser)
	User = user.NewUserClient(connUser)
}
