package main

import (
	"aoc2019/day13/intcode"
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

func makeGrid[T any](width, height int, fill T) [][]T {
	grid := make([][]T, height)
	for i := range height {
		row := make([]T, width)
		for j := range row {
			row[j] = fill
		}
		grid[i] = row
	}
	return grid
}

type Tile int

const (
	TileEmpty  Tile = 0
	TileWall   Tile = 1
	TileBlock  Tile = 2
	TilePaddle Tile = 3
	TileBall   Tile = 4
)

func (t Tile) String() string {
	switch t {
	case TileEmpty:
		return " "

	case TileWall:
		return "#"

	case TileBlock:
		return "*"

	case TilePaddle:
		return "_"

	case TileBall:
		return "@"
	}
	panic("unknown tile")
}

func solve() int {
	grid := makeGrid(44, 21, " ")
	pgrm := parse()
	pgrm[0] = 2
	ic := intcode.New(pgrm)

	ballX := 0
	paddleX := 0

	score := 0
	for {
		ic.Run()
		for step := range slices.Chunk(ic.Output(), 3) {
			x, y, tile := step[0], step[1], Tile(step[2])
			if x == -1 {
				score = int(tile)
				continue
			}
			if tile == TileBall {
				ballX = x
			}
			if tile == TilePaddle {
				paddleX = x
			}
			grid[y][x] = tile.String()
		}

		ic.ClearOutput()

		if paddleX > ballX {
			ic.AddInput(-1)
		} else if paddleX < ballX {
			ic.AddInput(1)
		} else {
			ic.AddInput(0)
		}

		if ic.Done() {
			break
		}
	}

	return score
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
