package main

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/consumer"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/stan"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	e := func(err error) {
		if err != nil {
			logger.Log.Error().Err(err).Send()
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		e(stan.Conn.Close())
	}()

	go func() {
		defer wg.Done()
		e(mongo.DB.Client().Disconnect(ctx))
	}()

	go func() {
		defer wg.Done()
		e(call.GrpcConn.Close())
	}()
	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{}, 1)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-signals
		cancel()
		cleanup()
		done <- struct{}{}
	}()

	err := consumer.URL(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}

	<-done
}
