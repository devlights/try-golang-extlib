package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bojanz/currency"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, time.Second, errors.New("too slow"))
		procCtx          = run(mainCtx)
		err              error
	)
	defer mainCxl()

	select {
	case <-mainCtx.Done():
		err = context.Cause(mainCtx)
	case <-procCtx.Done():
		if err = context.Cause(procCtx); errors.Is(err, context.Canceled) {
			err = nil
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) context.Context {
	ctx, cxl := context.WithCancelCause(ctx)

	go func() {
		cxl(proc(ctx))
	}()

	return ctx
}

func proc(_ context.Context) error {
	type data struct {
		v string
		c string
		l string
	}

	items := []data{
		{"129", "JPY", "ja"},
		{"129", "USD", "en"},
		{"129", "EUR", "fr"},
	}
	for _, item := range items {
		var (
			amount    currency.Amount
			locale    currency.Locale
			formatter *currency.Formatter
			err       error
		)
		locale = currency.NewLocale(item.l)
		formatter = currency.NewFormatter(locale)

		amount, err = currency.NewAmount(item.v, item.c)
		if err != nil {
			return err
		}

		fmt.Println(formatter.Format(amount))
	}

	return nil
}
