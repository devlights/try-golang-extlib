package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/bmatcuk/doublestar/v4"
)

func main() {
	const (
		BaseDir = "../../../../" // try-golang-extlib project root-dir
	)
	var (
		basepath string
		pattern  = filepath.Join(BaseDir, "**/main.go")
	)
	basepath, pattern = doublestar.SplitPattern(pattern)
	fmt.Printf("base: %s, pattern: %s\n", basepath, pattern)

	var (
		dirFs   = os.DirFS(basepath)
		matches []string
		err     error
	)
	matches, err = doublestar.Glob(dirFs, pattern)
	if err != nil {
		panic(err)
	}

	fmt.Printf("match count: %d\n", len(matches))
	for v := range slices.Values(matches) {
		fmt.Printf("\t%s\n", v)
	}
}
