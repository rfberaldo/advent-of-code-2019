package main

import (
	"aoc2019/lib/util"
	"fmt"
	"time"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Walker struct {
	Point
	steps    int
	portmask int
	level    int
}

func (w Walker) move(dir Direction) Walker {
	w.x += dir.dx
	w.y += dir.dy
	w.steps++
	return w
}

type Portal struct {
	to    Point
	mask  int
	outer bool
}

func (w Walker) teleport(port Portal) (Walker, bool) {
	if w.level == 0 && port.outer {
		return Walker{}, false
	}

	w.Point = port.to
	w.portmask |= port.mask
	w.steps-- // stepping out of port should not count
	if port.outer {
		w.level--
	} else {
		w.level++
	}
	return w, true
}

type State struct {
	Point
	portmask int
	level    int
}

func (w Walker) state() State {
	return State{w.Point, w.portmask, w.level}
}

var dirs = []Direction{
	{0, -1},
	{+1, 0},
	{0, +1},
	{-1, 0},
}

func isLetter(r byte) bool {
	return 'A' <= r && r <= 'Z'
}

func parse() (grid map[Point]byte, portalByPoint map[Point]Portal, pAA Point, pZZ Point) {
	grid = make(map[Point]byte)
	pointByMask := make(map[int][]Point)
	xMax, yMax := 0, 0

	for y, line := range util.Input() {
		yMax = max(yMax, y)
		for x := range line {
			xMax = max(xMax, x)
			if line[x] != ' ' && line[x] != '#' {
				grid[Point{x, y}] = line[x]
			}
		}
	}

	for p, b := range grid {
		if !isLetter(b) {
			continue
		}
		mask := 1 << (b - 'A')
		dot := false
		for _, dir := range dirs {
			edge := Point{p.x + dir.dx, p.y + dir.dy}
			if b2, ok := grid[edge]; ok && isLetter(byte(b2)) {
				mask |= 1 << (b2 - 'A')
			} else if b2 == '.' {
				dot = true
			}
		}
		if dot {
			pointByMask[mask] = append(pointByMask[mask], p)
		}
	}

	isOuter := func(p Point) bool {
		return p.x-2 < 0 || p.y-2 < 0 || p.x+2 >= xMax || p.y+2 >= yMax
	}

	portalByPoint = make(map[Point]Portal)
	for mask, points := range pointByMask {
		if len(points) == 2 {
			portalByPoint[points[0]] = Portal{points[1], mask, isOuter(points[0])}
			portalByPoint[points[1]] = Portal{points[0], mask, isOuter(points[1])}
		}
	}

	pZZ = pointByMask[1<<('Z'-'A')][0]
	pAA = pointByMask[1<<0][0]
	return
}

func solve() int {
	grid, portalByPoint, pAA, pZZ := parse()

	visited := make(map[State]bool)
	queue := []Walker{{Point: pZZ}}

	result := 1_000_000
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Point == pAA && curr.level == 0 {
			result = min(result, curr.steps-2)
		}

		if curr.steps >= result {
			continue
		}

		b, ok := grid[curr.Point]
		if !ok {
			continue
		}

		if visited[curr.state()] {
			continue
		}
		visited[curr.state()] = true

		for _, dir := range dirs {
			next := curr.move(dir)
			queue = append(queue, next)
		}

		if 'A' <= b && b <= 'Z' {
			port, ok := portalByPoint[curr.Point]
			if !ok {
				continue
			}
			next, ok := curr.teleport(port)
			if !ok {
				continue
			}
			queue = append(queue, next)
		}
	}

	return result
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
