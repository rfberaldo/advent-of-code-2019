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

func solve() int {
	pgrm := parse()

	sum := 0
	for y := range 50 {
		for x := range 50 {
			ic := intcode.New(pgrm)
			ic.AddInput(x, y)
			ic.Run()
			assert.True(ic.Done())
			sum += ic.LastOutput()
		}
	}

	return sum
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
