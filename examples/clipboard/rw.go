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
		return fmt.Errorf("clipboard.Write() fariled: %w", err)
	}

	// クリップボードから読み込み
	var (
		r io.Reader
	)

	r, err = cb.Read()
	if err != nil {
		return fmt.Errorf("clipboard.Read() fariled: %w", err)
	}

	_, err = io.Copy(buf, r)
	if err != nil {
		return fmt.Errorf("io.Copy() fariled: %w", err)
	}

	fmt.Println(buf.String())

	return nil
}
