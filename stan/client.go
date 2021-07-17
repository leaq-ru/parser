package stan

import (
	"github.com/google/uuid"
	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/parser/logger"
	s "github.com/nats-io/stan.go"
	"strings"
)

var Conn s.Conn

func init() {
	sc, err := s.Connect(
		config.Env.STAN.ClusterID,
		strings.Join([]string{
			config.ServiceName,
			uuid.New().String(),
		}, "-"),
		s.NatsURL(config.Env.NATS.URL),
	)
	logger.Must(err)
	Conn = sc
}
