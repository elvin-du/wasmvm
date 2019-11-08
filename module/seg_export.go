package module

import (
	"io"
	"log"
	"wasmvm/leb128"
)

type SegExport struct {
	Num     uint32
	Exports []*Export
}

type Export struct {
	NameLen uint32
	Name    []byte
	Kind    uint32
	Index   uint32
}

func (s *SegExport) EntryID(funcName string) uint32 {
	for _, e := range s.Exports {
		if funcName == string(e.Name) {
			return e.Index
		}
	}

	panic("Function not found" + funcName)
}

func (s *SegExport) Parse(r io.Reader) error {
	num, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		return err
	}
	s.Num = num
	log.Printf("[section export] num:0x%x", num)

	for i := uint32(0); i < num; i++ {
		s.parseExport(r)
	}
	return nil
}

func (s *SegExport) parseExport(r io.Reader) {
	export := &Export{}

	nameLen, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	export.NameLen = nameLen
	log.Printf("[section export][export] nameLen:0x%x", nameLen)

	name, err := readBytes(r, nameLen)
	if nil != err {
		panic(err)
	}
	export.Name = name
	log.Printf("[section export][export] name:0x%x, %s", name, string(name))

	kind, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	export.Kind = kind
	log.Printf("[section export][export] kind:0x%x", kind)

	index, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	export.Index = index
	log.Printf("[section export][export] index:0x%x", index)

	s.Exports = append(s.Exports, export)
}
