package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/beevik/etree"
)

var (
	logger *slog.Logger
)

func init() {
	var (
		opt = &slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch {
				case a.Key == slog.LevelKey:
					return slog.Attr{}
				case a.Key == slog.TimeKey:
					return slog.Attr{}
				case a.Key == slog.MessageKey:
					return slog.Attr{Key: "tag", Value: a.Value}
				default:
					return a
				}
			},
		}
		handler = slog.NewJSONHandler(os.Stdout, opt)
	)

	logger = slog.New(handler)
}

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(_ context.Context) error {
	//
	// etreeライブラリは、PythonのElementTreeライブラリに
	// インスパイアされて作成されているXML処理ライブラリ。
	//
	// 使いやすいAPIが備わっているので直感的に作業できる。
	//

	var (
		doc *etree.Document
		err error
	)
	doc = etree.NewDocument()
	if err = doc.ReadFromFile("langs.xml"); err != nil {
		return err
	}

	var (
		root      = doc.SelectElement("languages")
		languages = root.SelectElements("language")
	)
	for _, lang := range languages {
		var (
			attrs = make([]any, 0)
		)

		for _, a := range lang.Attr {
			attrs = append(attrs, slog.Attr{Key: a.Key, Value: slog.StringValue(a.Value)})
		}

		attrs = append(attrs, slog.Attr{Key: "text", Value: slog.StringValue(lang.Text())})

		logger.Info(lang.Tag, attrs...)
	}

	return nil
}
