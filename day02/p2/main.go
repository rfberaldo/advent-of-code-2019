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

const (
	OpAdd  = 1
	OpMul  = 2
	OpExit = 99
)

func run(pgrm []int) int {
	maxIndex := len(pgrm) - 1
	for i := 0; ; i += 4 {
		assert.True(i < maxIndex, "index [", i, "] out of bounds!")

		op := pgrm[i]
		if op == OpExit {
			return pgrm[0]
		}

		assert.True(i+3 < maxIndex, "index [", i, " + 3] out of bounds!")

		ai := pgrm[i+1]
		bi := pgrm[i+2]
		ri := pgrm[i+3]

		switch op {
		case OpAdd:
			pgrm[ri] = pgrm[ai] + pgrm[bi]
			// fmt.Printf("%v[%v] + %v[%v] = %v[%v]\n", pgrm[ai], ai, pgrm[bi], bi, pgrm[ri], ri)
		case OpMul:
			pgrm[ri] = pgrm[ai] * pgrm[bi]
			// fmt.Printf("%v[%v] * %v[%v] = %v[%v]\n", pgrm[ai], ai, pgrm[bi], bi, pgrm[ri], ri)
		}
	}
}

func main() {
	start := time.Now()
	input := util.Input()

	var pgrm1 []int
	for _, line := range input {
		for s := range strings.SplitSeq(line, ",") {
			n, err := strconv.Atoi(s)
			assert.NoErr(err)
			pgrm1 = append(pgrm1, n)
		}
	}

	result := 0

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			pgrm := slices.Clone(pgrm1)
			pgrm[1] = noun
			pgrm[2] = verb
			if run(pgrm) == 19690720 {
				fmt.Printf("noun=%v verb=%v\n", noun, verb)
				result = 100*noun + verb
				goto end
			}
		}
	}

end:
	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
