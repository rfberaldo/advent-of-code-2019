package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func mod(x, y *big.Int) *big.Int {
	return new(big.Int).Mod(x, y)
}

func modInv(x, y *big.Int) *big.Int {
	return new(big.Int).ModInverse(x, y)
}

func exp(x, y, m *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, m)
}

var deckSize = big.NewInt(119315717514047)
var shuffles = big.NewInt(101741582076661)

func shuffleN(a, b *big.Int) (*big.Int, *big.Int) {
	// naive method:
	// for range shuffles {
	// 	b, c := a, b
	// 	a, b = 1, 0
	// 	a = mod(a*c, deckSize)
	// 	b = mod(b*c+d, deckSize)
	// }

	a2 := exp(a, shuffles, deckSize)

	// b2 := mod(b*(a2-1)*modInv(a-1, deckSize), deckSize)
	b2 := mod(
		mul(
			mul(b, add(a2, big.NewInt(-1))),
			modInv(add(a, big.NewInt(-1)), deckSize),
		),
		deckSize,
	)
	return a2, b2
}

// https://codeforces.com/blog/entry/72593
func solve() int64 {
	a, b := big.NewInt(1), big.NewInt(0)
	for _, line := range util.Input() {
		c, d := big.NewInt(-1), big.NewInt(-1) // deal into new stack

		fields := strings.Fields(line)
		switch {
		case fields[0] == "cut":
			n, err := strconv.Atoi(fields[len(fields)-1])
			assert.NoErr(err)
			c.SetInt64(1)
			d.SetInt64(int64(-n))

		case fields[2] == "increment":
			n, err := strconv.Atoi(fields[len(fields)-1])
			assert.NoErr(err)
			c.SetInt64(int64(n))
			d.SetInt64(0)
		}

		// compose operation between two linear congruential functions
		// (a,b) ; (c,d) = (ac mod m, bc+d mod m)
		a = mod(mul(a, c), deckSize)         // a = mod(a*c, deckSize)
		b = mod(add(mul(b, c), d), deckSize) // b = mod(b*c+d, deckSize)
	}

	a, b = shuffleN(a, b)

	// r := mod((2020-b)*modInv(a, deckSize), deckSize)
	r := mod(
		mul(
			sub(big.NewInt(2020), b),
			modInv(a, deckSize),
		),
		deckSize,
	)

	assert.True(r.IsInt64())
	return r.Int64()
}

func main() {
	start := time.Now()
	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
