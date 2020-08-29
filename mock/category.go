package mock

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-proto/codegen/go/category"
	m "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type Category struct {
	m.Mock
}

func (c *Category) Find(ctx context.Context, in *category.FindRequest, opts ...grpc.CallOption) (*category.FindResponse, error) {
	args := c.Called(ctx, in, opts)
	return args.Get(0).(*category.FindResponse), args.Error(1)
}

func (c *Category) GetById(_ context.Context, _ *category.GetByIdRequest, _ ...grpc.CallOption) (*category.CategoryItem, error) {
	panic("implement me")
}

func (c *Category) GetBySlug(_ context.Context, _ *category.GetBySlugRequest, _ ...grpc.CallOption) (*category.CategoryItem, error) {
	panic("implement me")
}

func (c *Category) GetByIds(_ context.Context, _ *category.GetByIdsRequest, _ ...grpc.CallOption) (*category.CategoriesResponse, error) {
	panic("implement me")
}

func (c *Category) GetAll(_ context.Context, _ *empty.Empty, _ ...grpc.CallOption) (*category.CategoriesResponse, error) {
	panic("implement me")
}

func (c *Category) GetHints(_ context.Context, _ *category.GetHintsRequest, _ ...grpc.CallOption) (*category.CategoriesResponse, error) {
	panic("implement me")
}
