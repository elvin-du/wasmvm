package module

const (
	ValueTypeI32 int8 = 0x7F
	ValueTypeI64 int8 = 0x7E
	ValueTypeF32 int8 = 0x7D
	ValueTypeF64 int8 = 0x7C
)

const (
	ElementType  int8 = 0x70
	FunctionType int8 = 0x60
	ResultType   int8 = 0x40
)

const (
	SegmentIDCustom  int8 = 0x00
	SegmentIDType    int8 = 0x01
	SegmentIDImport  int8 = 0x02
	SegmentIDFuction int8 = 0x03
	SegmentIDTable   int8 = 0x04
	SegmentIDMemory  int8 = 0x05
	SegmentIDGlobal  int8 = 0x06
	SegmentIDExport  int8 = 0x07
	SegmentIDStart   int8 = 0x08
	SegmentIDElement int8 = 0x09
	SegmentIDCode    int8 = 0x0a
	SegmentIDData    int8 = 0x0b
)
