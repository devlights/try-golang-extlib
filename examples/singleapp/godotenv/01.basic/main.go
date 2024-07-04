package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
)

func main() {
	if err := run(); err != nil {
		pp.Fatal(err)
	}
}

func run() error {
	// 引数なしで Load() を呼ぶと .env が読み込まれる
	if err := godotenv.Load(); err != nil {
		return err
	}

	pp.Println(os.Getenv("HELLO"))

	return nil

	/*
	   $ task
	   task: [build] go build -o app .
	   task: [run] echo -n "HELLO=WORLD" > ./.env
	   task: [run] ./app
	   "WORLD"
	   task: [run] rm -f ./.env
	*/
}
