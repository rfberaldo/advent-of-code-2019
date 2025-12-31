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

const (
	StatusWall  = 0
	StatusOK    = 1
	StatusFound = 2
)

const (
	MoveUp    = 1
	MoveDown  = 2
	MoveLeft  = 3
	MoveRight = 4
)

var movePoint = map[int]Point{
	MoveUp:    {0, -1},
	MoveDown:  {0, +1},
	MoveLeft:  {-1, 0},
	MoveRight: {+1, 0},
}

type Point struct {
	x, y int
}

func (p Point) move(move int) Point {
	return Point{p.x + movePoint[move].x, p.y + movePoint[move].y}
}

func solve() int {
	visited := make(map[Point]bool)
	oxygen := Point{}
	grid := make(map[Point]bool)

	// walk1 maps the whole area (grid) and find the oxygen point
	var walk1 func(*intcode.IntCode, Point, int)
	walk1 = func(ic *intcode.IntCode, curr Point, wantMove int) {
		curr = curr.move(wantMove)
		if visited[curr] {
			return
		}

		ic = ic.Clone()
		ic.AddInput(wantMove)
		ic.Run()

		status := ic.LastOutput()
		ic.ClearOutput()

		switch status {
		case StatusWall:
			visited[curr] = true

		case StatusFound:
			oxygen = curr

		case StatusOK:
			visited[curr] = true
			grid[curr] = true
			for i := 1; i <= 4; i++ {
				walk1(ic, curr, i)
			}
		}
	}

	pgrm := parse()
	for i := 1; i <= 4; i++ {
		walk1(intcode.New(pgrm), Point{}, i)
	}

	visited = make(map[Point]bool)

	// walk2 walks grid and returns the longer path
	var walk2 func(Point) int
	walk2 = func(curr Point) int {
		if visited[curr] {
			return 0
		}

		visited[curr] = true

		steps := 0
		for i := 1; i <= 4; i++ {
			next := curr.move(i)
			if grid[next] {
				steps = max(steps, walk2(next))
			}
		}

		return 1 + steps
	}

	return walk2(oxygen) - 1
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
