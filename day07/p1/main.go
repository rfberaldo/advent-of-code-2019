package main

import (
	"aoc2019/day07/intcode"
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permutations(s []int) [][]int {
	var helper func(int)
	var res [][]int

	helper = func(n int) {
		if n == 1 {
			res = append(res, slices.Clone(s))
			return
		}

		helper(n - 1)
		for i := 0; i < n-1; i++ {
			if n%2 == 0 {
				s[n-1], s[i] = s[i], s[n-1]
			} else {
				s[n-1], s[0] = s[0], s[n-1]
			}
			helper(n - 1)
		}
	}

	helper(len(s))
	return res
}

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

func main() {
	start := time.Now()

	pgrm := parse()
	orders := permutations([]int{0, 1, 2, 3, 4})

	run := func(n ...int) int {
		ic := intcode.New(slices.Clone(pgrm))
		ic.AddInput(n...)
		return ic.Run()[0]
	}

	result := 0
	for _, order := range orders {
		output := run(order[0], 0)
		output = run(order[1], output)
		output = run(order[2], output)
		output = run(order[3], output)
		output = run(order[4], output)
		result = max(result, output)
	}

	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
