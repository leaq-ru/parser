package dnsimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/dns"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"golang.org/x/sync/errgroup"
	"net"
	"net/url"
	"sync"
	"time"
)

type server struct {
	parser.UnimplementedDnsServer
}

func NewServer() *server {
	return &server{}
}

func (*server) FindDns(
	ctx context.Context,
	req *parser.FindDnsRequest,
) (
	res *parser.FindDnsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	u, err := url.Parse(req.GetUrl())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	nss, err := net.DefaultResolver.LookupNS(ctx, u.Host)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var (
		mu     sync.Mutex
		dnsIDs []string
	)
	var eg errgroup.Group
	for _, _ns := range nss {
		ns := _ns
		if ns == nil {
			continue
		}

		eg.Go(func() error {
			fromDB, e := dns.DNS{}.Create(ctx, ns.Host)
			if e != nil {
				return e
			}

			mu.Lock()
			dnsIDs = append(dnsIDs, fromDB.ID.Hex())
			mu.Unlock()
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.FindDnsResponse{}
	for _, d := range dnsIDs {
		res.Ids = append(res.Ids, d)
	}
	return
}

func (*server) GetDnsByIds(ctx context.Context, req *parser.GetDnsByIdsRequest) (
	res *parser.GetDnsByIdsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dnss, err := dns.DNS{}.GetByIDs(ctx, req.GetIds())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetDnsByIdsResponse{}
	for _, d := range dnss {
		res.Dns = append(res.Dns, &parser.DnsItem{
			Id:   d.ID.Hex(),
			Name: d.Name,
		})
	}
	return
}

func (*server) GetDnsHints(ctx context.Context, req *parser.GetDnsHintsRequest) (
	res *parser.GetDnsHintsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(20)
	if req.GetLimit() > 100 || req.GetLimit() < 0 {
		err = errors.New("limit out of 1-100")
		return
	} else if req.GetLimit() != 0 {
		limit = int64(req.GetLimit())
	}

	dnss, err := dns.DNS{}.GetHints(ctx, req.GetName(), req.GetExcludeIds(), limit)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetDnsHintsResponse{}
	for _, d := range dnss {
		res.Dns = append(res.Dns, &parser.DnsItem{
			Id:   d.ID.Hex(),
			Name: d.Name,
		})
	}
	return
}
