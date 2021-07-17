package postimpl

import "github.com/leaq-ru/proto/codegen/go/parser"

type server struct {
	parser.UnimplementedPostServer
}

func NewServer() *server {
	return &server{}
}
