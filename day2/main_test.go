package main

import (
	"testing"
)

func TestNewInstruction(t *testing.T) {
	tests := []struct {
		i    string
		want Instruction
	}{
		{i: "forward 5", want: Instruction{command: "forward", unit: 5}},
		{i: "down 1", want: Instruction{command: "down", unit: 1}},
	}

	for _, test := range tests {
		got := NewInstruction(test.i)
		if !test.want.Equal(got) {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

// Testing out given/when/then in a table test case scenario.
func TestMove(t *testing.T) {
	tests := []struct {
		given Location
		when  Instruction
		then  Location
	}{
		{given: Location{x: 0, y: 0}, when: Instruction{command: "forward", unit: 5}, then: Location{x: 5, y: 0}},
		{given: Location{x: 0, y: 0}, when: Instruction{command: "down", unit: 1}, then: Location{x: 0, y: 1}},
		{given: Location{x: 0, y: 3}, when: Instruction{command: "up", unit: 3}, then: Location{x: 0, y: 0}},
		{given: Location{x: 5, y: 5}, when: Instruction{command: "up", unit: 3}, then: Location{x: 5, y: 2}},
		{given: Location{x: 0, y: 5}, when: Instruction{command: "down", unit: 3}, then: Location{x: 0, y: 8}},
		{given: Location{x: 10, y: 0}, when: Instruction{command: "forward", unit: 3}, then: Location{x: 13, y: 0}},
	}

	for _, test := range tests {
		got := test.given.Move(test.when)
		if !test.then.Equal(got) {
			t.Errorf("Expected %+v, got: %+v", test.then, got)
		}
	}
}

func TestSubmarineMove(t *testing.T) {
	tests := []struct {
		given *Submarine
		when  Instruction
		then  *Submarine
	}{
		{given: NewSubmarine(0, 0, 0), when: NewInstruction("down 5"), then: NewSubmarine(0, 0, 5)},
		{given: NewSubmarine(0, 0, 5), when: NewInstruction("up 5"), then: NewSubmarine(0, 0, 0)},
		{given: NewSubmarine(0, 0, 0), when: NewInstruction("forward 5"), then: NewSubmarine(5, 0, 0)},
		{given: NewSubmarine(0, 0, 1), when: NewInstruction("forward 5"), then: NewSubmarine(5, 5, 1)},
		{given: NewSubmarine(0, 0, 2), when: NewInstruction("forward 5"), then: NewSubmarine(5, 10, 2)},
	}
	for _, test := range tests {
		test.given.Move(test.when)
		if !test.then.Equal(*test.given) {
			t.Errorf("Expected %+v, got: %+v", test.then, test.given)
		}
	}
}
