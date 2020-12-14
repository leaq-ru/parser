package main

import (
	graceful "github.com/nnqq/scr-lib-graceful"
	"github.com/nnqq/scr-parser/companyimpl"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/postimpl"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"strings"
	"sync"
)

func main() {
	srv := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	parser.RegisterCompanyServer(srv, companyimpl.NewServer())
	parser.RegisterPostServer(srv, postimpl.NewServer())

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		config.Env.Grpc.Port,
	}, ":"))
	logger.Must(err)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		graceful.HandleSignals(srv.GracefulStop)
	}()
	go func() {
		logger.Must(srv.Serve(lis))
	}()
	wg.Wait()
}
