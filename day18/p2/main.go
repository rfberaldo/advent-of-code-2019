package main

import (
	"aoc2019/lib/util"
	"fmt"
	"runtime"
	"time"
)

// Remember to update the area in the middle of input:
//  ...       @#@
//  .@.  -->  ###
//  ...       @#@

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Walker struct {
	points  [4]Point
	active  int
	steps   int
	keyring string
	keymask int
}

func (w Walker) point() Point {
	return w.points[w.active]
}

func (w Walker) move(dir Direction) Walker {
	w.points[w.active].x += dir.dx
	w.points[w.active].y += dir.dy
	w.steps++
	return w
}

func (w Walker) teleport(id int) Walker {
	w.active = id
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

type State struct {
	points  [4]Point
	active  int
	keymask int
}

func (w Walker) state() State {
	return State{w.points, w.active, w.keymask}
}

func parse() (map[Point]byte, []Point, map[byte]Point) {
	grid := make(map[Point]byte)
	keys := make(map[byte]Point)
	var entry []Point
	for y, line := range util.Input() {
		for x, r := range line {
			if r == '@' {
				entry = append(entry, Point{x, y})
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

// runs in ~40s and uses ~2GB memory
func solve_dfs() int {
	grid, entries, keys := parse()

	visited := make(map[State]int)

	var walk func(Walker) int
	walk = func(curr Walker) int {
		b, ok := grid[curr.point()]
		if !ok || b == '#' {
			return -1
		}

		if 'A' <= b && b <= 'Z' && !curr.hasKey(b) {
			return -1
		}

		state := curr.state()
		if v, ok := visited[state]; ok {
			return v
		}
		visited[state] = -1

		var spawns []Walker
		if 'a' <= b && b <= 'z' && !curr.hasKey(b) {
			curr.addKey(b)

			// found a new key, spawn on other checkpoints
			for id := range 4 {
				if id == curr.active {
					continue
				}
				next := curr.teleport(id)
				spawns = append(spawns, next)
			}
		}

		if len(curr.keyring) == len(keys) {
			return 0
		}

		result := 1_000_000

		for _, next := range spawns {
			if steps := walk(next); steps >= 0 {
				result = min(result, steps)
			}
		}

		for _, dir := range dirs {
			if steps := walk(curr.move(dir)); steps >= 0 {
				result = min(result, 1+steps)
			}
		}

		if result >= 1_000_000 {
			delete(visited, state)
		} else {
			visited[state] = result
		}

		return result
	}

	return walk(Walker{points: [4]Point{entries[0], entries[1], entries[2], entries[3]}})
}

// runs in ~20s and uses ~6GB memory
func solve_bfs() int {
	grid, entries, keys := parse()

	visited := make(map[State]bool)

	queue := []Walker{
		{points: [4]Point{entries[0], entries[1], entries[2], entries[3]}},
	}

	result := 1_000_000
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.steps >= result {
			continue
		}

		b, ok := grid[curr.point()]
		if !ok || b == '#' {
			continue
		}

		if 'A' <= b && b <= 'Z' && !curr.hasKey(b) {
			continue
		}

		state := curr.state()
		if visited[state] {
			continue
		}
		visited[state] = true

		if 'a' <= b && b <= 'z' && !curr.hasKey(b) {
			curr.addKey(b)

			if len(curr.keyring) == len(keys) {
				result = min(result, curr.steps)
				continue
			}

			// found a new key, spawn on other checkpoints
			for id := range 4 {
				if id == curr.active {
					continue
				}
				next := curr.teleport(id)
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
	fmt.Println("Result (BFS):", solve_bfs())
	fmt.Println(time.Since(start))

	runtime.GC()

	start = time.Now()
	fmt.Println("Result (DFS):", solve_dfs())
	fmt.Println(time.Since(start))
}
