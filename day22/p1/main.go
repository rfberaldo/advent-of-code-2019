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

func newDeck() []int {
	deck := make([]int, 10007)
	for i := range len(deck) {
		deck[i] = i
	}
	return deck
}

func cut(deck []int, n int) {
	if n < 0 {
		n = len(deck) + n
	}

	if n == 0 {
		return
	}

	slices.Reverse(deck[:n])
	slices.Reverse(deck[n:])
	slices.Reverse(deck)
}

func dealInc(deck []int, n int) []int {
	deck2 := make([]int, len(deck))
	for i, j := 0, 0; i < len(deck); i, j = i+1, j+n {
		j %= len(deck)
		deck2[j] = deck[i]
	}
	return deck2
}

func solve() int {
	deck := newDeck()

	for _, line := range util.Input() {
		if line == "deal into new stack" {
			slices.Reverse(deck)
			continue
		}

		fields := strings.Fields(line)
		s := fields[len(fields)-1]
		n, err := strconv.Atoi(s)
		assert.NoErr(err)

		switch fields[0] {
		case "deal":
			deck = dealInc(deck, n)

		case "cut":
			cut(deck, n)
		}
	}

	return slices.Index(deck, 2019)
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
