package sets

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// Union -- Set.Union() の 動作確認のサンプルです。
func Union() error {

	// Set.Union の 確認 (和集合)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Union(s2)

	fmt.Println(s3) // -> "hello","world","golang","python"

	return nil
}
