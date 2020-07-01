package main

import (
	"bytes"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return line, nil
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

	fname := "/Users/denis/Downloads/ru_domains"
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	logger.Must(err)
	defer f.Close()

	for loopAlive {
		lines := make([][]byte, 0)
		for i := 0; i < 100; i += 1 {
			l, err := popLine(f)
			if err != nil {
				logger.Log.Error().Err(err).Msg("error read file next line")
				break
			}
			lines = append(lines, l)
		}
		if len(lines) == 0 {
			break
		}

		wg := sync.WaitGroup{}
		for _, line := range lines {
			wg.Add(1)

			go func(l []byte) {
				defer wg.Done()
				saveLine(l)
			}(line)
		}
		wg.Wait()
	}
}

func saveLine(line []byte) {
	values := strings.Split(string(line), "\t")

	url := strings.ToLower(values[0])
	registrant := strings.ToLower(values[1])
	timeRegistered, err := time.Parse("02.01.2006", values[2])
	logger.Must(err)

	site := model.Site{}
	site.Create(url, registrant, timeRegistered)
}
