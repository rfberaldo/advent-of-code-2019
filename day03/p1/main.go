package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

// https://en.wikipedia.org/wiki/Taxicab_geometry
func manhattanDistance(a, b Point) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func solve(input []string) int {
	dirs := map[byte]Point{
		'U': {0, -1},
		'D': {0, +1},
		'L': {-1, 0},
		'R': {+1, 0},
	}

	visited := make(map[Point]bool)
	start := Point{0, 0}
	var intersection []Point

	for i, line := range input {
		curr := start

		for s := range strings.SplitSeq(line, ",") {
			dir := s[0]
			n, err := strconv.Atoi(s[1:])
			assert.NoErr(err)
			for range n {
				curr.x += dirs[dir].x
				curr.y += dirs[dir].y
				if i == 0 {
					visited[curr] = true
				}
				if i == 1 && visited[curr] {
					intersection = append(intersection, curr)
				}
			}
		}
	}

	result := 1_000_000
	for _, p := range intersection {
		result = min(result, manhattanDistance(start, p))
	}

	return result
}

func main() {
	start := time.Now()
	input := util.Input()

	fmt.Println("Result:", solve(input))
	fmt.Println(time.Since(start))
}
