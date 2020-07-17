package stan

import (
	"github.com/google/uuid"
	stand "github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"strings"
)

var Conn stand.Conn

func init() {
	sc, err := stand.Connect(config.Env.STAN.ClusterID, strings.Join([]string{
		"parser",
		uuid.New().String(),
	}, "-"))
	logger.Must(err)
	Conn = sc
}
