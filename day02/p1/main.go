package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func run(pgrm []int) int {
	for i := 0; ; i += 4 {
		op := pgrm[i]
		if op == 99 {
			return pgrm[0]
		}

		ai := pgrm[i+1]
		bi := pgrm[i+2]
		ri := pgrm[i+3]

		switch op {
		case 1:
			pgrm[ri] = pgrm[ai] + pgrm[bi]
			// fmt.Printf("%v[%v] + %v[%v] = %v[%v]\n", pgrm[ai], ai, pgrm[bi], bi, pgrm[ri], ri)
		case 2:
			pgrm[ri] = pgrm[ai] * pgrm[bi]
			// fmt.Printf("%v[%v] * %v[%v] = %v[%v]\n", pgrm[ai], ai, pgrm[bi], bi, pgrm[ri], ri)
		}
	}
}

func main() {
	start := time.Now()
	input := util.Input()

	var pgrm []int
	for _, line := range input {
		for s := range strings.SplitSeq(line, ",") {
			n, err := strconv.Atoi(s)
			assert.NoErr(err)
			pgrm = append(pgrm, n)
		}
	}

	pgrm[1] = 12
	pgrm[2] = 2
	result := run(pgrm)

	fmt.Println("Result:", result)
	fmt.Println(time.Since(start))
}
