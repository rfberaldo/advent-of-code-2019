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
	var walk func(*intcode.IntCode, Point, int) (int, bool)
	visited := make(map[Point]bool)

	walk = func(ic *intcode.IntCode, curr Point, wantMove int) (int, bool) {
		curr = curr.move(wantMove)
		if visited[curr] {
			return 0, false
		}

		ic = ic.Clone()
		ic.AddInput(wantMove)
		ic.Run()

		status := ic.LastOutput()
		ic.ClearOutput()

		switch status {
		case StatusWall:
			visited[curr] = true
			return 0, false

		case StatusFound:
			return 1, true

		case StatusOK:
			visited[curr] = true
			for i := 1; i <= 4; i++ {
				if m, ok := walk(ic, curr, i); ok {
					return 1 + m, ok
				}
			}
			return 0, false
		}

		panic("unreacheable")
	}

	pgrm := parse()
	for i := 1; i <= 4; i++ {
		if m, ok := walk(intcode.New(pgrm), Point{}, i); ok {
			return m
		}
	}
	return 0
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
