package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input []string) int {
	var results []int

	bef, aft, ok := strings.Cut(input[0], "-")
	assert.True(ok)
	init, err := strconv.Atoi(bef)
	assert.NoErr(err)
	end, err := strconv.Atoi(aft)
	assert.NoErr(err)

outer:
	for num := init; num <= end; num++ {
		numStr := strconv.Itoa(num)
		hasAdj := false
		for i := 0; i < len(numStr)-1; i++ {
			if numStr[i] == numStr[i+1] {
				hasAdj = true
			}
			if numStr[i] > numStr[i+1] {
				continue outer
			}
		}

		if hasAdj {
			results = append(results, num)
		}
	}

	return len(results)
}

func main() {
	start := time.Now()
	input := util.Input()

	fmt.Println("Result:", solve(input))
	fmt.Println(time.Since(start))
}
