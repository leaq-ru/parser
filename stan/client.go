package stan

import (
	"github.com/google/uuid"
	s "github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"os"
	"strings"
	"syscall"
	"time"
)

var Conn s.Conn

func pollAlive(sc s.Conn) {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		<-t.C

		if sc.NatsConn().IsClosed() {
			p, err := os.FindProcess(os.Getpid())
			logger.Must(err)
			logger.Must(p.Signal(syscall.SIGTERM))
		}
	}
}

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
	go pollAlive(Conn)
}
