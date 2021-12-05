package main

import (
	"testing"
)

func Equal(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func MatrixEqual(a1, a2 [][]int) bool {
	for i, a := range a1 {
		if !Equal(a, a2[i]) {
			return false
		}

	}
	return true
}

func TestGetColumn(t *testing.T) {
	tests := []struct {
		b      Board
		column int
		want   []int
	}{
		{b: Board{{1, 2}, {3, 4}}, column: 0, want: []int{1, 3}},
		{b: Board{{1, 2}, {3, 4}}, column: 1, want: []int{2, 4}},
	}

	for _, test := range tests {
		got := test.b.getColumn(test.column)
		if !Equal(got, test.want) {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

func TestGetRow(t *testing.T) {
	tests := []struct {
		b      Board
		column int
		want   []int
	}{
		{b: Board{{1, 2}, {3, 4}}, column: 0, want: []int{1, 2}},
		{b: Board{{1, 2}, {3, 4}}, column: 1, want: []int{3, 4}},
	}

	for _, test := range tests {
		got := test.b.getRow(test.column)
		if !Equal(got, test.want) {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}

}

func TestGetWinningCombination(t *testing.T) {
	tests := []struct {
		b    Board
		want [][]int
	}{
		{b: Board{{1, 2}, {3, 4}}, want: [][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}}},
	}

	for _, test := range tests {
		got := test.b.getWinningCombinations()
		if !MatrixEqual(got, test.want) {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}

}

func TestIsWinning(t *testing.T) {
	b := Board{{1, 2}, {3, 4}}
	draw := []int{1, 5, 3, 2}

	winning, turn := b.isWinning(draw)
	if (winning == false) || (turn != 2) {
		t.Error("winning", winning, "turn", turn)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		a1   []int
		a2   []int
		want bool
	}{
		{a1: []int{2, 1}, a2: []int{1, 2, 3, 4}, want: true},
		{a1: []int{2, 5}, a2: []int{1, 2, 3, 4}, want: false},
	}

	for _, test := range tests {
		got := Contains(test.a1, test.a2)
		if got != test.want {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

func TestRemoveOccurences(t *testing.T) {
	tests := []struct {
		a1   []int
		a2   []int
		want []int
	}{
		{a1: []int{2, 1}, a2: []int{1, 2, 3, 4}, want: []int{3, 4}},
		{a1: []int{2, 5}, a2: []int{1, 2, 3, 4}, want: []int{1, 3, 4}},
		{a1: []int{2, 5}, a2: []int{1}, want: []int{1}},
		{a1: []int{2, 5, 6}, a2: []int{1, 2}, want: []int{1}},
	}

	for _, test := range tests {
		got := RemoveOccurences(test.a2, test.a1)
		if !Equal(got, test.want) {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}
