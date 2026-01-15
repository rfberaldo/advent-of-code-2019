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

func mod(x, y int) int {
	r := x % y
	if r < 0 {
		r += y
	}
	return r
}

// https://codeforces.com/blog/entry/72593
func solve_congruential() int {
	const deckSize = 10007

	a, b := 1, 0
	for _, line := range util.Input() {
		c, d := -1, -1 // deal into new stack

		fields := strings.Fields(line)
		switch {
		case fields[0] == "cut":
			n, err := strconv.Atoi(fields[len(fields)-1])
			assert.NoErr(err)
			c, d = 1, -n

		case fields[2] == "increment":
			n, err := strconv.Atoi(fields[len(fields)-1])
			assert.NoErr(err)
			c, d = n, 0
		}

		// compose operation between two linear congruential functions
		// (a,b) ; (c,d) = (ac mod m, bc+d mod m)
		a = mod(a*c, deckSize)
		b = mod(b*c+d, deckSize)
	}

	// linear congruential function, where m = deckSize
	// f(x) = ax + b mod m
	return mod(a*2019+b, deckSize)
}

func main() {
	start := time.Now()
	fmt.Println("Result (slices):", solve())
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("Result (congruential):", solve_congruential())
	fmt.Println(time.Since(start))
}
