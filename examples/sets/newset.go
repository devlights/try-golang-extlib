package sets

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// NewSet - mapset.NewSet() の サンプル
// REFERENCES:: https://github.com/deckarep/golang-set
func NewSet() error {

	// 新しい 集合 を生成
	s1 := mapset.NewSet()

	// データを設定
	s1.Add("hello")
	s1.Add("world")
	s1.Add("hello")

	// 集合なので重複項目は存在しない
	fmt.Println(s1) // -> "hello","world"

	return nil
}
