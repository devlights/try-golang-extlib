package clipboard

import (
	"bytes"
	"fmt"
	"io"

	cb "github.com/skanehira/clipboard-image/v2"
)

func ReadWrite() error {
	var (
		buf = bytes.NewBufferString("hello world")
		err error
	)

	// クリップボードに書き込み
	err = cb.Write(buf)
	if err != nil {
		return err
	}

	// クリップボードから読み込み
	var (
		r io.Reader
	)

	r, err = cb.Read()
	if err != nil {
		return err
	}

	_, err = io.Copy(buf, r)
	if err != nil {
		return err
	}

	fmt.Println(buf.String())

	return nil
}
