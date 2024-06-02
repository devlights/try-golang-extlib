package main

import (
	"log"

	"github.com/integrii/flaggy"
)

type Args struct {
	Val1 string
	Val2 int
	Val3 bool
}

var (
	args Args
)

func init() {
	// ルートレベルのフラグを追加
	flaggy.String(&args.Val1, "v1", "val1", "value 1")
	flaggy.Int(&args.Val2, "", "val2", "value 2")
	flaggy.Bool(&args.Val3, "v3", "", "value 3")

	// ヘルプに表示される情報やバージョン情報など
	flaggy.SetName("app")
	flaggy.SetDescription("flaggy sample app")
	flaggy.SetVersion("v1.0.0")

	flaggy.Parse()
}

func main() {
	log.SetFlags(0)
	log.Printf("%+v", args)
}
