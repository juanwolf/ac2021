package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const FORWARD_STRING = "forward"
const DOWNARD_STRING = "down"
const UPWARD_STRING = "up"

type Location struct {
	x int
	y int
}

func (l Location) Equal(location Location) bool {
	return l.x == location.x && l.y == location.y
}

func (l Location) Move(i Instruction) Location {
	res := Location{
		x: l.x,
		y: l.y,
	}
	switch i.command {
	case FORWARD_STRING:
		res.x += i.unit
	case DOWNARD_STRING:
		res.y += i.unit
	case UPWARD_STRING:
		res.y -= i.unit
	}

	return res
}

type Submarine struct {
	location Location
	aim      int
}

func NewSubmarine(x, y, aim int) *Submarine {
	return &Submarine{
		location: Location{x: x, y: y},
		aim:      aim,
	}
}

func (s *Submarine) Move(i Instruction) {
	switch i.command {
	case DOWNARD_STRING:
		s.aim += i.unit
	case UPWARD_STRING:
		s.aim -= i.unit
	case FORWARD_STRING:
		s.location = s.location.Move(i)
		s.location.y += s.aim * i.unit
	}
}

func (s Submarine) Equal(s1 Submarine) bool {
	return s.location.Equal(s1.location) && s.aim == s1.aim
}

type Instruction struct {
	command string
	unit    int
}

func (i Instruction) Equal(instruction Instruction) bool {
	return i.command == instruction.command && i.unit == instruction.unit
}

func NewInstruction(s string) Instruction {
	args := strings.Split(s, " ")
	command := args[0]
	unit, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	return Instruction{
		command: command,
		unit:    unit,
	}
}

func main() {
	log.Println("Hello")

	instructions := []Instruction{}

	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := NewInstruction(scanner.Text())
		instructions = append(instructions, instruction)
	}

	currentLocation := Location{}
	submarine := NewSubmarine(0, 0, 0)

	for _, instruction := range instructions {
		currentLocation = currentLocation.Move(instruction)
		submarine.Move(instruction)
	}

	log.Printf("Final location: %+v", currentLocation)
	log.Println("Part 1 response:", currentLocation.x*currentLocation.y)
	log.Println("Part 2: Submarine ", submarine, "Response: ", submarine.location.x*submarine.location.y)

}
