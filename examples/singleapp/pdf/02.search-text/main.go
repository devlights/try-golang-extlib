package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/ledongthuc/pdf"
)

type (
	Args struct {
		inFile string
		word   string
	}
)

var (
	args Args
)

func init() {
	log.SetFlags(log.Lmicroseconds)

	flag.StringVar(&args.inFile, "in", "", "input file")
	flag.StringVar(&args.word, "word", "", "search word")
}

func main() {
	flag.Parse()

	if args.inFile == "" || args.word == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		file *os.File
		pdfR *pdf.Reader
		err  error
	)

	//
	// 暗号化されているPDFを開こうとすると
	//   malformed PDF: 256-bit encryption key
	// というエラーが出る
	//
	// 01.split-pages で利用した IPA の「安全なウェブサイトの作り方」のPDFは
	// ページ分割は出来るが、テキストを抽出しようとすると
	//   malformed PDF: 256-bit encryption key
	// と出て無理だった.
	//
	// サンプルなので、自前で適当に作成したPDFファイルで試す
	//
	file, pdfR, err = pdf.Open(args.inFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		buf bytes.Buffer
		r   io.Reader
	)
	r, err = pdfR.GetPlainText()
	if err != nil {
		return err
	}
	io.Copy(&buf, r)

	var (
		data = buf.Bytes()
		sep  = []byte(args.word)
		idx  = bytes.Index(data, sep)
	)

	if idx < 0 {
		log.Println("no hit.")
		return nil
	}

	log.Printf("hit: %d,%s", idx, data[idx:idx+len(sep)])

	return nil
}
