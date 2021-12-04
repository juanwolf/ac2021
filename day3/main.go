package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Gamma(inputs []string) string {
	res := []byte{}
	for i := 0; i < len(inputs[0]); i++ {
		occurences := map[byte]int{}
		for _, input := range inputs {
			occurences[input[i]] += 1
		}
		maxOccurences := 0
		var byteMostFound byte
		for k, v := range occurences {
			if maxOccurences < v {
				maxOccurences = v
				byteMostFound = k

			} else if maxOccurences == v {
				byteMostFound = '1'
			}
		}
		res = append(res, byteMostFound)
		log.Println(occurences)
	}

	return string(res)
}

func Epsilon(inputs []string) string {
	res := []byte{}
	for i := 0; i < len(inputs[0]); i++ {
		occurences := map[byte]int{}
		for _, input := range inputs {
			occurences[input[i]] += 1
		}
		minOccurences := math.MaxInt32
		var byteMostFound byte
		for k, v := range occurences {
			if minOccurences > v {
				minOccurences = v
				byteMostFound = k

			} else if minOccurences == v {
				byteMostFound = '0'
			}
		}
		res = append(res, byteMostFound)
	}

	return string(res)
}

func OxygenGeneratorRating(inputs []string) string {
	prefix := ""
	res := ""

	for i := 0; i < len(inputs[0]); i++ {
		set := []string{}
		gamma := Gamma(inputs)
		prefix += string(gamma[i])
		for _, input := range inputs {
			if strings.HasPrefix(input, prefix) {
				set = append(set, input)
				if len(input) == len(prefix) {
					res = input
				}
			}
		}
		inputs = set
	}
	return res
}

func CO2Scrubber(inputs []string) string {
	prefix := ""
	res := ""

	for i := 0; i < len(inputs[0]); i++ {
		set := []string{}
		gamma := Epsilon(inputs)
		prefix += string(gamma[i])
		for _, input := range inputs {
			if strings.HasPrefix(input, prefix) {
				set = append(set, input)
				if len(input) == len(prefix) {
					res = input
				}
			}
		}
		inputs = set
	}
	return res
}

func convertStringBytesToInt(s string) int64 {

	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	inputs := []string{}
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	gamma := Gamma(inputs)
	epsilon := Epsilon(inputs)

	log.Println("Got Gamma: ", gamma)
	log.Println("Got Epsilon: ", epsilon)

	log.Println("Solution:", convertStringBytesToInt(gamma)*convertStringBytesToInt(epsilon))

	log.Println("------ Part 2  ---------")
	o := OxygenGeneratorRating(inputs)
	c := CO2Scrubber(inputs)
	log.Println("Got OxygenGenerator Rating: ", o)
	log.Println("Got CO2Scrubber:", c)
	log.Println("Solution:", convertStringBytesToInt(o)*convertStringBytesToInt(c))
}
