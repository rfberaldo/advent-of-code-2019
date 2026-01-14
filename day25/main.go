package main

import (
	"aoc2019/day23/intcode"
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	input := util.Input()
	var pgrm []int
	for s := range strings.SplitSeq(input[0], ",") {
		n, err := strconv.Atoi(s)
		assert.NoErr(err)
		pgrm = append(pgrm, n)
	}
	return pgrm
}

func toInt(str string) []int {
	s := make([]int, len(str))
	for i := range str {
		s[i] = int(str[i])
	}
	return s
}

func toString(s []int) string {
	var sb strings.Builder
	for i := range s {
		sb.WriteByte(byte(s[i]))
	}
	return sb.String()
}

// hologram, space heater, astronaut ice cream, antenna

func main() {
	pgrm := parse()
	ic := intcode.New(pgrm)

	for {
		ic.Run()
		fmt.Println(toString(ic.Output()))
		ic.ClearOutput()
		if ic.WaitingInput() {
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			assert.NoErr(err)
			ic.AddInput(toInt(line)...)
			continue
		}

		return
	}
}
