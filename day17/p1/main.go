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

func makeGridMap(s []int) map[Point]byte {
	grid := make(map[Point]byte)
	p := Point{}
	for i := range s {
		b := byte(s[i])
		if b == '\n' {
			p.x = 0
			p.y++
			continue
		}
		grid[p] = b
		p.x++
	}

	return grid
}

type Point struct {
	x, y int
}

func (p Point) inc(dir Point) Point {
	return Point{p.x + dir.x, p.y + dir.y}
}

func solve() int {
	pgrm := parse()
	ic := intcode.New(pgrm)
	ic.Run()
	assert.True(ic.Done())

	grid := makeGridMap(ic.Output())

	var dirs = []Point{
		{0, -1},
		{0, +1},
		{-1, 0},
		{+1, 0},
	}

	sum := 0
outer:
	for p, v := range grid {
		if v != '#' {
			continue
		}
		for _, dir := range dirs {
			if grid[p.inc(dir)] != '#' {
				continue outer
			}
		}
		sum += p.x * p.y
	}

	return sum
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
