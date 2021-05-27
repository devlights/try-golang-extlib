package sets

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// Difference -- Set.Difference() の動作確認のサンプルです。
func Difference() error {

	// Set.Difference の 確認 (差集合)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Difference(s2)

	fmt.Println(s3) // -> "hello"

	return nil
}
