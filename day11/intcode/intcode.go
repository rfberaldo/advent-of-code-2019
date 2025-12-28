package intcode

import (
	"aoc2019/lib/assert"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type OpCode int

const (
	OpCodeAdd           OpCode = 1
	OpCodeMul           OpCode = 2
	OpCodeInput         OpCode = 3
	OpCodeOutput        OpCode = 4
	OpCodeGotoIfTrue    OpCode = 5
	OpCodeGotoIfFalse   OpCode = 6
	OpCodeLessThan      OpCode = 7
	OpCodeEqual         OpCode = 8
	OpCodeUpdateRelBase OpCode = 9
	OpCodeExit          OpCode = 99
)

// implements [fmt.Stringer]
func (opcode OpCode) String() string {
	switch opcode {
	case OpCodeAdd:
		return "ADD"
	case OpCodeMul:
		return "MUL"
	case OpCodeInput:
		return "INPUT"
	case OpCodeOutput:
		return "OUTPUT"
	case OpCodeGotoIfTrue:
		return "GOTO_IF_TRUE"
	case OpCodeGotoIfFalse:
		return "GOTO_IF_FALSE"
	case OpCodeLessThan:
		return "LESS_THAN"
	case OpCodeEqual:
		return "EQUAL"
	case OpCodeUpdateRelBase:
		return "UPDATE_REL_BASE"
	case OpCodeExit:
		return "EXIT"
	}
	return "UNKNOWN"
}

type PMode int // Parameter mode

const (
	PModePosition  PMode = 0 // default
	PModeImmediate PMode = 1
	PModeRelative  PMode = 2
)

var paramCountByOpCode = map[OpCode]int{
	OpCodeAdd:           3,
	OpCodeMul:           3,
	OpCodeInput:         1,
	OpCodeOutput:        1,
	OpCodeGotoIfTrue:    2,
	OpCodeGotoIfFalse:   2,
	OpCodeLessThan:      3,
	OpCodeEqual:         3,
	OpCodeUpdateRelBase: 1,
}

type IntCode struct {
	name             string
	pgrm             []int
	position         int
	paramsMode       map[int]PMode
	inputQueue       []int
	feedback         *IntCode
	feedbackPosition int
	relBase          int
	output           []int
	done             bool
	debug            bool
}

func New(pgrm []int) *IntCode {
	assert.True(len(pgrm) > 0, "expect len > 0, got: ", len(pgrm))
	return &IntCode{
		pgrm:       slices.Clone(pgrm),
		paramsMode: make(map[int]PMode),
	}
}

func (ic *IntCode) SetDebug(debug bool) *IntCode {
	ic.debug = debug
	return ic
}

func (ic *IntCode) SetName(name string) *IntCode {
	ic.name = name
	return ic
}

func (ic *IntCode) AddInput(n ...int) *IntCode {
	ic.inputQueue = append(ic.inputQueue, n...)
	return ic
}

func (ic *IntCode) AddFeedback(feedback *IntCode) *IntCode {
	ic.feedback = feedback
	return ic
}

func (ic *IntCode) Done() bool {
	return ic.done
}

func (ic *IntCode) Output() []int {
	return ic.output
}

func (ic *IntCode) ClearOutput() {
	ic.output = ic.output[:0]
}

func (ic *IntCode) LastOutput() int {
	ic.must(len(ic.output) > 0, "output length must be > 0")
	return ic.output[len(ic.output)-1]
}

func (ic *IntCode) fatal(format string, v ...any) {
	if ic.name != "" {
		fmt.Printf("%v: ", ic.name)
	}
	fmt.Printf(format+"\n", v...)
	os.Exit(1)
}

func (ic *IntCode) must(b bool, format string, v ...any) {
	if b {
		return
	}

	ic.fatal(format, v...)
}

func (ic *IntCode) log(format string, v ...any) {
	if !ic.debug {
		return
	}

	if ic.name != "" {
		fmt.Printf("%v: ", ic.name)
	}
	fmt.Printf(format+"\n", v...)
}

func (ic *IntCode) setParamsMode(opcode OpCode) OpCode {
	ic.must(opcode > 99, "opcode must be > 99, got=%v", opcode)

	digits := strconv.Itoa(int(opcode))
	n := 1
	for i := len(digits) - 3; i >= 0; i-- {
		ic.paramsMode[n] = PMode(int(digits[i] - '0'))
		n++
	}

	return OpCode(opcode % 100)
}

func (ic *IntCode) step(inc int) {
	ic.position += inc
	ic.must(
		0 <= ic.position && ic.position < len(ic.pgrm),
		"position=%v out of bounds length=%v", ic.position, len(ic.pgrm),
	)
	ic.log("position=%v", ic.position)
}

func (ic *IntCode) checkAddr(addr int) {
	ic.must(addr >= 0, "addr must be >= 0, got=%v", addr)

	if delta := addr - len(ic.pgrm); delta >= 0 {
		ic.pgrm = slices.Grow(ic.pgrm, delta+1)
		ic.pgrm = ic.pgrm[:addr+1]
		ic.log("expand length=%v", len(ic.pgrm))
	}

	ic.must(addr < len(ic.pgrm), "addr=%v out of bounds length=%v", addr, len(ic.pgrm))
}

func (ic *IntCode) deref(addr int) int {
	ic.checkAddr(addr)
	ic.log("deref addr=%v value=%v", addr, ic.pgrm[addr])
	return ic.pgrm[addr]
}

func (ic *IntCode) valueOf(addr int) int {
	ic.checkAddr(addr)
	ic.log("valueOf addr=%v value=%v", addr, ic.pgrm[addr])
	return ic.pgrm[addr]
}

func (ic *IntCode) writeTo(addr, v int) {
	ic.checkAddr(addr)
	ic.log("writeTo addr=%v value=%v", addr, v)
	ic.pgrm[addr] = v
}

func (ic *IntCode) writeToParam(n, v int) {
	ic.must(n > 0, "parameter must be > 0")

	switch ic.paramsMode[n] {
	case PModeImmediate:
		ic.fatal("write parameter mode must not be immediate")

	case PModeRelative:
		ic.writeTo(ic.relBase+ic.valueOf(ic.position+n), v)

	// default is position mode
	default:
		ic.writeTo(ic.deref(ic.position+n), v)
	}
}

func (ic *IntCode) valueOfParam(n int) int {
	ic.must(n > 0, "parameter must be > 0")

	switch ic.paramsMode[n] {
	case PModeImmediate:
		ic.log("ParamMode=Immediate")
		return ic.valueOf(ic.position + n)

	case PModeRelative:
		ic.log("ParamMode=Relative base=%v", ic.relBase)
		return ic.valueOf(ic.relBase + ic.valueOf(ic.position+n))

	// default is position mode
	default:
		ic.log("ParamMode=Position")
		return ic.valueOf(ic.deref(ic.position + n))
	}
}

func (ic *IntCode) handleInput() bool {
	if len(ic.inputQueue) > 0 {
		ic.log("input value=%v", ic.inputQueue[0])
		ic.writeToParam(1, ic.inputQueue[0])
		ic.inputQueue = ic.inputQueue[1:]
		return true
	}

	if ic.feedback == nil || ic.feedbackPosition >= len(ic.feedback.output) {
		ic.log("no more input/feedback, call Run again when ready")
		return false
	}

	ic.log("input position=%v from feedback=%v", ic.feedbackPosition, ic.feedback.output)
	ic.writeToParam(1, ic.feedback.output[ic.feedbackPosition])
	ic.feedbackPosition++
	return true
}

func (ic *IntCode) Run() {
forLoop:
	for {
		ic.log("------------")
		ic.paramsMode = make(map[int]PMode)
		opcode := OpCode(ic.pgrm[ic.position])
		ic.log("opcode=%d [%v]", opcode, opcode)

		if opcode >= 100 {
			opcode = ic.setParamsMode(opcode)
			ic.log("parameter mode opcode=%d [%v] paramsMode=%v", opcode, opcode, ic.paramsMode)
		}

		switch opcode {
		case OpCodeExit:
			ic.done = true
			return

		case OpCodeAdd:
			ic.writeToParam(3, ic.valueOfParam(1)+ic.valueOfParam(2))

		case OpCodeMul:
			ic.writeToParam(3, ic.valueOfParam(1)*ic.valueOfParam(2))

		case OpCodeInput:
			if !ic.handleInput() {
				return
			}

		case OpCodeOutput:
			v := ic.valueOfParam(1)
			ic.output = append(ic.output, v)

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

		case OpCodeUpdateRelBase:
			ic.relBase += ic.valueOfParam(1)
			ic.log("relative base=%v", ic.relBase)

		default:
			ic.fatal("unknown opcode=%d", opcode)
		}

		if v, ok := paramCountByOpCode[opcode]; !ok {
			ic.fatal("unknown param count for opcode=%v", opcode)
		} else {
			ic.step(v + 1)
		}
	}
}
