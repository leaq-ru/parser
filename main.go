package main

import (
	"context"
	"github.com/nnqq/scr-parser/iterator"
)

func main() {
	iterator.FileParse(context.Background(), "/Users/denis/Downloads/ru_domains")
}
