package stan

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-proto/codegen/go/event"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProduceCompanyNew(msg *event.CompanyNew) error {
	b, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}

	return Conn.Publish(config.Env.STAN.SubjectCompanyNew, b)
}
