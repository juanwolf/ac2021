package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Returns the number of time the depth increased from the report
func NumberDepthIncreased(measurements []int) int {
	nb_increases := 0
	previous_measurement := measurements[0]
	for _, measurement := range measurements {
		if previous_measurement < measurement {
			nb_increases += 1
		}
		previous_measurement = measurement
	}

	return nb_increases
}

func main() {
	inputs := []int{}

	file, err := os.Open("./report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Report contained something else than integer", err)
		}
		inputs = append(inputs, input)

	}

	res := NumberDepthIncreased(inputs)

	log.Println("Result found: ", res)
}
