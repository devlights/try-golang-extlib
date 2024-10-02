package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/bcicen/jstream"
	"github.com/k0kubun/pp/v3"
)

type (
	Args struct {
		depth int
	}
)

var (
	args Args
)

func init() {
	flag.IntVar(&args.depth, "depth", 0, "depth")
}

func main() {
	flag.Parse()

	if args.depth < 0 {
		args.depth = 0
	}

	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		pp.Fatal(err)
	}
}

func run(_ context.Context) error {
	var (
		file    *os.File
		decoder *jstream.Decoder
		err     error
	)
	file, err = os.Open("data.json")
	if err != nil {
		return fmt.Errorf("os.Open() failed: %w", err)
	}
	defer file.Close()

	decoder = jstream.NewDecoder(file, args.depth)
	for v := range decoder.Stream() {
		pp.Println(v.Value)
	}

	return nil
}
