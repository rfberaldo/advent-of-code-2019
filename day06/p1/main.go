package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strings"
	"time"
)

var nodes = make(map[string]string)

func walk(key string) int {
	if _, ok := nodes[key]; !ok {
		return 0
	}

	return 1 + walk(nodes[key])
}

func solve(input []string) int {
	for _, line := range input {
		parent, child, ok := strings.Cut(line, ")")
		assert.True(ok)

		nodes[child] = parent
	}

	result := 0
	for k := range nodes {
		result += walk(k)
	}
	return result
}

func main() {
	start := time.Now()
	input := util.Input()

	fmt.Println("Result:", solve(input))
	fmt.Println(time.Since(start))
}
