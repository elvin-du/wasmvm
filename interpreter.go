package main

import "wasmvm/module"

func Exec(index, param1, param2 uint32) {
	seg := globalModule.Segments[module.SegmentIDCode]
	code := seg.(*module.SegCode)
	_ = code.FunctionBody[index]

}

func exec()  {
	
}
