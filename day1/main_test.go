package main

import (
	"testing"
)

func TestNumberDepthIncreased(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{0}, want: 0},
		{input: []int{0, 1}, want: 1},
		{input: []int{0, 1, 0}, want: 1},
		{input: []int{0, 1, 2}, want: 2},
		{input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, want: 7},
	}

	for _, test := range tests {
		got := NumberDepthIncreased(test.input)
		if got != test.want {
			t.Errorf("Expected %d, got: %d", test.want, got)
		}
	}

}
