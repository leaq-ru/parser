package main

import (
	"github.com/nnqq/scr-parser/iterator"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	iterator.FileParse("/Users/denis/Downloads/ru_domains")
}
