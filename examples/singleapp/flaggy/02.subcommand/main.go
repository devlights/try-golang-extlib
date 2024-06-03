package main

import (
	"log"

	"github.com/integrii/flaggy"
)

var (
	args Args
	subA *flaggy.Subcommand
	subB *flaggy.Subcommand
)

func init() {
	// ルートレベルのフラグ
	flaggy.Int(&args.Val1, "d", "val", "value 1")

	// サブコマンドの作成
	subA = flaggy.NewSubcommand("subA")
	subA.Int(&args.SubA.SubVal1, "v1", "val1", "subA value 1")
	subB = flaggy.NewSubcommand("subB")
	subB.String(&args.SubB.SubVal1, "v1", "val1", "subB value 1")

	// サブコマンドをアタッチ
	// 第２引数は relativePosition となっていて
	// このサブコマンドが置かれる位置を指定することになっているが
	// 基本は１で良い
	flaggy.AttachSubcommand(subA, 1)
	flaggy.AttachSubcommand(subB, 1)

	// ヘルプに表示される情報やバージョン情報など
	subA.Description = "Subcommand A"
	subB.Description = "Subcommand B"

	flaggy.SetName("app")
	flaggy.SetDescription("flaggy sample app")
	flaggy.SetVersion("v1.0.0")

	flaggy.Parse()
}

func main() {
	log.SetFlags(0)

	// どのサブコマンドが指定されたかは *flaggy.Subcommand.Used でわかる
	switch {
	case subA.Used:
		log.Println("SubA used")
	case subB.Used:
		log.Println("SubB used")
	default:
		log.Println("No subcommand")
	}

	log.Printf("%#v", args)
}
