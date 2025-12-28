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
		{
			// Quine - takes no input and produces a copy of itself as output.
			pgrm:         []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			expectOutput: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			// Output a 16-digit number.
			pgrm:         []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			expectOutput: []int{1219070632396864},
		},
		{
			// Output the large number in the middle.
			pgrm:         []int{104, 1125899906842624, 99},
			expectOutput: []int{1125899906842624},
		},
		{
			// Input relative mode.
			pgrm:       []int{109, 1, 203, 4, 99},
			expect:     []int{109, 1, 203, 4, 99, 42},
			inputQueue: []int{42},
		},
	}

	for _, tc := range tests {
		ic := New(tc.pgrm)
		ic.inputQueue = tc.inputQueue
		ic.Run()

		if !ic.done {
			t.Fatalf("input: %v not done", tc.pgrm)
		}

		if len(tc.expect) > 0 && !slices.Equal(ic.pgrm, tc.expect) {
			t.Fatalf("\ninput: %v\n want: %v\n  got: %v", tc.pgrm, tc.expect, ic.pgrm)
		}

		if !slices.Equal(ic.output, tc.expectOutput) {
			t.Fatalf("\n      input: %v\nwant output: %v\n got output: %v", tc.pgrm, tc.expectOutput, ic.output)
		}
	}
}

func TestIntCode_feedback_loop(t *testing.T) {
	pgrm := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	icA := New(pgrm).SetName("Amp A")
	icB := New(pgrm).SetName("Amp B")
	icC := New(pgrm).SetName("Amp C")
	icD := New(pgrm).SetName("Amp D")
	icE := New(pgrm).SetName("Amp E")

	icA.AddInput(9, 0).AddFeedback(icE)
	icB.AddInput(8).AddFeedback(icA)
	icC.AddInput(7).AddFeedback(icB)
	icD.AddInput(6).AddFeedback(icC)
	icE.AddInput(5).AddFeedback(icD)

	for !icE.done {
		icA.Run()
		icB.Run()
		icC.Run()
		icD.Run()
		icE.Run()
	}

	if !icE.done {
		t.Fatal("expected to be done")
	}

	const expect = 139629729
	if icE.LastOutput() != expect {
		t.Fatalf("\nwant: %v\n got: %v", expect, icE.LastOutput())
	}
}

func TestIntCode_length_expand(t *testing.T) {
	pgrm := []int{1, 0, 0, 9, 99}
	expect := []int{1, 0, 0, 9, 99, 0, 0, 0, 0, 2}

	ic := New(pgrm).SetDebug(true)
	ic.Run()

	if !ic.done {
		t.Fatalf("not done")
	}

	if !slices.Equal(ic.pgrm, expect) {
		t.Fatalf("\n  want: %v\n   got: %v", expect, ic.pgrm)
	}
}

func TestIntCode_quine(t *testing.T) {
	pgrm := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	ic := New(pgrm).SetDebug(true)
	ic.Run()

	if !ic.done {
		t.Fatalf("not done")
	}

	if !slices.Equal(ic.output, pgrm) {
		t.Fatalf("\n  want: %v\n   got: %v", pgrm, ic.output)
	}
}
