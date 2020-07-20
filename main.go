package main

import (
	graceful "github.com/nnqq/scr-lib-graceful"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/url"
)

func main() {
	urlConsumer := url.NewConsumer()
	go graceful.HandleSignals(urlConsumer.GracefulStop)

	logger.Err(urlConsumer.Serve())
}
