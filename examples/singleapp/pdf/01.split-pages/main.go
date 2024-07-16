package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type (
	Args struct {
		inFile string
		outDir string
	}
)

var (
	args Args
)

func init() {
	log.SetFlags(log.Lmicroseconds)

	flag.StringVar(&args.inFile, "in", "", "input file")
	flag.StringVar(&args.outDir, "out", "", "output dir")
}

func main() {
	flag.Parse()

	if args.inFile == "" || args.outDir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := api.SplitFile(args.inFile, args.outDir, 1, nil); err != nil {
		return fmt.Errorf("error splitting PDF: %w", err)
	}

	return nil
}
