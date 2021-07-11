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
	qGroup      string
	sub         stan.Subscription
	maxInFlight int
	cb          cb
}

type cb = func(msg *stan.Msg)

func NewConsumer(logger zerolog.Logger, stan stan.Conn, subject, qGroup string, maxInFlight int, cb cb) (*consumer, error) {
	c := &consumer{
		logger:      logger.With().Str("package", "consumer").Str("subject", subject).Logger(),
		stan:        stan,
		subject:     subject,
		qGroup:      qGroup,
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
	opts := []stan.SubscriptionOption{
		stan.SetManualAckMode(),
		stan.DurableName(c.subject),
	}
	if c.maxInFlight != 0 {
		opts = append(opts, stan.MaxInflight(c.maxInFlight))
	}

	sub, err := c.stan.QueueSubscribe(
		c.subject,
		c.qGroup,
		c.cb,
		opts...,
	)
	if err != nil {
		c.logger.Error().Err(err).Send()
		return err
	}

	c.sub = sub
	return nil
}
