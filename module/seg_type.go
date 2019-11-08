package module

import (
	"fmt"
	"io"
	"log"
	"wasmvm/leb128"
)

type SegType struct {
	Count   uint32
	Entries []*FuncType
}

type FuncType struct {
	Form        int8 //fixed: 0x60
	ParamCount  uint32
	ParamTypes  []int8
	ReturnCount uint8
	ReturnType  int8
}

func (s *SegType) Parse(r io.Reader) error {
	count, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		return err
	}
	s.Count = count
	log.Printf("[section type] count:0x%x", count)

	for i := uint32(0); i < count; i++ {
		s.parseFuncType(r)
	}
	return nil
}

func (s *SegType) parseFuncType(r io.Reader) {
	funcType := &FuncType{}

	form, _, err := leb128.ReadVarUint32Size(r)
	if FunctionType != int8(form) {
		panic(fmt.Sprintf("expect 0x60,got %x", form))
	}
	funcType.Form = int8(form)
	log.Printf("[section type][func type] form:0x%x", form)

	paramCount, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	funcType.ParamCount = paramCount
	log.Printf("[section type][func type] paramCount:0x%x", paramCount)

	paramTypes := []int8{}
	for i := uint32(0); i < paramCount; i++ {
		valType, _, err := leb128.ReadVarUint32Size(r)
		if nil != err {
			panic(err)
		}
		log.Printf("[section type][func type] paramType:0x%x", valType)

		paramTypes = append(paramTypes, int8(valType))
	}
	funcType.ParamTypes = paramTypes

	retCount, _, err := leb128.ReadVarUint32Size(r)
	if nil != err {
		panic(err)
	}
	log.Printf("[section type][func type] returnCount:0x%x", retCount)
	funcType.ReturnCount = uint8(retCount)

	for i := uint8(0); i < uint8(retCount); i++ {
		retType, _, err := leb128.ReadVarUint32Size(r)
		if nil != err {
			panic(err)
		}
		log.Printf("[section type][func type] returnType:0x%x", retType)

		funcType.ReturnType = int8(retType)
	}

	s.Entries = append(s.Entries, funcType)
}

func (s *SegType) String() {
	log.Printf(`section type count:0x%x`, s.Count)
	for _, e := range s.Entries {
		log.Printf("form:%x", e.Form)
		log.Printf("ParamCount:0x%x", e.ParamCount)
		for _, e2 := range e.ParamTypes {
			log.Printf("ParamTypes:0x%x", e2)
		}
		log.Printf("ReturnCount:0x%x", e.ReturnCount)
		log.Printf("ReturnType:0x%x", e.ReturnType)
	}
}
