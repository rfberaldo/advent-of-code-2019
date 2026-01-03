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

func makeGridMap(s []int) (map[Point]byte, Robot) {
	var dirs = map[byte]Direction{
		'^': {0, -1},
		'v': {0, +1},
		'<': {-1, 0},
		'>': {+1, 0},
	}

	grid := make(map[Point]byte)
	p := Point{}
	init := Robot{}
	for i := range s {
		b := byte(s[i])
		if b == '^' || b == 'v' || b == '<' || b == '>' {
			init = Robot{p, dirs[b]}
			grid[p] = '#'
			p.x++
			continue
		}
		if b == '\n' {
			p.x = 0
			p.y++
			continue
		}
		grid[p] = b
		p.x++
	}

	return grid, init
}

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Robot struct {
	Point
	Direction
}

// https://en.wikipedia.org/wiki/Rotation_matrix
func (p Robot) cw() Robot {
	p.dx, p.dy = -p.dy, p.dx
	return p
}

// https://en.wikipedia.org/wiki/Rotation_matrix
func (p Robot) ccw() Robot {
	p.dx, p.dy = p.dy, -p.dx
	return p
}

func (p Robot) fwd(steps int) Robot {
	p.x += p.dx * steps
	p.y += p.dy * steps
	return p
}

func toInt(s string) []int {
	ss := make([]int, len(s))
	for i := range s {
		ss[i] = int(s[i])
	}
	return ss
}

func solve() int {
	pgrm := parse()
	ic := intcode.New(pgrm)
	ic.Run()
	assert.True(ic.Done())

	grid, start := makeGridMap(ic.Output())

	var walkFwd func(Robot) int
	walkFwd = func(curr Robot) int {
		curr = curr.fwd(1)
		if grid[curr.Point] != '#' {
			return 0
		}

		return 1 + walkFwd(curr)
	}

	hasPath := func(curr Robot) bool {
		return grid[curr.fwd(1).Point] == '#' ||
			grid[curr.ccw().fwd(1).Point] == '#' ||
			grid[curr.cw().fwd(1).Point] == '#'
	}

	var walk func(Robot) []byte
	walk = func(curr Robot) []byte {
		var path []byte

		if !hasPath(curr) {
			return path
		}

		// forward
		if steps := walkFwd(curr); steps > 0 {
			curr = curr.fwd(steps)
			path = append(path, []byte(strconv.Itoa(steps))...)
			path = append(path, ',')
		}

		// left
		if steps := walkFwd(curr.ccw()); steps > 0 {
			curr = curr.ccw().fwd(steps)
			path = append(path, 'L')
			path = append(path, ',')
			path = append(path, []byte(strconv.Itoa(steps))...)
			path = append(path, ',')
		}

		// right
		if steps := walkFwd(curr.cw()); steps > 0 {
			curr = curr.cw().fwd(steps)
			path = append(path, 'R')
			path = append(path, ',')
			path = append(path, []byte(strconv.Itoa(steps))...)
			path = append(path, ',')
		}

		return append(path, walk(curr)...)
	}

	fmt.Println(string(walk(start)))

	pgrm[0] = 2
	ic = intcode.New(pgrm)

	// handmande...
	ic.AddInput(toInt("A,B,A,C,B,C,B,C,A,C\n")...)
	ic.AddInput(toInt("R,12,L,10,R,12\n")...)
	ic.AddInput(toInt("L,8,R,10,R,6\n")...)
	ic.AddInput(toInt("R,12,L,10,R,10,L,8\n")...)
	ic.AddInput('n', '\n')
	ic.Run()
	assert.True(ic.Done())
	return ic.LastOutput()
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
