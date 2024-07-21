package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/bodgit/sevenzip"
)

type (
	Args struct {
		inFile string
	}
)

var (
	args Args
)

func init() {
	log.SetFlags(0)

	flag.StringVar(&args.inFile, "in", "", "input file")
	flag.Parse()
}

func main() {
	if args.inFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		reader *sevenzip.ReadCloser
		err    error
	)
	if reader, err = sevenzip.OpenReader(args.inFile); err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		err = func(file *sevenzip.File) error {
			var (
				in  io.ReadCloser
				out bytes.Buffer
			)
			if in, err = file.Open(); err != nil {
				return err
			}
			defer in.Close()

			if _, err = io.Copy(&out, in); err != nil {
				return err
			}

			log.Printf("[FILE   ] %s", file.Name)
			log.Printf("[CRC32  ] %x", file.CRC32)
			log.Printf("[CONTENT] %d byte(s)", len(out.Bytes()))

			return nil
		}(file)

		if err != nil {
			return err
		}
	}

	return nil
}
