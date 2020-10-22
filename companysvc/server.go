package companysvc

import "github.com/nnqq/scr-proto/codegen/go/parser"

type server struct {
	parser.UnimplementedCompanyServer
}

func NewServer() *server {
	return &server{}
}
