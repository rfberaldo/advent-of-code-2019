package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func isValid(str string) bool {
	digitCount := make(map[byte]int)
	for i := 0; i < len(str)-1; i++ {
		digitCount[str[i]]++
		if str[i] > str[i+1] {
			return false
		}
	}
	digitCount[str[len(str)-1]]++ // last digit

	for _, v := range digitCount {
		if v == 2 {
			return true
		}
	}

	return false
}

func solve(input []string) int {
	var results []int

	bef, aft, ok := strings.Cut(input[0], "-")
	assert.True(ok)
	init, err := strconv.Atoi(bef)
	assert.NoErr(err)
	end, err := strconv.Atoi(aft)
	assert.NoErr(err)

	for num := init; num <= end; num++ {
		if isValid(strconv.Itoa(num)) {
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
