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
		Home string `env:"HOME"`
		Dir1 string `env:"DIR1,required"` // 存在するかどうかだけが判断され、値が空の場合はOKとなる
		Dir2 string `env:"DIR2,notempty"` // 値が空の場合にエラーとなる
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
		opts env.Options
		err  error
	)

	// 環境変数の値を取得して値を構造体に設定してもらう
	err = env.ParseWithOptions(&cfg, opts)
	if err != nil {
		pp.Println(err)
	}

	// Pretty Print
	pp.Println(cfg)

	return nil
}
