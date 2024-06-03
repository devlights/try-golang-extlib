package main

import (
	"log"

	"github.com/integrii/flaggy"
)

type (
	SubAArgs struct {
		SubSubA SubSubAArgs
	}

	SubSubAArgs struct {
		Val1 string
	}

	Args struct {
		Val1 int
		SubA SubAArgs
	}
)

var (
	args    Args
	subA    *flaggy.Subcommand
	subsubA *flaggy.Subcommand
)

func init() {
	// ルートレベルのフラグ
	flaggy.Int(&args.Val1, "d", "val", "value 1")

	// サブコマンドの作成
	subA = flaggy.NewSubcommand("subA")

	// サブサブコマンドの作成
	subsubA = flaggy.NewSubcommand("subsubA")
	subsubA.String(&args.SubA.SubSubA.Val1, "v1", "val1", "subsubA value 1")

	// アタッチ
	subA.AttachSubcommand(subsubA, 1)
	flaggy.AttachSubcommand(subA, 1)

	// ヘルプに表示される情報やバージョン情報など
	subA.Description = "Subcommand A"
	subsubA.Description = "Nested-Subcommand A"

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
		switch {
		case subsubA.Used:
			log.Println("subsubA Used")
		default:
			log.Println("invalid nested-sub-command")
		}
	default:
		log.Println("No subcommand")
	}

	log.Printf("%#v", args)
}
