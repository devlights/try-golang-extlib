# これは何？

[flaggy](https://github.com/integrii/flaggy) パッケージの使い方をメモしたものです。

[flaggy](https://github.com/integrii/flaggy) パッケージは、サブコマンドも扱うことが可能なシンプルなフラグライブラリです。

以下のライブラリであると記載されています。以下、引用。

> Sensible and fast command-line flag parsing with excellent support for subcommands and positional values. Flags can be at any position. Flaggy has no required project or package layout like Cobra requires, and no external dependencies!

> (サブコマンドと位置引数の優れたサポートで、コマンドラインのフラグを賢明かつ高速に解析します。フラグはどの位置にも置くことができます。Flaggyは、Cobraが必要とするようなプロジェクトやパッケージのレイアウトを必要とせず、外部依存関係もありません！)

ライブラリが必要とする依存関係がとても少ないので、導入し易いです。

[flaggy](https://github.com/integrii/flaggy) を使って、標準の ```flag``` パッケージと同じような定義をするには以下のようにします。


```go
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
```

デフォルトで、```flag``` パッケージでは出来ない ロングバージョン の引数が定義出来るのでとても便利。 (```-v1``` と ```--val1```のように)

ロングバージョンの引数のみとしたい場合はショートバージョンの引数の値を空文字にすれば良いです。

```flaggy.Parse()``` で解析に失敗したら内部で ```log.Panicln``` が呼ばれます。

本サンプルを実行すると以下のようになります。

```sh
$ task
task: [default] go run . help
app - flaggy sample app

  Flags: 
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -v1 --val1      value 1
       --val2      value 2 (default: 0)
    -v3           value 3

task: [default] go run . version
Version: v1.0.0
task: [default] go run . -v1 helloworld --val2 999 -v3
{Val1:helloworld Val2:999 Val3:true}
task: [default] go run . --val1 hello --val2 888
{Val1:hello Val2:888 Val3:false}
```

## 参考情報

- https://github.com/integrii/flaggy
- https://ericgreer.info/post/a-better-flags-package-for-go/
