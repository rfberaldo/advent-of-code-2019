package intcode

import (
	"slices"
	"testing"
)

func TestIntCode(t *testing.T) {
	tests := []struct {
		pgrm         []int
		expect       []int
		inputQueue   []int
		expectOutput []int
	}{
		{
			pgrm:   []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expect: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			pgrm:   []int{1, 0, 0, 0, 99},
			expect: []int{2, 0, 0, 0, 99},
		},
		{
			pgrm:   []int{2, 3, 0, 3, 99},
			expect: []int{2, 3, 0, 6, 99},
		},
		{
			pgrm:   []int{2, 4, 4, 5, 99, 0},
			expect: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			pgrm:   []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expect: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
		{
			pgrm:   []int{1002, 4, 3, 4, 33},
			expect: []int{1002, 4, 3, 4, 99},
		},
		{
			pgrm:   []int{1101, 100, -1, 4, 0},
			expect: []int{1101, 100, -1, 4, 99},
		},
		{
			pgrm:         []int{3, 1, 4, 1, 99},
			expect:       []int{3, 69, 4, 1, 99},
			inputQueue:   []int{69},
			expectOutput: []int{69},
		},
		{
			// Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputQueue:   []int{8},
			expectOutput: []int{1},
		},
		{
			// Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputQueue:   []int{7},
			expectOutput: []int{0},
		},
		{
			// Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputQueue:   []int{7},
			expectOutput: []int{1},
		},
		{
			// Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputQueue:   []int{8},
			expectOutput: []int{0},
		},
		{
			// Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputQueue:   []int{8},
			expectOutput: []int{1},
		},
		{
			// Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputQueue:   []int{7},
			expectOutput: []int{0},
		},
		{
			// Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputQueue:   []int{7},
			expectOutput: []int{1},
		},
		{
			// Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			pgrm:         []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputQueue:   []int{8},
			expectOutput: []int{0},
		},
		{
			// Using position mode, jump test that take an input, then output 0 if the input was zero or 1 if the input was non-zero.
			pgrm:         []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputQueue:   []int{1},
			expectOutput: []int{1},
		},
		{
			// Using position mode, jump test that take an input, then output 0 if the input was zero or 1 if the input was non-zero.
			pgrm:         []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputQueue:   []int{0},
			expectOutput: []int{0},
		},
		{
			// Using immediate mode, jump test that take an input, then output 0 if the input was zero or 1 if the input was non-zero.
			pgrm:         []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputQueue:   []int{1},
			expectOutput: []int{1},
		},
		{
			// Using immediate mode, jump test that take an input, then output 0 if the input was zero or 1 if the input was non-zero.
			pgrm:         []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputQueue:   []int{0},
			expectOutput: []int{0},
		},
		{
			// The program will then output 999 if the input value is below 8,
			// output 1000 if the input value is equal to 8,
			// or output 1001 if the input value is greater than 8.
			pgrm:         []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputQueue:   []int{7},
			expectOutput: []int{999},
		},
		{
			// The program will then output 999 if the input value is below 8,
			// output 1000 if the input value is equal to 8,
			// or output 1001 if the input value is greater than 8.
			pgrm:         []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputQueue:   []int{8},
			expectOutput: []int{1000},
		},
		{
			// The program will then output 999 if the input value is below 8,
			// output 1000 if the input value is equal to 8,
			// or output 1001 if the input value is greater than 8.
			pgrm:         []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputQueue:   []int{9},
			expectOutput: []int{1001},
		},
	}

	for _, tc := range tests {
		ic := New(tc.pgrm)
		ic.inputQueue = tc.inputQueue
		input := slices.Clone(tc.pgrm)
		output := ic.Run()

		if len(tc.expect) > 0 && !slices.Equal(ic.pgrm, tc.expect) {
			t.Fatalf("\n  input: %v\n   want: %v\n    got: %v", input, tc.expect, ic.pgrm)
		}

		if !slices.Equal(output, tc.expectOutput) {
			t.Fatalf("\n  input: %v\n   want output: %v\n    got output: %v", input, tc.expectOutput, output)
		}
	}
}
