package main

import (
	"context"
	"os"

	"github.com/beevik/etree"
)

var (
	langs = map[string]string{
		"golang": "https://go.dev/",
		"csharp": "https://learn.microsoft.com/ja-jp/dotnet/csharp/",
		"java":   "https://docs.oracle.com/javase/jp/21/index.html",
		"python": "https://docs.python.org/ja/3/",
		"rust":   "https://www.rust-lang.org/ja/learn",
	}
)

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(_ context.Context) error {
	//
	// etreeライブラリは、PythonのElementTreeライブラリに
	// インスパイアされて作成されているXML処理ライブラリ。
	//
	// 使いやすいAPIが備わっているので直感的に作業できる。
	//

	var (
		doc       *etree.Document
		languages *etree.Element
		language  *etree.Element
	)
	doc = etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	languages = doc.CreateElement("languages")
	for lang, url := range langs {
		language = languages.CreateElement("language")
		language.CreateAttr("name", lang)
		language.SetText(url)
	}

	doc.Indent(4)
	if _, err := doc.WriteTo(os.Stdout); err != nil {
		return err
	}

	return nil
}
