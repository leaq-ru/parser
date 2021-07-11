package call

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/classifier"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"github.com/nnqq/scr-proto/codegen/go/wappalyzer"
	"google.golang.org/grpc"
)

var (
	Image      image.ImageClient
	Role       user.RoleClient
	User       user.UserClient
	Classifier classifier.ClassifierClient
	Wappalyzer wappalyzer.WappalyzerClient
)

func init() {
	connImage, err := grpc.Dial(config.Env.Service.Image, grpc.WithInsecure())
	logger.Must(err)
	Image = image.NewImageClient(connImage)

	connUser, err := grpc.Dial(config.Env.Service.User, grpc.WithInsecure())
	logger.Must(err)
	Role = user.NewRoleClient(connUser)
	User = user.NewUserClient(connUser)

	connClassifier, err := grpc.Dial(config.Env.Service.Classifier, grpc.WithInsecure())
	logger.Must(err)
	Classifier = classifier.NewClassifierClient(connClassifier)

	connWappalyzer, err := grpc.Dial(config.Env.Service.Wappalyzer, grpc.WithInsecure())
	logger.Must(err)
	Wappalyzer = wappalyzer.NewWappalyzerClient(connWappalyzer)
}
