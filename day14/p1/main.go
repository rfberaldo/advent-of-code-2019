package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Material struct {
	name  string
	count int
}

type Recipe struct {
	output Material
	input  []Material
}

var recipeByMatName = make(map[string]Recipe)

func parse() {
	parseMat := func(s string) Material {
		fields := strings.Fields(s)
		n, err := strconv.Atoi(fields[0])
		assert.NoErr(err)
		return Material{fields[1], n}
	}

	input := util.Input()
	for _, line := range input {
		in, out, ok := strings.Cut(line, "=>")
		assert.True(ok)
		recipe := Recipe{output: parseMat(out)}
		for field := range strings.SplitSeq(in, ",") {
			recipe.input = append(recipe.input, parseMat(field))
		}
		recipeByMatName[recipe.output.name] = recipe
	}
}

func ceilDiv(n, d int) int {
	return int(math.Ceil(float64(n) / float64(d)))
}

var bank = make(map[string]int)

func solve(curr string, need int) int {
	if curr == "ORE" {
		return need
	}

	recipe := recipeByMatName[curr]

	if bank[curr] >= need {
		bank[curr] -= need
		return 0
	}

	need -= bank[curr]
	times := max(1, ceilDiv(need, recipe.output.count))
	bank[curr] = recipe.output.count*times - need

	cost := 0
	for _, mat := range recipe.input {
		cost += solve(mat.name, times*mat.count)
	}

	return cost
}

func main() {
	start := time.Now()
	parse()
	fmt.Println("Result:", solve("FUEL", 1))
	fmt.Println(time.Since(start))
}
