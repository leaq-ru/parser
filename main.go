package main

import (
	"context"
	graceful "github.com/leaq-ru/lib-graceful"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/companyimpl"
	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/parser/dnsimpl"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/postimpl"
	"github.com/leaq-ru/parser/reviewimpl"
	"github.com/leaq-ru/parser/stan"
	"github.com/leaq-ru/parser/technologyimpl"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"strconv"
	"strings"
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

	url, err := stan.NewConsumer(
		logger.Log,
		stan.Conn,
		config.Env.STAN.SubjectURL,
		config.ServiceName,
		urlMaxInFlight,
		comp.ConsumeURL,
	)
	logger.Must(err)

	analyzeResult, err := stan.NewConsumer(
		logger.Log,
		stan.Conn,
		config.Env.STAN.SubjectAnalyzeResult,
		config.ServiceName,
		0,
		comp.ConsumeAnalyzeResult,
	)
	logger.Must(err)

	imageUploadResult, err := stan.NewConsumer(
		logger.Log,
		stan.Conn,
		config.Env.STAN.SubjectImageUploadResult,
		config.ServiceName,
		0,
		comp.ConsumeImageUploadResult,
	)
	logger.Must(err)

	var eg errgroup.Group
	eg.Go(func() error {
		graceful.HandleSignals(srv.GracefulStop, cancel)
		return nil
	})
	eg.Go(func() error {
		return srv.Serve(lis)
	})
	eg.Go(func() error {
		url.Serve(ctx)
		return nil
	})
	eg.Go(func() error {
		analyzeResult.Serve(ctx)
		return nil
	})
	eg.Go(func() error {
		imageUploadResult.Serve(ctx)
		return nil
	})
	logger.Must(eg.Wait())
}
