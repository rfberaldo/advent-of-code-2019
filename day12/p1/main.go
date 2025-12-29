package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y, z int
}

type Velocity struct {
	vx, vy, vz int
}

type Moon struct {
	Point
	Velocity
}

func (m *Moon) applyVelocity() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func (m *Moon) total() int {
	return (abs(m.x) + abs(m.y) + abs(m.z)) * (abs(m.vx) + abs(m.vy) + abs(m.vz))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func applyGravity(a, b *Moon) {
	apply := func(get func(m *Moon) int, inc, dec func(m *Moon)) {
		if get(a) == get(b) {
			return
		}
		if get(a) > get(b) {
			dec(a)
			inc(b)
			return
		}
		inc(a)
		dec(b)
	}

	apply(func(m *Moon) int { return m.x }, func(m *Moon) { m.vx++ }, func(m *Moon) { m.vx-- })
	apply(func(m *Moon) int { return m.y }, func(m *Moon) { m.vy++ }, func(m *Moon) { m.vy-- })
	apply(func(m *Moon) int { return m.z }, func(m *Moon) { m.vz++ }, func(m *Moon) { m.vz-- })
}

func parse() []Moon {
	input := util.Input()
	var moons []Moon
	for _, line := range input {
		moon := Moon{}
		line = strings.Trim(line, "<>")
		for field := range strings.SplitSeq(line, ", ") {
			id, val, ok := strings.Cut(field, "=")
			assert.True(ok)

			n, err := strconv.Atoi(val)
			assert.NoErr(err)

			switch id {
			case "x":
				moon.x = n
			case "y":
				moon.y = n
			case "z":
				moon.z = n
			}
		}
		moons = append(moons, moon)
	}
	return moons
}

func solve() int {
	moons := parse()

	for range 1000 {
		for i := 0; i < len(moons)-1; i++ {
			for j := i + 1; j < len(moons); j++ {
				applyGravity(&moons[i], &moons[j])
			}
		}

		for i := range moons {
			moons[i].applyVelocity()
		}
	}

	result := 0
	for _, v := range moons {
		result += v.total()
	}
	return result
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
