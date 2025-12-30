package main

import (
	"aoc2019/day13/intcode"
	"aoc2019/lib/assert"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"

	"golang.org/x/term"
)

func printGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, s := range row {
			fmt.Print(s)
		}
		fmt.Print("\r\n")
	}
}

func clearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func setupTerm() (restore func()) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	assert.NoErr(err)
	return func() { term.Restore(fd, oldState) }
}

type Key string

const (
	KeyArrowLeft  Key = "\x1b[D"
	KeyArrowRight Key = "\x1b[C"
	KeyQ          Key = "q\x00\x00"
)

func readKey() Key {
	buf := make([]byte, 3)
	_, err := os.Stdin.Read(buf)
	assert.NoErr(err)
	return Key(buf)
}

func play() int {
	grid := makeGrid(44, 21, " ")
	pgrm := parse()
	pgrm[0] = 2
	ic := intcode.New(pgrm)
	ic.Run()

	restore := setupTerm()
	defer restore()

	score := 0
	for ; !ic.Done(); ic.Run() {
		for step := range slices.Chunk(ic.Output(), 3) {
			x, y, tile := step[0], step[1], Tile(step[2])
			if x == -1 {
				score = int(tile)
				continue
			}
			grid[y][x] = tile.String()
		}

		ic.ClearOutput()
		clearTerm()
		printGrid(grid)

		key := readKey()
		if key == KeyQ {
			break
		} else if key == KeyArrowLeft {
			ic.AddInput(-1)
		} else if key == KeyArrowRight {
			ic.AddInput(1)
		} else {
			ic.AddInput(0)
		}

		time.Sleep(20 * time.Microsecond)
	}

	return score
}
