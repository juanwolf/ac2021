package main

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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

func TestGroupMeasurements(t *testing.T) {
	tests := []struct {
		m    []int
		w    int
		want []int
	}{
		{m: []int{1}, w: 1, want: []int{1}},
		{m: []int{0, 1}, w: 2, want: []int{1}},
		{m: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, w: 3, want: []int{607, 618, 618, 617, 647, 716, 769, 792}},
	}

	for _, test := range tests {
		got := GroupMeasurements(test.m, test.w)
		if !Equal(got, test.want) {
			t.Errorf("Expected %d, got: %d", test.want, got)
		}
	}

}
