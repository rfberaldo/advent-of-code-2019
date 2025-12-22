package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Walker struct {
	Point
	steps int
}

func solve(input []string) int {
	dirs := map[byte]Point{
		'U': {0, -1},
		'D': {0, +1},
		'L': {-1, 0},
		'R': {+1, 0},
	}

	stepsByPoint := make(map[Point]int)
	start := Point{0, 0}
	var intersections []int

	for i, line := range input {
		curr := Walker{start, 0}

		for cmd := range strings.SplitSeq(line, ",") {
			dir := cmd[0]
			n, err := strconv.Atoi(cmd[1:])
			assert.NoErr(err)

			for range n {
				curr.x += dirs[dir].x
				curr.y += dirs[dir].y
				curr.steps++

				if i == 0 && stepsByPoint[curr.Point] == 0 {
					stepsByPoint[curr.Point] = curr.steps
				}

				if i == 1 && stepsByPoint[curr.Point] > 0 {
					intersections = append(intersections, curr.steps+stepsByPoint[curr.Point])
				}
			}
		}
	}

	assert.True(len(intersections) > 0, "no intersections found!")
	return slices.Min(intersections)
}

func main() {
	start := time.Now()
	input := util.Input()

	fmt.Println("Result:", solve(input))
	fmt.Println(time.Since(start))
}
