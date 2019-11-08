package module

import (
	"io"
	"log"
	"wasmvm/leb128"
)

type SegCode struct {
	Num          uint32
	FunctionBody []*FunctionBody
}

type FunctionBody struct {
	Size uint32
	Code []byte
}

func (s *SegCode) Parse(r io.Reader) error {
	num, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		return err
	}
	s.Num = num
	log.Printf("[section code] num:0x%x", num)

	for i := uint32(0); i < num; i++ {
		s.parseCodeBody(r)
	}
	return nil
}

func (s *SegCode) parseCodeBody(r io.Reader) {
	body := &FunctionBody{}

	size, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	body.Size = size
	log.Printf("[section code][function body] size:0x%x", size)

	code, err := readBytes(r, size)
	if nil != err {
		panic(err)
	}
	body.Code = code
	log.Printf("[section code][function body] code:0x%x", code)

	s.FunctionBody = append(s.FunctionBody, body)
}
