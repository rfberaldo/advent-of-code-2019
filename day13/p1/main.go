package main

import (
	"aoc2019/day13/intcode"
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

type Tile int

const (
	TileEmpty  Tile = 0
	TileWall   Tile = 1
	TileBlock  Tile = 2
	TilePaddle Tile = 3
	TileBall   Tile = 4
)

func solve() int {
	ic := intcode.New(parse())
	ic.Run()
	assert.True(ic.Done())
	out := ic.Output()

	result := 0
	for i := 2; i < len(out); i += 3 {
		if Tile(out[i]) == TileBlock {
			result++
		}
	}

	return result
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
