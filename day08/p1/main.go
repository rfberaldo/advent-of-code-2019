package main

import (
	"aoc2019/lib/util"
	"fmt"
	"slices"
	"strings"
	"time"
)

func digitsCount(layer []string) (n0, n1, n2 int) {
	for _, v := range layer {
		switch v {
		case "0":
			n0++
		case "1":
			n1++
		case "2":
			n2++
		}
	}
	return
}

func main() {
	start := time.Now()
	input := util.Input()[0]

	const width = 25
	const height = 6
	const pixelsPerLayer = width * height

	r0 := 100_000
	r1 := 0
	r2 := 0

	for layer := range slices.Chunk(strings.Split(input, ""), pixelsPerLayer) {
		n0, n1, n2 := digitsCount(layer)
		if n0 < r0 {
			r0, r1, r2 = n0, n1, n2
		}
	}

	fmt.Println("Result:", r1*r2)
	fmt.Println(time.Since(start))
}
