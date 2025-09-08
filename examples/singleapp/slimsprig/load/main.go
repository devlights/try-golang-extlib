package main

import (
	"cmp"
	"fmt"
	"log"
	"maps"
	"slices"
	"text/template"

	sprig "github.com/go-task/slim-sprig/v3"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// slim-sprigの関数マップをロードしてテンプレート処理
	//
	const (
		TEXT = `
{{- .Str | upper | repeat 3}} {{now | date "2006-01-02"}}
{{list "hello" "world" | join "-"}}`
	)
	var (
		m   = sprig.FuncMap()
		t   = template.Must(template.New("my").Funcs(m).Parse(TEXT))
		d   = map[string]any{"Str": "hello"}
		err error
	)
	if err = t.Execute(log.Writer(), d); err != nil {
		return err
	}

	//
	// ライブラリで公開される関数を出力
	//
	log.Println("\n------- funcs")

	const (
		nChunks = 8
	)
	var (
		funcNames = slices.Collect(maps.Keys(m))
		maxWords  = slices.MaxFunc(funcNames, func(x, y string) int {
			return cmp.Compare(len(x), len(y))
		})
	)
	slices.Sort(funcNames)
	{
		for i, name := range funcNames {
			fmt.Fprintf(log.Writer(), "%-*s", len(maxWords)+1, name)

			if (i+1)%nChunks == 0 {
				log.Println("")
			}
		}
	}
	log.Println("")

	return nil

	/*
		$ task
		task: [default] go run .
		HELLOHELLOHELLO 2025-09-08
		hello-world
		------- funcs
		add                        add1                       adler32sum                 ago                        all                        any                        append                     atoi
		b32dec                     b32enc                     b64dec                     b64enc                     base                       biggest                    cat                        ceil
		chunk                      clean                      coalesce                   compact                    concat                     contains                   date                       dateInZone
		dateModify                 date_in_zone               date_modify                deepEqual                  default                    dict                       dig                        dir
		div                        duration                   durationRound              empty                      env                        expandenv                  ext                        fail
		first                      float64                    floor                      fromJson                   get                        getHostByName              has                        hasKey
		hasPrefix                  hasSuffix                  hello                      htmlDate                   htmlDateInZone             indent                     initial                    int
		int64                      isAbs                      join                       keys                       kindIs                     kindOf                     last                       list
		lower                      max                        maxf                       min                        minf                       mod                        mul                        mustAppend
		mustChunk                  mustCompact                mustDateModify             mustFirst                  mustFromJson               mustHas                    mustInitial                mustLast
		mustPrepend                mustPush                   mustRegexFind              mustRegexFindAll           mustRegexMatch             mustRegexReplaceAll        mustRegexReplaceAllLiteral mustRegexSplit
		mustRest                   mustReverse                mustSlice                  mustToDate                 mustToJson                 mustToPrettyJson           mustToRawJson              mustUniq
		mustWithout                must_date_modify           nindent                    now                        omit                       osBase                     osClean                    osDir
		osExt                      osIsAbs                    pick                       pluck                      plural                     prepend                    push                       quote
		randInt                    regexFind                  regexFindAll               regexMatch                 regexQuoteMeta             regexReplaceAll            regexReplaceAllLiteral     regexSplit
		repeat                     replace                    rest                       reverse                    round                      seq                        set                        sha1sum
		sha256sum                  slice                      sortAlpha                  split                      splitList                  splitn                     squote                     sub
		substr                     ternary                    title                      toDate                     toDecimal                  toJson                     toPrettyJson               toRawJson
		toString                   toStrings                  trim                       trimAll                    trimPrefix                 trimSuffix                 trimall                    trunc
		tuple                      typeIs                     typeIsLike                 typeOf                     uniq                       unixEpoch                  unset                      until
		untilStep                  upper                      urlJoin                    urlParse                   values                     without
	*/
}
