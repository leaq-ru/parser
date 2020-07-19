package call

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"google.golang.org/grpc"
)

var (
	GrpcConn *grpc.ClientConn
	Image    image.ImageClient
)

func init() {
	conn, err := grpc.Dial(config.Env.Service.Image, grpc.WithInsecure())
	logger.Must(err)

	GrpcConn = conn
	Image = image.NewImageClient(conn)
}
