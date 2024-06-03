# これは何？

[flaggy](https://github.com/integrii/flaggy) パッケージの使い方をメモしたものです。

サブコマンドはネストさせることも出来ます。

```go
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
```

上記のような構成にすると、コマンドラインにて以下のように指定できるようになります。

```sh
$ app subA subsubA -v1 helloworld
```

```docker compose ls``` のような形のコマンドラインが作れるということです。

## 参考情報

- https://github.com/integrii/flaggy
- https://ericgreer.info/post/a-better-flags-package-for-go/
