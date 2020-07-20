package main

import (
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleSignals(stopFunc ...func() error) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	<-signals
	wg := sync.WaitGroup{}
	wg.Add(len(stopFunc))
	for _, f := range stopFunc {
		logger.Err(f())
	}
	wg.Wait()
}

func main() {
	urlConsumer := url.NewConsumer()
	go handleSignals(urlConsumer.GracefulStop)

	logger.Err(urlConsumer.Serve())
}
