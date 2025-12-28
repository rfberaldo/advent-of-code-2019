package main

import (
	"aoc2019/lib/util"
	"cmp"
	"fmt"
	"math"
	"slices"
	"time"
)

type Point struct {
	x, y int
}

type Asteroid struct {
	Point
	dist  int
	angle float64
}

type Laser struct {
	Point
	asteroids []Asteroid
}

func parse() []Point {
	input := util.Input()
	var points []Point
	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func angleBetween(a, b Point) float64 {
	return math.Atan2(float64(b.y-a.y), float64(b.x-a.x))
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

func solve() int {
	points := parse()

	const (
		deg90  = math.Pi / 2 // 90deg in radians
		deg360 = math.Pi * 2 // 360deg in radians
	)

	var lasers []Laser
	for i, a := range points {
		laser := Laser{a, nil}
		for j, b := range points {
			if i == j {
				continue
			}

			angle := angleBetween(a, b)
			angle += deg90 // add 90deg so when it's sorted it starts from top
			if angle < 0 {
				angle += deg360 // keep angles 0-360
			}
			dist := manhattanDistance(a, b)
			ast := Asteroid{b, dist, angle}

			laser.asteroids = append(laser.asteroids, ast)
		}
		lasers = append(lasers, laser)
	}

	var bestLaser Laser
	var count = 0

	for _, laser := range lasers {
		var angles []float64
		for _, ast := range laser.asteroids {
			angles = append(angles, ast.angle)
		}
		slices.Sort(angles)
		angles = slices.Compact(angles)
		if len(angles) > count {
			bestLaser = laser
			count = len(angles)
		}
	}

	// sort by angle then distance
	slices.SortFunc(bestLaser.asteroids, func(a, b Asteroid) int {
		if c := cmp.Compare(a.angle, b.angle); c != 0 {
			return c
		}
		return cmp.Compare(a.dist, b.dist)
	})

	destroyed := make(map[Point]bool)
	lastAngle := -10.0
	for len(destroyed) < len(bestLaser.asteroids) {
		for _, ast := range bestLaser.asteroids {
			if destroyed[ast.Point] || ast.angle == lastAngle {
				continue
			}

			destroyed[ast.Point] = true
			lastAngle = ast.angle

			if len(destroyed) == 200 {
				return ast.x*100 + ast.y
			}
		}
	}

	panic("unreacheable")
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
