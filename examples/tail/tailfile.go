package tail

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	tailpkg "github.com/nxadm/tail"
	"golang.org/x/sync/errgroup"
)

// TailFile は、github.com/nxadm/tail のサンプルです。
//
// # REFERENCES
//
//   - https://github.com/nxadm/tail
func TailFile() error {
	const (
		FILE_PATH = "messages.log"
	)

	log.SetFlags(0)

	//
	// 対象となるファイルを作成
	//
	var (
		file *os.File
		err  error
	)

	file, err = os.Create(FILE_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	//
	// 1秒ごとにデータを書き込み
	//
	var (
		rootCtx            = context.Background()
		mainCtx, mainCxl   = context.WithCancel(rootCtx)
		g, gCtx            = errgroup.WithContext(mainCtx)
		writeCtx, writeCxl = context.WithTimeout(mainCtx, 5300*time.Millisecond)
	)
	defer mainCxl()

	g.Go(func() error {
		defer writeCxl()

		for {
			select {
			case <-writeCtx.Done():
				return nil
			case <-gCtx.Done():
				return nil
			case t := <-time.After(1 * time.Second):
				_, err = file.WriteString(t.Format(time.TimeOnly) + " helloworld\n")
				if err != nil {
					return err
				}
			}
		}
	})

	//
	// ファイルをtailする
	//
	var (
		lines = make(chan *tailpkg.Line, 128)
		t     *tailpkg.Tail
		c     = tailpkg.Config{
			Location: &tailpkg.SeekInfo{
				Offset: 0,
				Whence: io.SeekEnd,
			},
			Follow: true,
			ReOpen: true,
		}
	)

	t, err = tailpkg.TailFile(FILE_PATH, c)
	if err != nil {
		return err
	}

	g.Go(func() error {
		stop := func() error {
			if err := t.Stop(); err != nil {
				return err
			}

			return nil
		}

		for {
			select {
			case <-writeCtx.Done():
				return stop()
			case <-gCtx.Done():
				return stop()
			case line, ok := <-t.Lines:
				if !ok {
					return nil
				}

				lines <- line
			}
		}
	})

	// 出力
LOOP:
	for {
		select {
		case <-t.Dead():
			break LOOP
		case line, ok := <-lines:
			if !ok {
				break LOOP
			}

			log.Printf("[tail] %s", line.Text)
		}
	}

	return nil
}
