package main

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

const linesInParallel = 500

type offset struct {
	Index int `bson:"index"`
}

func oneTimeFileParse() {
	loopAlive := true

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, os.Kill)
	go func() {
		<-exitCh
		loopAlive = false
		logger.Log.Info().Bool("loopAlive", loopAlive).Msg("waiting for last iteration and exit")
	}()

	fileBytes, err := ioutil.ReadFile("/Users/denis/Downloads/ru_domains")
	logger.Must(err)

	file := strings.Split(string(fileBytes), "\n")

	for loopAlive {
		o := offset{}
		err := mongo.DebugFileOffset.FindOne(context.Background(), bson.D{}).Decode(&o)
		logger.Must(err)

		lines := make([]string, 0)
		for i := 0; i < linesInParallel; i += 1 {
			lines = append(lines, file[o.Index+i])
		}

		if len(lines) == 0 {
			break
		}

		_, err = mongo.DebugFileOffset.UpdateOne(context.Background(), bson.D{}, bson.M{
			"$inc": bson.M{
				"index": len(lines),
			},
		})
		logger.Must(err)

		wg := sync.WaitGroup{}
		for _, line := range lines {
			wg.Add(1)

			go func(l string) {
				defer wg.Done()
				saveLine(l)
			}(line)
		}
		wg.Wait()
	}
}

func saveLine(line string) {
	values := strings.Split(line, "\t")

	url := strings.ToLower(values[0])
	registrant := strings.ToLower(values[1])
	timeRegistered, err := time.Parse("02.01.2006", values[2])
	logger.Must(err)

	site := model.Site{}
	site.Create(url, registrant, timeRegistered)
}
