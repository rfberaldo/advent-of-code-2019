package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"time"
)

func fuel(n int) int {
	n = n/3 - 2
	if n <= 0 {
		return 0
	}
	return n + fuel(n)
}

func main() {
	start := time.Now()
	input := util.Input()

	result := 0

	for _, line := range input {
		n, err := strconv.Atoi(line)
		assert.NoErr(err)
		result += fuel(n)
	}

	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
