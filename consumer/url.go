package consumer

import (
	"context"
	"encoding/json"
	stand "github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-url-producer/protocol"
)

func cb(ctx context.Context) func(*stand.Msg) {
	return func(m *stand.Msg) {
		go func() {
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
			}
		}()
	}
}

func URL(ctx context.Context) (err error) {
	_, err = stan.Conn.QueueSubscribe(
		"url",
		config.ServiceName,
		cb(ctx),
		stand.DurableName(config.ServiceName),
		stand.SetManualAckMode(),
		stand.MaxInflight(30),
	)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	<-ctx.Done()
	return
}
