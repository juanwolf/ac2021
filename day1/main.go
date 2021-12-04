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

// Return a list of measurement grouped by the size of the window
func GroupMeasurements(measurements []int, windowSize int) []int {
	res := []int{}

	for index := 0; index <= len(measurements)-windowSize; index++ {
		window := measurements[index : index+windowSize]
		windowRes := 0
		for _, element := range window {
			windowRes += element
		}

		res = append(res, windowRes)
	}
	return res
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

	log.Println("Number of Depth increase found: ", res)
	log.Println("Number of depth increase for a 3 measurements window", NumberDepthIncreased(GroupMeasurements(inputs, 3)))
}
