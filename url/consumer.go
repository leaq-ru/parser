package url

import (
	"context"
	"encoding/json"
	s "github.com/nats-io/stan.go"
	logger "github.com/nnqq/scr-lib-logger"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-url-producer/protocol"
	"time"
)

type consumer struct {
	done chan struct{}
}

func NewConsumer() *consumer {
	return &consumer{
		done: make(chan struct{}),
	}
}

func (c *consumer) Serve() (err error) {
	_, err = stan.Conn.QueueSubscribe(
		"url",
		config.ServiceName,
		cb,
		s.DurableName(config.ServiceName),
		s.SetManualAckMode(),
		s.MaxInflight(25),
	)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	<-c.done
	return
}

func (c *consumer) GracefulStop() (err error) {
	err = stan.Conn.Close()
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	close(c.done)
	return
}

func cb(_m *s.Msg) {
	go func(m *s.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		msg := protocol.URLMessage{}
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		compModel := company.Company{}
		compModel.UpdateOrCreate(ctx, msg.URL, msg.Registrar, msg.RegistrationDate)
		err = m.Ack()
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		logger.Log.Debug().Str("url", msg.URL).Msg("URL consumed")
	}(_m)
}
