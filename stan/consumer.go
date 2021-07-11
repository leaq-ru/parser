package stan

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/rs/zerolog"
	"time"
)

type consumer struct {
	logger      zerolog.Logger
	stan        stan.Conn
	subject     string
	sub         stan.Subscription
	maxInFlight int
	cb          cb
}

type cb = func(msg *stan.Msg)

func NewConsumer(logger zerolog.Logger, stan stan.Conn, subject string, maxInFlight int, cb cb) (*consumer, error) {
	c := &consumer{
		logger:      logger.With().Str("package", "consumer").Str("subject", subject).Logger(),
		stan:        stan,
		subject:     subject,
		maxInFlight: maxInFlight,
		cb:          cb,
	}
	err := c.subscribe()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *consumer) Serve(ctx context.Context) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if !c.sub.IsValid() {
				err := c.subscribe()
				if err != nil {
					c.logger.Error().Err(err).Send()
				}
			}
		}
	}
}

func (c *consumer) subscribe() error {
	sub, err := c.stan.Subscribe(
		c.subject,
		c.cb,
		stan.SetManualAckMode(),
		stan.DurableName(c.subject),
		stan.MaxInflight(c.maxInFlight),
	)
	if err != nil {
		c.logger.Error().Err(err).Send()
		return err
	}

	c.sub = sub
	return nil
}
