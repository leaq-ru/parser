package mock

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-proto/codegen/go/city"
	m "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type City struct {
	m.Mock
}

func (c *City) Find(ctx context.Context, in *city.FindRequest, opts ...grpc.CallOption) (*city.FindResponse, error) {
	args := c.Called(ctx, in, opts)
	return args.Get(0).(*city.FindResponse), args.Error(1)
}

func (c *City) GetById(_ context.Context, _ *city.GetByIdRequest, _ ...grpc.CallOption) (*city.CityItem, error) {
	panic("implement me")
}

func (c *City) GetBySlug(_ context.Context, _ *city.GetBySlugRequest, _ ...grpc.CallOption) (*city.CityItem, error) {
	panic("implement me")
}

func (c *City) GetByIds(_ context.Context, _ *city.GetByIdsRequest, _ ...grpc.CallOption) (*city.CitiesResponse, error) {
	panic("implement me")
}

func (c *City) GetAll(_ context.Context, _ *empty.Empty, _ ...grpc.CallOption) (*city.CitiesResponse, error) {
	panic("implement me")
}

func (c *City) GetHints(_ context.Context, _ *city.GetHintsRequest, _ ...grpc.CallOption) (*city.CitiesResponse, error) {
	panic("implement me")
}
