package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strings"
	"time"
)

var nodes = make(map[string][]string)
var visited = make(map[string]bool)

func walk(key string) (int, bool) {
	if visited[key] {
		return 0, false
	}

	visited[key] = true

	if key == "" {
		return 0, false
	}

	if key == "SAN" {
		return 0, true
	}

	for _, k := range nodes[key] {
		if n, ok := walk(k); ok {
			return 1 + n, true
		}
	}

	return 0, false
}

func solve(input []string) int {
	for _, line := range input {
		parent, child, ok := strings.Cut(line, ")")
		assert.True(ok)

		nodes[child] = append(nodes[child], parent)
		nodes[parent] = append(nodes[parent], child)
	}

	n, ok := walk("YOU")
	assert.True(ok)
	return n - 2
}

func main() {
	start := time.Now()
	input := util.Input()

	fmt.Println("Result:", solve(input))
	fmt.Println(time.Since(start))
}
