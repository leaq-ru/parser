package iterator

import (
	"context"
	"errors"
	"fmt"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

const linesInParallel = 50

type offset struct {
	Index int `bson:"index"`
}

var loopAlive = true

func FileParse(localPath string) {
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, os.Kill)
	go func() {
		<-exitCh
		loopAlive = false
		logger.Log.Info().Bool("loopAlive", loopAlive).Msg("waiting for last iteration and exit")
	}()

	fileBytes, err := ioutil.ReadFile(localPath)
	logger.Must(err)

	file := strings.Split(string(fileBytes), "\n")

	for loopAlive {
		o := offset{}
		err := mongo.FileOffset.FindOne(context.Background(), bson.D{}).Decode(&o)
		if err != nil && !errors.Is(err, mongo2.ErrNoDocuments) {
			logger.Log.Panic().Err(err).Send()
		}

		lines := make([]string, 0)
		for i := 0; i < linesInParallel; i += 1 {
			lines = append(lines, file[o.Index+i])
		}

		if len(lines) == 0 {
			break
		}

		opts := options.UpdateOptions{}
		opts.SetUpsert(true)
		_, err = mongo.FileOffset.UpdateOne(context.Background(), bson.D{}, bson.M{
			"$inc": bson.M{
				"index": len(lines),
			},
		}, &opts)
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
	defer func() {
		r := recover()
		if r != nil {
			logger.Log.Info().Msg("recovered")
			fmt.Println(r)
			loopAlive = false
			logger.Log.Info().Bool("loopAlive", loopAlive).Msg("waiting for last iteration and exit")
		}
	}()

	values := strings.Split(line, "\t")

	url := strings.ToLower(values[0])
	registrant := strings.ToLower(values[1])
	timeRegistered, err := time.Parse("02.01.2006", values[2])
	logger.Must(err)

	company := model.Company{}
	company.Upsert(url, registrant, timeRegistered)
}
