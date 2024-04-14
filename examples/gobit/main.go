package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/nokute78/go-bit/v2"
)

type (
	OneByte struct {
		B1 bit.Bit
		B2 bit.Bit
		B3 bit.Bit
		B4 bit.Bit
		B5 bit.Bit
		B6 bit.Bit
		B7 bit.Bit
		B8 bit.Bit
	}

	CompositeBitField struct {
		Length [4]bit.Bit
		Value  [3]bit.Bit
		Flag   bit.Bit
	}
)

func (me OneByte) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v", me.B1, me.B2, me.B3, me.B4, me.B5, me.B6, me.B7, me.B8)
}

func init() {
	log.SetFlags(0)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	/*
	   $ task
	   task: [default] go run main.go
	     1 : 00000001
	     2 : 00000010
	     3 : 00000011
	     7 : 00000111
	    15 : 00001111
	    16 : 00010000
	   127 : 01111111
	   128 : 10000000
	   129 : 10000001
	   255 : 11111111
	   Length= 8([0 0 0 1]), Value=1([1 0 0]), Flag=0(0)
	   Length= 4([0 0 1 0]), Value=3([1 1 0]), Flag=1(1)
	   Length=15([1 1 1 1]), Value=0([0 0 0]), Flag=1(1)
	*/

}

func run() error {
	if err := readOneByte(); err != nil {
		return err
	}

	if err := readComposite(); err != nil {
		return err
	}

	return nil
}

func readOneByte() error {
	var (
		source = []byte{1, 2, 3, 7, 15, 16, 127, 128, 129, 255}
		buf    = bytes.NewBuffer(source)
		data   [10]OneByte
		err    error
	)

	err = bit.Read(buf, binary.BigEndian, &data)
	if err != nil {
		return err
	}

	for i, b := range source {
		log.Printf("%3d : %s", b, data[i])
	}

	return nil
}

func readComposite() error {
	var (
		source = []byte{0x82, 0x47, 0xF1}
		endian = binary.BigEndian
		buf    = bytes.NewBuffer(source)
		data   [3]CompositeBitField
		err    error
	)

	err = bit.Read(buf, endian, &data)
	if err != nil {
		return err
	}

	for _, d := range data {
		var (
			l = bit.BitsToBytes(d.Length[:], endian)[0]
			v = bit.BitsToBytes(d.Value[:], endian)[0]
			r = bit.BitsToBytes([]bit.Bit{d.Flag}, endian)[0]
		)

		log.Printf("Length=%2d(%v), Value=%d(%v), Flag=%d(%v)", l, d.Length, v, d.Value, r, d.Flag)
	}

	return nil
}
