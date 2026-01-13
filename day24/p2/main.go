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

func newEmptyGrid() map[Point]byte {
	grid := make(map[Point]byte)
	for y := range 5 {
		for x := range 5 {
			grid[Point{x, y}] = '.'
		}
	}
	return grid
}

func bugCount(grids []map[Point]byte) int {
	sum := 0
	for i := range grids {
		for _, b := range grids[i] {
			if b == '#' {
				sum++
			}
		}
	}
	return sum
}

func timelapse(grids []map[Point]byte) {
	adjBugCount := func(i int, p Point) int {
		if p.x == 2 && p.y == 2 {
			return 0
		}

		dirs := [4]Point{
			{0, -1}, // top
			{0, +1}, // bottom
			{-1, 0}, // left
			{+1, 0}, // right
		}

		// index based on dirs index
		outerEdge := [4]Point{
			{2, 1}, // top
			{2, 3}, // bottom
			{1, 2}, // left
			{3, 2}, // right
		}

		// index based on dirs index
		innerEdges := [4][5]Point{
			{{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}}, // bottom row
			{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // top row
			{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}}, // right column
			{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}, // left column
		}

		count := 0
		for j, dir := range dirs {
			edge := Point{p.x + dir.x, p.y + dir.y}
			isInner := edge.x == 2 && edge.y == 2
			if b, exists := grids[i][edge]; exists {
				if b == '#' {
					count++
					continue
				}

				if !isInner || i == len(grids)-1 {
					continue
				}

				for _, edge := range innerEdges[j] {
					if grids[i+1][edge] == '#' {
						count++
					}
				}
				continue
			}

			if i == 0 {
				continue
			}

			if grids[i-1][outerEdge[j]] == '#' {
				count++
			}
		}

		return count
	}

	state := make([]map[Point]byte, len(grids))
	for i := range grids {
		state[i] = make(map[Point]byte)
		for p, b := range grids[i] {
			count := adjBugCount(i, p)
			if b == '#' && count != 1 {
				state[i][p] = '.'
			}
			if b == '.' && (1 <= count && count <= 2) {
				state[i][p] = '#'
			}
		}
	}

	for i := range grids {
		maps.Copy(grids[i], state[i])
	}
}

func solve() int {
	const levels = 200

	grids := make([]map[Point]byte, levels+1)
	for i := range len(grids) {
		if i == levels/2 {
			grids[i] = parse()
		} else {
			grids[i] = newEmptyGrid()
		}
	}

	for range levels {
		timelapse(grids)
	}

	return bugCount(grids)
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
