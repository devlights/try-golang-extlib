package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
)

const (
	ENV_FILE = "my.env"
)

func main() {
	if err := run(); err != nil {
		pp.Fatal(err)
	}
}

func run() error {
	// ファイルを指定してロード
	if err := godotenv.Load(ENV_FILE); err != nil {
		return err
	}

	pp.Println(os.Getenv("HELLO"))

	return nil

	/*
		$ task
		task: [build] go build -o app .
		task: [run] echo -n "HELLO=WORLD" > ./my.env
		task: [run] ./app
		"WORLD"
		task: [run] rm -f ./my.env
	*/
}
