package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parse() []int {
	input := util.Input()
	var numbers []int
	for s := range strings.SplitSeq(input[0], "") {
		n, err := strconv.Atoi(s)
		assert.NoErr(err)
		numbers = append(numbers, n)
	}
	return numbers
}

var defaultPattern = []int{0, 1, 0, -1}

func patternValue(i, loop int) int {
	i = (i / loop) % len(defaultPattern)
	return defaultPattern[i]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solve() string {
	const phases = 100
	numbers := parse()

	for phase := 1; phase <= phases; phase++ {
		for i := range numbers {
			sum := 0
			for j := i; j < len(numbers); j++ {
				sum += numbers[j] * patternValue(j+1, i+1)
			}
			numbers[i] = abs(sum) % 10
		}
	}

	result := ""
	for _, n := range numbers[:8] {
		result += strconv.Itoa(n)
	}
	return result
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
