package intcode

import (
	"slices"
	"testing"
)

func TestIntCode(t *testing.T) {
	tests := []struct {
		pgrm   []int
		expect []int
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
			pgrm:   []int{3, 1, 4, 1, 99},
			expect: []int{3, 69, 4, 1, 99},
		},
	}

	for _, tc := range tests {
		ic := New(tc.pgrm)
		input := slices.Clone(tc.pgrm)
		got := ic.Run()
		if !slices.Equal(got, tc.expect) {
			t.Fatalf("\n  input: %v\n   want: %v\n    got: %v", input, tc.expect, got)
		}
	}
}
