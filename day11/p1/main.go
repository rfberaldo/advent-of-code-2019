package main

import (
	"aoc2019/day11/intcode"
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

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

// https://en.wikipedia.org/wiki/Rotation_matrix
func cw(dir Direction) Direction {
	return Direction{-dir.dy, dir.dx}
}

// https://en.wikipedia.org/wiki/Rotation_matrix
func ccw(dir Direction) Direction {
	return Direction{dir.dy, -dir.dx}
}

func solve() int {
	panel := make(map[Point]int)
	robot := Point{0, 0}
	dir := Direction{0, -1} // up

	ic := intcode.New(parse())

	for !ic.Done() {
		color := panel[robot]
		ic.AddInput(color)
		ic.Run()

		out := ic.Output()
		assert.True(len(out) == 2)
		ic.ClearOutput()

		color, turn := out[0], out[1]
		panel[robot] = color
		if turn == 0 {
			dir = ccw(dir)
		} else {
			dir = cw(dir)
		}

		// walk
		robot.x += dir.dx
		robot.y += dir.dy
	}

	return len(panel)
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
