package main

import (
	"aoc2019/day15/intcode"
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
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

func toInt(ss ...string) []int {
	s := strings.Join(ss, "\n") + "\n"
	slc := make([]int, len(s))
	for i := range s {
		slc[i] = int(s[i])
	}
	return slc
}

func solve() int {
	pgrm := parse()

	// jump if D is safe, and there's no ground on A or C
	// J = D && (!A || !C)

	ic := intcode.New(pgrm)
	ic.AddInput(toInt(
		"NOT A T",
		"NOT C J",
		"OR T J",
		"AND D J",
		"WALK",
	)...)
	ic.Run()
	assert.True(ic.Done())
	for _, v := range ic.Output() {
		if 0 <= v && v <= 255 {
			fmt.Print(string(v))
		} else {
			return v
		}
	}

	return 0
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
