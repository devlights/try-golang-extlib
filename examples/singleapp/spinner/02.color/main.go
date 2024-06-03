package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	var (
		spin = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	)
	defer spin.Stop()

	spin.Color("green")
	spin.Start()

	<-time.After(3 * time.Second)
}
