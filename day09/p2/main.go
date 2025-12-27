package main

import (
	"aoc2019/day09/intcode"
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

func main() {
	start := time.Now()

	ic := intcode.New(parse()).AddInput(2)
	ic.Run()
	assert.True(ic.Done())

	fmt.Println("Result:", ic.LastOutput())
	fmt.Println(time.Since(start))
}
