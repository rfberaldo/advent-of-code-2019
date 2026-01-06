package main

import (
	"aoc2019/lib/util"
	"fmt"
	"runtime"
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
	id      int
	keyring string
	keymask int
	steps   int
}

func (w Walker) move(dir Direction) Walker {
	w.x += dir.dx
	w.y += dir.dy
	w.steps++
	return w
}

func (w *Walker) addKey(b byte) {
	w.keyring += string(b)
	w.keymask |= 1 << (b - 'a')
}

func (w Walker) hasKey(b byte) bool {
	if 'A' <= b && b <= 'Z' {
		b += 'a' - 'A' // to lower
	}
	return (w.keymask & (1 << (b - 'a'))) != 0
}

func (w Walker) key() string {
	return fmt.Sprintf("%d-%d-%d", w.x, w.y, w.keymask)
}

func parse() (map[Point]byte, Point, map[byte]Point) {
	grid := make(map[Point]byte)
	keys := make(map[byte]Point)
	var entry Point
	for y, line := range util.Input() {
		for x, r := range line {
			if r == '@' {
				entry = Point{x, y}
			}
			if 'a' <= r && r <= 'z' {
				keys[byte(r)] = Point{x, y}
			}
			grid[Point{x, y}] = byte(r)
		}
	}
	return grid, entry, keys
}

var dirs = []Direction{
	{0, -1},
	{+1, 0},
	{0, +1},
	{-1, 0},
}

func solve_dfs() int {
	grid, start, keys := parse()

	visited := make(map[string]int)

	var walk func(Walker) int
	walk = func(curr Walker) int {
		b, ok := grid[curr.Point]
		if !ok || b == '#' {
			return -1
		}

		if v, ok := visited[curr.key()]; ok {
			return v
		}

		if 'A' <= b && b <= 'Z' && !curr.hasKey(b) {
			return -1
		}

		if 'a' <= b && b <= 'z' && !curr.hasKey(b) {
			curr.addKey(b)
		}

		if len(curr.keyring) == len(keys) {
			return 0
		}

		visited[curr.key()] = -1

		result := 1_000_000
		for _, dir := range dirs {
			if steps := walk(curr.move(dir)); steps >= 0 {
				result = min(result, 1+steps)
			}
		}

		if result >= 1_000_000 {
			delete(visited, curr.key())
		} else {
			visited[curr.key()] = result
		}

		return result
	}

	return walk(Walker{Point: start})
}

func solve_bfs() int {
	grid, start, keys := parse()

	visited := make(map[string]bool)
	queue := []Walker{{Point: start}}

	result := 1_000_000
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		b, ok := grid[curr.Point]
		if !ok || b == '#' {
			continue
		}

		if visited[curr.key()] {
			continue
		}

		if 'A' <= b && b <= 'Z' && !curr.hasKey(b) {
			continue
		}

		if 'a' <= b && b <= 'z' && !curr.hasKey(b) {
			curr.addKey(b)
		}

		if len(curr.keyring) == len(keys) {
			result = min(result, curr.steps)
			continue
		}

		visited[curr.key()] = true

		for _, dir := range dirs {
			next := curr.move(dir)
			queue = append(queue, next)
		}
	}

	return result
}

func main() {
	start := time.Now()
	fmt.Println("Result (BFS):", solve_bfs())
	fmt.Println(time.Since(start))

	runtime.GC()

	start = time.Now()
	fmt.Println("Result (DFS):", solve_dfs())
	fmt.Println(time.Since(start))
}
