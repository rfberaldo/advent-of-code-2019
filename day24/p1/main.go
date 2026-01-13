package main

import (
	"aoc2019/lib/util"
	"fmt"
	"maps"
	"time"
)

type Point struct {
	x, y int
}

func parse() map[Point]byte {
	grid := make(map[Point]byte)
	for y, line := range util.Input() {
		for x := range line {
			grid[Point{x, y}] = line[x]
		}
	}
	return grid
}

func points(grid map[Point]byte) int {
	sum := 0
	for p, b := range grid {
		if b == '#' {
			sum |= 1 << (5*p.y + p.x)
		}
	}
	return sum
}

func timelapse(grid map[Point]byte) {
	adjs := func(p Point) int {
		count := 0
		for _, dir := range []Point{{0, -1}, {0, +1}, {-1, 0}, {+1, 0}} {
			b, ok := grid[Point{p.x + dir.x, p.y + dir.y}]
			if ok && b == '#' {
				count++
			}
		}
		return count
	}

	updt := make(map[Point]byte)
	for p, b := range grid {
		count := adjs(p)
		if b == '#' && count != 1 {
			updt[p] = '.'
		}
		if b == '.' && (1 <= count && count <= 2) {
			updt[p] = '#'
		}
	}
	maps.Copy(grid, updt)
}

func solve() int {
	grid := parse()

	history := make(map[int]bool)
	for {
		timelapse(grid)
		pts := points(grid)
		if history[pts] {
			return pts
		}
		history[pts] = true
	}
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
