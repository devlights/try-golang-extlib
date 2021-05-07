package builder

import (
	"github.com/devlights/try-golang/pkg/mappings"
)

// BuildMappings は、サンプル実行のためのマッピング情報を構築します.
func BuildMappings() mappings.ExampleMapping {
	m := mappings.NewSampleMapping()

	m.MakeMapping(
	)

	return m
}