package module

import (
	"fmt"
	"io"
)

type SegmentTable struct {
	Table map[int8]ISegment
}

var globalSegmentTable = &SegmentTable{
	Table: map[int8]ISegment{
		//0: "",         //Custom segment
		1: &SegType{Entries: []*FuncType{}}, //Type segment
		//2:  "",         //Import segment
		3: &SegFunction{SigIndexes: []uint32{}}, //Function segment
		//4:  "",         //Table segment
		//5:  "",         //Memory segment
		//6:  "",         //Global segment
		7: &SegExport{Exports: []*Export{}}, //Export segment
		//8:  "",         //Start segment
		//9:  "",         //Element segment
		10: &SegCode{FunctionBody: []*FunctionBody{}}, //Code segment
		//11: "",         //Data segment
	},
}

func (s *SegmentTable) Segment(id int8) ISegment {
	seg, ok := s.Table[id]
	if !ok {
		panic(fmt.Sprintf("Invalid Section ID:0x%x", id))
	}
	return seg
}

type ISegment interface {
	Parse(io.Reader) error
	//String()
}

//type Segment struct {
//	ID          uint8
//	PayLoadLen  uint32
//	NameLen     uint32
//	Name        []byte
//	PayLoadData []byte
//}
