package main

import (
	"log"
	"wasmvm/leb128"
)

const (
	GetLocal = 0x20
	I32Add   = 0x6A
	End      = 0x0B
)

func (vm *VM) GetLocal() {
	index, err := leb128.ReadVarUint32(vm.r)
	if nil != err {
		panic(err)
	}
	param := vm.local[index]
	vm.push(param)
}

func (vm *VM) I32Add() {
	param1 := vm.pop()
	param2 := vm.pop()
	vm.push(param1 + param2)
}

func (vm *VM) End() {
	log.Printf("value of top stack:%d", vm.pop())
}
