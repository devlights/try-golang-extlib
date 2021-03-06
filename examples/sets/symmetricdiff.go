package sets

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// SymmetricDiff -- Set.SymmetricDifference() の動作確認のサンプルです。
func SymmetricDiff() error {

	// Set.SymmetricDifference() の 確認 (対象差)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.SymmetricDifference(s2)

	fmt.Println(s3) // -> "hello","golang","python"

	return nil
}
