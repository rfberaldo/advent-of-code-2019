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
}

func (w Walker) move(dir Direction) Walker {
	w.x += dir.dx
	w.y += dir.dy
	w.steps++
	return w
}

type Portal struct {
	to   Point
	mask int
}

func (w Walker) teleport(p Portal) Walker {
	w.Point = p.to
	w.portmask |= p.mask
	w.steps-- // stepping out of port should not count
	return w
}

type State struct {
	Point
	portmask int
}

func (w Walker) state() State {
	return State{w.Point, w.portmask}
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

func parse() (grid map[Point]byte, portals map[int][]Point) {
	grid = make(map[Point]byte)
	portals = make(map[int][]Point)

	for y, line := range util.Input() {
		for x := range line {
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
			// it's a letter between a dot and another letter
			portals[mask] = append(portals[mask], p)
		}
	}
	return grid, portals
}

func solve() int {
	grid, portals := parse()

	start := portals[1<<('Z'-'A')][0]
	end := portals[1<<0][0]

	portalByPoint := make(map[Point]Portal)
	for mask, points := range portals {
		if len(points) != 2 {
			continue
		}
		portalByPoint[points[0]] = Portal{points[1], mask}
		portalByPoint[points[1]] = Portal{points[0], mask}
	}

	visited := make(map[State]bool)
	queue := []Walker{{Point: start}}

	result := 1_000_000
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Point == end {
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

		if 'A' <= b && b <= 'Z' {
			if p, ok := portalByPoint[curr.Point]; ok {
				next := curr.teleport(p)
				queue = append(queue, next)
			}
		}

		for _, dir := range dirs {
			next := curr.move(dir)
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
