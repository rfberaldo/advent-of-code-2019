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

func join(s []int) string {
	str := ""
	for _, n := range s {
		str += strconv.Itoa(n)
	}
	return str
}

func joinInt(s []int) int {
	str := ""
	for _, n := range s {
		str += strconv.Itoa(n)
	}
	n, err := strconv.Atoi(str)
	assert.NoErr(err)
	return n
}

// This works because the last half of the input is a cumulative sum:
//
// 1  +  0  + -1  +  0  +  1  +  0  + -1  +  0
// 0  +  1  +  1  +  0  +  0  + -1  + -1  +  0
// 0  +  0  +  1  +  1  +  1  +  0  +  0  +  0
// 0  +  0  +  0  +  1  +  1  +  1  +  1  +  0
// 0  +  0  +  0  +  0  +  1  +  1  +  1  +  1 = 4
// 0  +  0  +  0  +  0  +  0  +  1  +  1  +  1 = 3
// 0  +  0  +  0  +  0  +  0  +  0  +  1  +  1 = 2
// 0  +  0  +  0  +  0  +  0  +  0  +  0  +  1 = 1
func solve() string {
	const phases = 100
	const repeat = 10_000
	numbers := parse()
	offset := joinInt(numbers[:7])
	numbers = slices.Repeat(numbers, repeat)

	// this only works when offset is >= len/2
	assert.True(offset >= len(numbers)/2)

	for phase := 1; phase <= phases; phase++ {
		sum := numbers[len(numbers)-1]
		for i := len(numbers) - 2; i >= offset; i-- {
			sum = numbers[i] + sum
			numbers[i] = sum % 10
		}
	}

	return join(numbers[offset : offset+8])
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
