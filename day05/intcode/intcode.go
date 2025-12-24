package intcode

import (
	"aoc2019/lib/assert"
	"fmt"
	"log"
	"strconv"
	"testing"
)

const DEBUG = false

type OpCode int

const (
	OpCodeAdd    OpCode = 1
	OpCodeMul    OpCode = 2
	OpCodeInput  OpCode = 3
	OpCodeOutput OpCode = 4
	OpCodeExit   OpCode = 99
)

type PMode int // Parameter mode

const (
	PModePosition  PMode = 0 // default
	PModeImmediate PMode = 1
)

var stepsByOp = map[OpCode]int{
	OpCodeAdd:    4,
	OpCodeMul:    4,
	OpCodeInput:  2,
	OpCodeOutput: 2,
}

type IntCode struct {
	pgrm       []int
	position   int
	opcode     OpCode
	paramsMode map[int]PMode
}

func New(pgrm []int) *IntCode {
	assert.True(len(pgrm) > 0, "expect len > 0, got: ", len(pgrm))
	return &IntCode{pgrm: pgrm, paramsMode: make(map[int]PMode)}
}

func (ic *IntCode) log(format string, v ...any) {
	if !DEBUG {
		return
	}
	fmt.Printf(format, v...)
}

func (ic *IntCode) setParamsMode() {
	assert.True(ic.opcode > 99, "opcode must be > 99, got: ", ic.opcode)

	digits := strconv.Itoa(int(ic.opcode))
	n := 1
	for i := len(digits) - 3; i >= 0; i-- {
		ic.paramsMode[n] = PMode(int(digits[i] - '0'))
		n++
	}

	ic.opcode = ic.opcode % 100
}

func (ic *IntCode) step() {
	ic.position += stepsByOp[ic.opcode]
	ic.log("position=%v (%v)\n", ic.position, stepsByOp[ic.opcode])
	ic.checkPosition(ic.position)
	ic.paramsMode = make(map[int]PMode)
}

func (ic *IntCode) checkPosition(pos int) {
	assert.True(0 <= pos && pos < len(ic.pgrm), "position=", pos, " out of bounds")
}

func (ic *IntCode) valueOf(pos int) int {
	ic.checkPosition(pos)
	ic.log("valueOf: pos=%v value=%v\n", pos, ic.pgrm[pos])
	return ic.pgrm[pos]
}

func (ic *IntCode) valueOfAddr(pos int) int {
	ic.checkPosition(pos)
	addr := ic.pgrm[pos]
	ic.checkPosition(addr)
	ic.log("valueOfAddr: pos=%v addr=%v value=%v\n", pos, addr, ic.pgrm[addr])
	return ic.pgrm[addr]
}

func (ic *IntCode) writeToParam(n, v int) {
	assert.True(n > 0, "parameter must be > 0, got: ", n)

	// write is always position mode
	ic.checkPosition(ic.position + n)
	addr := ic.pgrm[ic.position+n]
	ic.checkPosition(addr)
	ic.log("writing: addr=%v value=%v\n", addr, v)
	ic.pgrm[addr] = v
}

func (ic *IntCode) valueOfParam(n int) int {
	assert.True(n > 0, "parameter must be > 0, got: ", n)

	if ic.paramsMode[n] == PModeImmediate {
		return ic.valueOf(ic.position + n)
	}

	// default is position mode
	return ic.valueOfAddr(ic.position + n)
}

func (ic *IntCode) Run() []int {
	for {
		ic.opcode = OpCode(ic.pgrm[ic.position])
		ic.log("opcode=%v\n", ic.opcode)

		if ic.opcode >= 100 {
			ic.setParamsMode()
			ic.log("parameter mode: opcode=%v paramsMode=%v\n", ic.opcode, ic.paramsMode)
		}

		switch ic.opcode {
		case OpCodeExit:
			ic.log("exiting\n")
			return ic.pgrm

		case OpCodeAdd:
			ic.writeToParam(3, ic.valueOfParam(1)+ic.valueOfParam(2))

		case OpCodeMul:
			ic.writeToParam(3, ic.valueOfParam(1)*ic.valueOfParam(2))

		case OpCodeInput:
			fmt.Print("Enter an integer: ")
			var input int
			if testing.Testing() {
				input = 69
			} else {
				_, err := fmt.Scanln(&input)
				assert.NoErr(err)
			}
			ic.writeToParam(1, input)

		case OpCodeOutput:
			fmt.Printf("[OUTPUT] %v\n", ic.valueOfParam(1))

		default:
			log.Fatalf("unknown opcode=%v", ic.opcode)
		}

		ic.step()
	}
}
