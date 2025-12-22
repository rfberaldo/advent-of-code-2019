package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	input := util.Input()

	result := 0

	for _, line := range input {
		n, err := strconv.Atoi(line)
		assert.NoErr(err)
		result += n/3 - 2
	}

	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
