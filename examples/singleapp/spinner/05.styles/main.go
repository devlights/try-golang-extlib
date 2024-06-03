package main

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	ctx := context.Background()

	for i := 0; i < len(spinner.CharSets); i++ {
		<-play(ctx, i).Done()
	}
}

func play(ctx context.Context, no int) context.Context {
	ctx, cxl := context.WithCancel(ctx)
	go func() {
		defer cxl()

		spin := spinner.New(spinner.CharSets[no], 100*time.Millisecond)

		spin.Color("green")
		spin.Suffix = fmt.Sprintf(" CharSets[%d]", no)
		spin.FinalMSG = fmt.Sprintf("[DONE]%s\n", spin.Suffix)

		spin.Start()
		<-time.After(500 * time.Millisecond)
		spin.Stop()
	}()

	return ctx
}
