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

	read := func(x, y int) int {
		ic := intcode.New(pgrm).AddInput(x, y)
		ic.Run()
		assert.True(ic.Done())
		return ic.LastOutput()
	}

	x, y := 0, 0
	for {
		if read(x+99, y) == 1 {
			if read(x, y+99) == 1 {
				break
			}
			x++
			continue
		}
		y++
	}

	return x*10_000 + y
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
