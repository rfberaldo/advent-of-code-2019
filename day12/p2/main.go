package main

import (
	"aoc2019/lib/assert"
	"aoc2019/lib/util"
	"fmt"
	"maps"
	"slices"
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
	init Point
}

func (m *Moon) applyVelocity() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func (m *Moon) isInitByAxis(axis string) bool {
	switch axis {
	case "x":
		return m.x == m.init.x && m.vx == 0

	case "y":
		return m.y == m.init.y && m.vy == 0

	case "z":
		return m.z == m.init.z && m.vz == 0
	}
	panic("wrong axis")
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Greatest Common Divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

// Least Common Multiple
func lcm(a, b int) int {
	assert.True(a != 0 && b != 0, "LCM of zero is undefined")
	return abs(a*b) / gcd(a, b)
}

func sliceLCM(s []int) int {
	assert.True(len(s) > 0, "LCM slice is empty")
	result := 0
	for i, n := range s {
		if i == 0 {
			result = n
			continue
		}
		result = lcm(result, n)
	}
	return result
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
				moon.init.x = n
			case "y":
				moon.y = n
				moon.init.y = n
			case "z":
				moon.z = n
				moon.init.z = n
			}
		}
		moons = append(moons, moon)
	}
	return moons
}

func solve() int {
	moons := parse()
	stepsByAxis := make(map[string]int)

	isCycleByAxis := func(axis string) bool {
		if _, exists := stepsByAxis[axis]; exists {
			return false
		}
		for _, m := range moons {
			if !m.isInitByAxis(axis) {
				return false
			}
		}
		return true
	}

	for step := 1; ; step++ {
		for i := 0; i < len(moons)-1; i++ {
			for j := i + 1; j < len(moons); j++ {
				applyGravity(&moons[i], &moons[j])
			}
		}

		for i := range moons {
			moons[i].applyVelocity()
		}

		if axis := "x"; isCycleByAxis(axis) {
			stepsByAxis[axis] = step
		}
		if axis := "y"; isCycleByAxis(axis) {
			stepsByAxis[axis] = step
		}
		if axis := "z"; isCycleByAxis(axis) {
			stepsByAxis[axis] = step
		}

		if len(stepsByAxis) == 3 {
			break
		}
	}

	return sliceLCM(slices.Collect(maps.Values(stepsByAxis)))
}

func main() {
	start := time.Now()

	fmt.Println("Result:", solve())
	fmt.Println(time.Since(start))
}
