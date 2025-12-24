package intcode

import (
	"aoc2019/lib/assert"
	"fmt"
	"log"
	"strconv"
)

const DEBUG = false

type OpCode int

const (
	OpCodeAdd         OpCode = 1
	OpCodeMul         OpCode = 2
	OpCodeInput       OpCode = 3
	OpCodeOutput      OpCode = 4
	OpCodeGotoIfTrue  OpCode = 5
	OpCodeGotoIfFalse OpCode = 6
	OpCodeLessThan    OpCode = 7
	OpCodeEqual       OpCode = 8
	OpCodeExit        OpCode = 99
)

type PMode int // Parameter mode

const (
	PModePosition  PMode = 0 // default
	PModeImmediate PMode = 1
)

var stepsByOpCode = map[OpCode]int{
	OpCodeAdd:         4,
	OpCodeMul:         4,
	OpCodeInput:       2,
	OpCodeOutput:      2,
	OpCodeGotoIfTrue:  3,
	OpCodeGotoIfFalse: 3,
	OpCodeLessThan:    4,
	OpCodeEqual:       4,
}

type IntCode struct {
	pgrm       []int
	position   int
	opcode     OpCode
	paramsMode map[int]PMode
	inputQueue []int
	output     []int
}

func New(pgrm []int) *IntCode {
	assert.True(len(pgrm) > 0, "expect len > 0, got: ", len(pgrm))
	return &IntCode{pgrm: pgrm, paramsMode: make(map[int]PMode)}
}

func (ic *IntCode) AddInput(n int) {
	ic.inputQueue = append(ic.inputQueue, n)
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
	ic.position += stepsByOpCode[ic.opcode]
	ic.log("position=%v (%v)\n", ic.position, stepsByOpCode[ic.opcode])
	ic.checkPosition(ic.position)
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

func (ic *IntCode) Run() {
forLoop:
	for {
		ic.paramsMode = make(map[int]PMode)
		ic.opcode = OpCode(ic.pgrm[ic.position])
		ic.log("opcode=%v\n", ic.opcode)

		if ic.opcode >= 100 {
			ic.setParamsMode()
			ic.log("parameter mode: opcode=%v paramsMode=%v\n", ic.opcode, ic.paramsMode)
		}

		switch ic.opcode {
		case OpCodeExit:
			ic.log("exiting\n")
			return

		case OpCodeAdd:
			ic.writeToParam(3, ic.valueOfParam(1)+ic.valueOfParam(2))

		case OpCodeMul:
			ic.writeToParam(3, ic.valueOfParam(1)*ic.valueOfParam(2))

		case OpCodeInput:
			assert.True(len(ic.inputQueue) > 0, "input queue should be > 0, got: ", len(ic.inputQueue))

			ic.log("using input=%v then discarding\n", ic.inputQueue[0])
			ic.writeToParam(1, ic.inputQueue[0])
			ic.inputQueue = ic.inputQueue[1:]

		case OpCodeOutput:
			v := ic.valueOfParam(1)
			ic.output = append(ic.output, v)
			fmt.Printf("[OUTPUT] %v\n", v)

		case OpCodeGotoIfTrue:
			if ic.valueOfParam(1) != 0 {
				ic.position = ic.valueOfParam(2)
				continue forLoop
			}

		case OpCodeGotoIfFalse:
			if ic.valueOfParam(1) == 0 {
				ic.position = ic.valueOfParam(2)
				continue forLoop
			}

		case OpCodeLessThan:
			v := 0
			if ic.valueOfParam(1) < ic.valueOfParam(2) {
				v = 1
			}
			ic.writeToParam(3, v)

		case OpCodeEqual:
			v := 0
			if ic.valueOfParam(1) == ic.valueOfParam(2) {
				v = 1
			}
			ic.writeToParam(3, v)

		default:
			log.Fatalf("unknown opcode=%v", ic.opcode)
		}

		ic.step()
	}
}
