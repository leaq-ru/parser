package companyimpl

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	st "github.com/nats-io/stan.go"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-parser/technologyimpl"
	"github.com/nnqq/scr-proto/codegen/go/event"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-url-producer/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

type server struct {
	parser.UnimplementedCompanyServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) ConsumeURL(_m *st.Msg) {
	go func(m *st.Msg) {
		if m.RedeliveryCount >= 3 {
			err := m.Ack()
			if err != nil {
				logger.Log.Error().Err(err).Send()
			}
			return
		}

		logger.Log.Debug().
			Str("value", string(m.Data)).
			Str("subject", "url").
			Msg("recieved message")

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

		_, err = s.reindex(context.Background(), &parser.ReindexRequest{
			Url:              msg.URL,
			Registrar:        msg.Registrar,
			RegistrationDate: registrationDate,
		}, true)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		logger.Log.Debug().Str("url", msg.URL).Msg("url consumed")
		return
	}(_m)
}

func (s *server) ConsumeAnalyzeResult(_m *st.Msg) {
	go func(m *st.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		logger.Log.Debug().
			Str("value", string(m.Data)).
			Str("subject", "analyze-result").
			Msg("received message")

		msg := &event.AnalyzeResult{}
		err := protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		}.Unmarshal(m.Data, msg)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		compID, err := primitive.ObjectIDFromHex(msg.GetCompanyId())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		techIDs, err := technologyimpl.NewServer().RetrieveTechIDs(ctx, msg.GetResult())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		err = company.SetTechIDs(ctx, compID, techIDs)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		err = m.Ack()
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		if config.Env.LogLevel == "debug" {
			techIDStrs := make([]string, len(techIDs))
			for i, id := range techIDs {
				techIDStrs[i] = id.Hex()
			}
			logger.Log.Debug().
				Str("companyId", msg.GetCompanyId()).
				Strs("techIDs", techIDStrs).
				Msg("analyze-result consumed")
		}
		return
	}(_m)
}

func (s *server) ConsumeImageUploadResult(_m *st.Msg) {
	go func(m *st.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ack := func() {
			err := m.Ack()
			if err != nil {
				logger.Log.Error().Err(err).Send()
			}
		}

		msg := &event.ImageUploadResult{}
		err := protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		}.Unmarshal(m.Data, msg)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			ack()
			return
		}
		newAvatar := msg.GetAvatarUrl()
		if newAvatar == "" || msg.GetCompanyId() == "" {
			ack()
			return
		}

		compID, err := primitive.ObjectIDFromHex(msg.GetCompanyId())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		oldAvatar, err := company.GetAvatar(ctx, compID)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		if newAvatar == oldAvatar {
			ack()
			return
		}

		err = company.SetAvatar(ctx, compID, newAvatar)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		err = stan.ProduceDeleteImage(&event.DeleteImage{
			S3Url: oldAvatar,
		})
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		ack()
		return
	}(_m)
}
