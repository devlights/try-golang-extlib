package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/k0kubun/pp/v3"
)

type (
	// Config は、環境変数の値を保持する構造体です.
	//
	// # REFERENCES
	//   - https://github.com/caarlos0/env
	Config struct {
		Home string
		Path []string `envSeparator:":"`
	}
)

func main() {
	if err := run(); err != nil {
		pp.Fatal(err)
	}
}

func run() error {
	var (
		cfg  Config
		opts = env.Options{
			UseFieldNameByDefault: true,
		}
		err error
	)

	// 環境変数の値を取得して値を構造体に設定してもらう
	err = env.ParseWithOptions(&cfg, opts)
	if err != nil {
		return err
	}

	// Pretty Print
	pp.Println(cfg)

	return nil
}
