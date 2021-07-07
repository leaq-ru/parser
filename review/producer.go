package review

import (
	"context"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-proto/codegen/go/event"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProduceModeration(ctx context.Context, review *parser.ReviewItem) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	b, err := protojson.Marshal(&event.ReviewModeration{
		Review: review,
	})
	if err != nil {
		return err
	}

	return stan.Conn.Publish(config.Env.STAN.SubjectReviewModeration, b)
}
