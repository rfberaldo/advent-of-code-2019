package main

// import (
// 	"aoc2019/day11/intcode"
// 	"aoc2019/lib/assert"
// 	"aoc2019/lib/util"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"golang.org/x/term"
// )

// func parse() []int {
// 	input := util.Input()
// 	var pgrm []int
// 	for s := range strings.SplitSeq(input[0], ",") {
// 		n, err := strconv.Atoi(s)
// 		assert.NoErr(err)
// 		pgrm = append(pgrm, n)
// 	}
// 	return pgrm
// }

// func printNearby(m map[Point]string, center Point) {
// 	const width, height = 60, 20

// 	zeroX := center.x - width/2
// 	zeroY := center.y - height/2

// 	for y := range height {
// 		for x := range width {
// 			s, ok := m[Point{zeroX + x, zeroY + y}]
// 			if !ok {
// 				s = " "
// 			}
// 			fmt.Print(s)
// 		}
// 		fmt.Print("\r\n")
// 	}
// }

// func clearTerm() {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

// func setupTerm() (restore func()) {
// 	fd := int(os.Stdin.Fd())
// 	oldState, err := term.MakeRaw(fd)
// 	assert.NoErr(err)
// 	return func() { term.Restore(fd, oldState) }
// }

// type Key string

// const (
// 	KeyArrowUp    = "\x1b[A"
// 	KeyArrowDown  = "\x1b[B"
// 	KeyArrowRight = "\x1b[C"
// 	KeyArrowLeft  = "\x1b[D"
// 	KeyQ          = "q\x00\x00"
// )

// func readKey() Key {
// 	buf := make([]byte, 3)
// 	_, err := os.Stdin.Read(buf)
// 	assert.NoErr(err)
// 	return Key(buf)
// }

// const (
// 	StatusWall  = 0
// 	StatusOK    = 1
// 	StatusFound = 2
// )

// const (
// 	MoveUp    = 1
// 	MoveDown  = 2
// 	MoveLeft  = 3
// 	MoveRight = 4
// )

// var movePoint = map[int]Point{
// 	MoveUp:    {0, -1},
// 	MoveDown:  {0, +1},
// 	MoveLeft:  {-1, 0},
// 	MoveRight: {+1, 0},
// }

// type Point struct {
// 	x, y int
// }

// func (p Point) move(move int) Point {
// 	return Point{p.x + movePoint[move].x, p.y + movePoint[move].y}
// }

// func play() int {
// 	restore := setupTerm()
// 	defer restore()

// 	ic := intcode.New(parse())

// 	droid := Point{}
// 	moves := 0

// 	grid := make(map[Point]string)
// 	grid[droid] = "&"

// 	printNearby(grid, droid)

// 	for {
// 		wantMove := 0

// 		switch readKey() {
// 		case KeyQ:
// 			return 0
// 		case KeyArrowUp:
// 			wantMove = MoveUp
// 		case KeyArrowDown:
// 			wantMove = MoveDown
// 		case KeyArrowLeft:
// 			wantMove = MoveLeft
// 		case KeyArrowRight:
// 			wantMove = MoveRight
// 		}
// 		ic.AddInput(wantMove)

// 		ic.Run()

// 		status := -1
// 		if len(ic.Output()) > 0 {
// 			status = ic.LastOutput()
// 		}
// 		ic.ClearOutput()

// 		update := true
// 		switch status {
// 		case StatusWall:
// 			if grid[droid.move(wantMove)] == "#" {
// 				update = false
// 			}
// 			grid[droid.move(wantMove)] = "#"

// 		case StatusOK:
// 			grid[droid] = "."
// 			droid = droid.move(wantMove)
// 			grid[droid] = "&"
// 			moves++

// 		case StatusFound:
// 			moves++
// 			return moves
// 		}

// 		if update {
// 			clearTerm()
// 			printNearby(grid, droid)
// 		}
// 	}
// }

// func main() {
// 	start := time.Now()

// 	fmt.Println("Result:", play())
// 	fmt.Println(time.Since(start))
// }
