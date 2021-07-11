package companyimpl

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-url-producer/protocol"
)

type server struct {
	parser.UnimplementedCompanyServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) ConsumeURL(m *stan.Msg) {
	go func() {
		msg := protocol.URLMessage{}
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		registrationDate, err := ptypes.TimestampProto(msg.RegistrationDate)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		_, err = s.Reindex(context.Background(), &parser.ReindexRequest{
			Url:              msg.URL,
			Registrar:        msg.Registrar,
			RegistrationDate: registrationDate,
		})
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		err = m.Ack()
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		logger.Log.Debug().Str("url", msg.URL).Msg("consumed")
	}()
}
