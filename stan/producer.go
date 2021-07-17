package stan

import (
	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/proto/codegen/go/event"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProduceCompanyNew(msg *event.CompanyNew) error {
	b, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}

	return Conn.Publish(config.Env.STAN.SubjectCompanyNew, b)
}

func ProduceDeleteImage(msg *event.DeleteImage) error {
	b, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}

	return Conn.Publish(config.Env.STAN.SubjectDeleteImage, b)
}
