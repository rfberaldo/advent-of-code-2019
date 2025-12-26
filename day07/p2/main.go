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
	orders := permutations([]int{5, 6, 7, 8, 9})

	result := 0
	for _, order := range orders {
		icA := intcode.New(pgrm).SetName("Amp A")
		icB := intcode.New(pgrm).SetName("Amp B")
		icC := intcode.New(pgrm).SetName("Amp C")
		icD := intcode.New(pgrm).SetName("Amp D")
		icE := intcode.New(pgrm).SetName("Amp E")

		icA.AddInput(order[0], 0).AddFeedback(icE)
		icB.AddInput(order[1]).AddFeedback(icA)
		icC.AddInput(order[2]).AddFeedback(icB)
		icD.AddInput(order[3]).AddFeedback(icC)
		icE.AddInput(order[4]).AddFeedback(icD)

		for !icE.Done() {
			icA.Run()
			icB.Run()
			icC.Run()
			icD.Run()
			icE.Run()
		}

		result = max(result, icE.LastOutput())
	}

	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
