package builder

import (
	"github.com/devlights/try-golang-extlib/examples"
	"github.com/devlights/try-golang/mapping"
)

// BuildMappings は、サンプル実行のためのマッピング情報を構築します.
func BuildMappings() mapping.ExampleMapping {
	m := mapping.NewSampleMapping()

	m.MakeMapping(
		examples.NewRegister(),
	)

	return m
}
