package main

import (
	"aoc2019/day05/intcode"
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

	ic := intcode.New(pgrm)
	ic.AddInput(5)
	ic.Run()

	fmt.Println(time.Since(start))
}
