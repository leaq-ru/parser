package main

import (
	graceful "github.com/nnqq/scr-lib-graceful"
	"github.com/nnqq/scr-parser/companyimpl"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/url"
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
	urlConsumer := url.NewConsumer()
	go graceful.HandleSignals(srv.GracefulStop, urlConsumer.GracefulStop)

	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	parser.RegisterCompanyServer(srv, companyimpl.NewServer())

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		config.Env.Grpc.Port,
	}, ":"))
	logger.Must(err)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		logger.Must(srv.Serve(lis))
	}()
	go func() {
		defer wg.Done()
		logger.Must(urlConsumer.Serve())
	}()
	wg.Wait()
}
