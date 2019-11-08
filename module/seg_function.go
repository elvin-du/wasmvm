package module

import (
	"io"
	"log"
	"wasmvm/leb128"
)

type SegFunction struct {
	Num        uint32
	SigIndexes []uint32
}

func (s *SegFunction) Parse(r io.Reader) error {
	num, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		return err
	}
	s.Num = num
	log.Printf("[section function] num:0x%x", num)

	for i := uint32(0); i < num; i++ {
		s.parseSigIndex(r)
	}
	return nil
}

func (s *SegFunction) parseSigIndex(r io.Reader) {
	index, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	log.Printf("[section function][signature index] index:0x%x", index)

	s.SigIndexes = append(s.SigIndexes, index)
}
