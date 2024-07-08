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
	var (
		myEnv map[string]string
		err   error
	)

	// 環境変数を弄られたくない場合は、godotenv.Read()で内容を取得できる
	myEnv, err = godotenv.Read()
	if err != nil {
		return err
	}

	pp.Println(myEnv["HELLO"])
	if _, ok := os.LookupEnv("HELLO"); !ok {
		pp.Println("環境変数側にHELLOは存在しない")
	}

	return nil

	/*
		$ task
		task: [build] go build -o app .
		task: [run] echo -n "HELLO=WORLD" > ./.env
		task: [run] ./app
		"WORLD"
		"環境変数側にHELLOは存在しない"
		task: [run] rm -f ./.env
	*/
}
