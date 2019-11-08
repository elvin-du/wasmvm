package main

import (
	"log"
	"os"
	"strconv"
	"wasmvm/module"
)

var (
	wasmFileName = "./demo.wasm"
	globalModule *module.Module
)

var (
	params string
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	globalModule = module.NewModule()
	globalModule.Init(wasmFileName)
	globalModule.ParseSegments()

	funcName := os.Args[1]
	param1 := os.Args[2]
	param2 := os.Args[3]
	log.Printf("funcName:%s,param1:%s,param2:%s", funcName, param1, param2)

	seg := globalModule.Segments[module.SegmentIDExport]
	export := seg.(*module.SegExport)
	entryID := export.EntryID(funcName)

	p1, err := strconv.Atoi(param1)
	if nil != err {
		panic(err)
	}
	p2, err := strconv.Atoi(param2)
	if nil != err {
		panic(err)
	}

	Exec(entryID, uint32(p1), uint32(p2))
}
