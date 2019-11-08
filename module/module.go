package module

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"wasmvm/leb128"
)

var (
	WasmMagicNumber = []byte{0x00, 0x61, 0x73, 0x6d}
	WasmMinLength   = 4
)

type Module struct {
	Code     []byte
	Segments map[int8]ISegment
}

func NewModule() *Module {
	return &Module{Segments: make(map[int8]ISegment)}
}

func (m *Module) Init(filename string) {
	bin, err := ioutil.ReadFile(filename)
	if nil != err {
		panic(err)
	}

	m.CheckMagicNumber(bin)
	m.CheckVerion(bin)

	if nil == m.Code {
		m.Code = bin
	}
}

func (m *Module) ParseSegments() {
	size := uint(len(m.Code))
	log.Printf("code length:0x%x", size)
	pc := uint(8)
	r := bytes.NewReader(m.Code[pc:])

	for {
		id, _, err := leb128.ReadVarUint32Size(r)
		if nil != err && io.EOF == err {
			return
		}
		segID := int8(id)
		log.Printf("[section] id:0x%x", segID)

		sectionSize, _, err := leb128.ReadVarUint32Size(r)
		if nil != err {
			panic(err)
		}
		log.Printf("[section] size:0x%x", sectionSize)

		segment := globalSegmentTable.Segment(segID)
		segment.Parse(r)

		m.Segments[segID] = segment
		//segment.String()
	}
}

func (m *Module) CheckVerion(bin []byte) {
	if 0x1 != binary.LittleEndian.Uint32(bin[4:8]) {
		panic("invalid wasm format: vesion != 0x1")
	}
}

func (m *Module) CheckMagicNumber(bin []byte) {
	if len(bin) < WasmMinLength {
		panic("invalid wasm format: length < 4")
	}

	if bytes.Compare(WasmMagicNumber, bin[:4]) != 0 {
		panic("invalid wasm format: first 4 bytes  != 0asm")
	}
}
