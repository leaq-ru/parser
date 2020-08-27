package stan

import (
	"github.com/google/uuid"
	s "github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"strings"
)

var Conn s.Conn

func connect() (sc s.Conn, err error) {
	return s.Connect(
		config.Env.STAN.ClusterID,
		strings.Join([]string{
			"parser",
			uuid.New().String(),
		}, "-"),
		s.NatsURL(config.Env.NATS.URL),
	)
}

func init() {
	var err error
	Conn, err = connect()
	logger.Must(err)
}
