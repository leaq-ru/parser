package vk

import (
	"time"
)

type tokenPool struct {
	tokens chan string
	maxRPS int
	curRPS int
}

func newTokenPool(perTokenLimit int, tokens ...string) *tokenPool {
	c := make(chan string, len(tokens))
	for _, t := range tokens {
		c <- t
	}

	tp := &tokenPool{
		tokens: c,
		maxRPS: len(tokens) * perTokenLimit,
	}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			<-t.C
			if tp.curRPS > 0 {
				tp.curRPS -= 1
			}
		}
	}()

	return tp
}

func (tp *tokenPool) Get() string {
	if tp.curRPS >= tp.maxRPS {
		time.Sleep(time.Second)
		return tp.Get()
	}

	token := <-tp.tokens
	tp.tokens <- token

	return token
}
