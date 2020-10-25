package postimpl

import "github.com/nnqq/scr-proto/codegen/go/parser"

type server struct {
	parser.UnimplementedPostServer
}

func NewServer() *server {
	return &server{}
}
