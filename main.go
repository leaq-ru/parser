package main

import (
	"context"
	graceful "github.com/nnqq/scr-lib-graceful"
	"github.com/nnqq/scr-parser/categoryimpl"
	"github.com/nnqq/scr-parser/cityimpl"
	"github.com/nnqq/scr-parser/companyimpl"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/dnsimpl"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/postimpl"
	"github.com/nnqq/scr-parser/reviewimpl"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-parser/technologyimpl"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"strconv"
	"strings"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	srv := grpc.NewServer()
	comp := companyimpl.NewServer()

	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	parser.RegisterCompanyServer(srv, comp)
	parser.RegisterPostServer(srv, postimpl.NewServer())
	parser.RegisterReviewServer(srv, reviewimpl.NewServer())
	parser.RegisterCategoryServer(srv, categoryimpl.NewServer())
	parser.RegisterCityServer(srv, cityimpl.NewServer())
	parser.RegisterDnsServer(srv, dnsimpl.NewServer())
	parser.RegisterTechnologyServer(srv, technologyimpl.NewServer())

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		config.Env.Grpc.Port,
	}, ":"))
	logger.Must(err)

	urlMaxInFlight, err := strconv.Atoi(config.Env.STAN.URLMaxInFlight)
	logger.Must(err)

	urlConsumer, err := stan.NewConsumer(
		logger.Log,
		stan.Conn,
		config.Env.STAN.SubjectURL,
		urlMaxInFlight,
		comp.ConsumeURL,
	)
	logger.Must(err)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		graceful.HandleSignals(srv.GracefulStop, cancel)
	}()
	go func() {
		defer wg.Done()
		logger.Must(srv.Serve(lis))
	}()
	go func() {
		defer wg.Done()
		urlConsumer.Serve(ctx)
	}()
	wg.Wait()
}
