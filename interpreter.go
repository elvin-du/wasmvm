package main

import (
	"bytes"
	"io"
	"log"
	"wasmvm/leb128"
	"wasmvm/module"
)

type VM struct {
	r     io.Reader
	stack []int64
	local []int64
}

func NewVM(r io.Reader) *VM {
	vm := &VM{}
	vm.r = r
	vm.stack = []int64{}
	vm.local = []int64{}
	return vm
}

func Exec(index, param1, param2 uint32) {
	seg := globalModule.Segments[module.SegmentIDCode]
	code := seg.(*module.SegCode)
	r := bytes.NewReader(code.FunctionBody[index].Code)
	vm := NewVM(r)
	vm.local = append(vm.local, int64(param1))
	vm.local = append(vm.local, int64(param2))
	vm.exec(r)
}

func (vm *VM) push(i int64) {
	vm.stack = append(vm.stack, i)
}

func (vm *VM) pop() int64 {
	i := vm.stack[len(vm.stack)-1]

	vm.stack = vm.stack[:len(vm.stack)-1]
	return i
}

func (vm *VM) exec(r io.Reader) {
	localCount, err := leb128.ReadVarint32(r)
	if nil != err {
		panic(err)
	}
	log.Printf("[exec] localCount:%d", localCount)

	for {
		op, err := leb128.ReadVarUint32(r)
		if nil != err {
			panic(err)
		}
		log.Printf("[exec]opcode:0x%x", op)
		switch op {
		case GetLocal:
			vm.GetLocal()
		case I32Add:
			vm.I32Add()
		case End:
			vm.End()
			return
		default:
			log.Printf("unkown opcode")
		}
	}
}
