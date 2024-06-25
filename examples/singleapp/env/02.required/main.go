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
		HomeDir string `env:"HOME"`
		MyEnv1  string `env:"MYENV1,required"`
		MyEnv2  int    `env:"MYENV2,required"`
	}
)

func main() {
	if err := run(); err != nil {
		pp.Fatal(err)
	}
}

func run() error {
	var (
		cfg Config
		err error
	)

	// 環境変数の値を取得して値を構造体に設定してもらう
	err = env.Parse(&cfg)
	if err != nil {
		pp.Println(err)
	}

	pp.Println(cfg)

	return nil
}
