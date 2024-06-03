package main

import (
	"context"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	spin := spinner.New(spinner.CharSets[16], 100*time.Millisecond)

	spin.Color("green")
	spin.Suffix = " NOW LOADING..."
	spin.Start()

	<-work().Done()
	spin.Stop()
}

func work() context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		<-time.After(3 * time.Second)
	}()

	return ctx
}
