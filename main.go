package main

import (
	"context"
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

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := stan.Conn.Close()
		if err != nil {
			logger.Log.Error().Err(err).Send()
		}
	}()

	go func() {
		defer wg.Done()
		err := mongo.DB.Client().Disconnect(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Send()
		}
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
