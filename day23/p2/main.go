package main

import (
	"aoc2019/day23/intcode"
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

	xNat, yNat := 0, 0
	yLast := 0

	for {
		for i := range size {
			pcs[i].Run()

			if len(pcs[i].Output()) == 0 {
				continue
			}

			for out := range slices.Chunk(pcs[i].Output(), 3) {
				j, x, y := out[0], out[1], out[2]
				if j == 255 {
					xNat, yNat = x, y
					continue
				}
				pcs[j].AddInput(x, y)
				pcs[i].ClearOutput()
			}
		}

		idle := 0
		for i := range size {
			if pcs[i].WaitingInput() {
				idle++

				if idle >= size {
					if yNat == yLast {
						return yNat
					}
					yLast = yNat
					pcs[0].AddInput(xNat, yNat)
				}
			}
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
