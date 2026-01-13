package main

import (
	"aoc2019/day15/intcode"
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"slices"
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
	const size = 50
	pgrm := parse()
	pcs := make([]*intcode.IntCode, size)
	for i := range size {
		pcs[i] = intcode.New(pgrm).AddInput(i, -1)
	}

	for i := 0; ; i++ {
		if i == size {
			i = 0
		}
		pcs[i].Run()
		if len(pcs[i].Output()) == 0 {
			continue
		}

		for out := range slices.Chunk(pcs[i].Output(), 3) {
			j, x, y := out[0], out[1], out[2]
			if j == 255 {
				return y
			}
			pcs[j].AddInput(x, y)
			pcs[i].ClearOutput()
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
