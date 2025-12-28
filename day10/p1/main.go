package main

import (
	"aoc2019/lib/util"
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
	angles []float64
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

func solve() int {
	points := parse()

	var asteroids []Asteroid
	for i, a := range points {
		asteroid := Asteroid{a, nil}
		for j, b := range points {
			if i == j {
				continue
			}
			angle := angleBetween(a, b)
			asteroid.angles = append(asteroid.angles, angle)
		}
		asteroids = append(asteroids, asteroid)
	}

	result := 0

	// remove duplicate angles and get the highest length
	for _, v := range asteroids {
		slices.Sort(v.angles)
		v.angles = slices.Compact(v.angles)
		result = max(result, len(v.angles))
	}

	return result
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
