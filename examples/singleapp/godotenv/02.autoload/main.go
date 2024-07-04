package main

import (
	"os"

	// autoloadパッケージを暗黙インポートすると自動で .env を読み込んで展開してくれる
	//   "If you're even lazier than that, you can just take advantage of the autoload package which will read in .env on import"
	_ "github.com/joho/godotenv/autoload"
	"github.com/k0kubun/pp/v3"
)

func main() {
	if err := run(); err != nil {
		pp.Fatal(err)
	}
}

func run() error {
	// godotenv.Load() を呼んでないが .env に設定されている環境変数は展開されている
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
