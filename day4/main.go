package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Board [][]int

func (b Board) getColumn(x int) []int {
	column := []int{}
	for _, row := range b {
		column = append(column, row[x])
	}
	return column
}

func (b Board) getRow(x int) []int {
	return b[x]
}

func (b Board) getWinningCombinations() [][]int {
	res := [][]int{}
	for i := 0; i < len(b[0]); i++ {
		res = append(res, b.getRow(i))
		res = append(res, b.getColumn(i))
	}

	return res
}

func (b Board) isWinning(input []int) (bool, int) {
	draw := input[0:len(b[0])]
	for i := len(b[0]); i < len(input); i++ {
		draw = append(draw, input[i])
		for _, winningCombination := range b.getWinningCombinations() {
			if Contains(winningCombination, draw) {
				return true, i
			}
		}

	}
	return false, 0
}

func (b Board) Flatten() []int {
	res := []int{}
	for _, r := range b {
		for _, v := range r {
			res = append(res, v)
		}
	}
	return res
}

func Contains(smallSlice []int, bigSlice []int) bool {
	copySmall := make([]int, len(smallSlice))
	copyBig := make([]int, len(bigSlice))
	copy(copySmall, smallSlice)
	copy(copyBig, bigSlice)

	sortedSmallSlice := sort.IntSlice(copySmall)
	sortedBigSlice := sort.IntSlice(copyBig)

	sortedSmallSlice.Sort()
	sortedBigSlice.Sort()

	for len(sortedSmallSlice) > 0 {
		switch {
		case len(sortedBigSlice) == 0:
			return false
		case sortedSmallSlice[0] == sortedBigSlice[0]:
			sortedSmallSlice = sortedSmallSlice[1:]
			sortedBigSlice = sortedBigSlice[1:]
		case sortedSmallSlice[0] < sortedBigSlice[0]:
			return false
		case sortedSmallSlice[0] > sortedBigSlice[0]:
			sortedBigSlice = sortedBigSlice[1:]
		}

	}
	return true
}

func RemoveOccurences(source []int, occ []int) []int {
	copyOcc := make([]int, len(occ))
	copy(copyOcc, occ)
	copySource := make([]int, len(source))
	copy(copySource, source)

	sortedSmallSlice := sort.IntSlice(copyOcc)
	sortedBigSlice := sort.IntSlice(copySource)

	sortedSmallSlice.Sort()
	sortedBigSlice.Sort()

	res := []int{}

	for len(sortedSmallSlice) > 0 && len(sortedBigSlice) > 0 {
		switch {
		case sortedSmallSlice[0] == sortedBigSlice[0]:
			sortedSmallSlice = sortedSmallSlice[1:]
			sortedBigSlice = sortedBigSlice[1:]

		case sortedSmallSlice[0] < sortedBigSlice[0]:
			sortedSmallSlice = sortedSmallSlice[1:]

		case sortedSmallSlice[0] > sortedBigSlice[0]:
			res = append(res, sortedBigSlice[0])
			sortedBigSlice = sortedBigSlice[1:]

		}

	}

	if len(sortedBigSlice) == 0 {
		return res
	}
	if len(sortedSmallSlice) == 0 {
		return sortedBigSlice
	}

	return res
}

func main() {
	start := time.Now()
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	drawStr := scanner.Text()
	drawStrArr := strings.Split(drawStr, ",")
	draw := []int{}

	log.Println("Formatting drawing..... Content : ", drawStrArr)

	for _, k := range drawStrArr {
		v, err := strconv.Atoi(k)
		if err != nil {
			log.Fatal(err)
		}
		draw = append(draw, v)
	}

	boards := []Board{}

	for scanner.Scan() {
		// Beggining of a new board
		if scanner.Text() == "" {
			board := Board{}
			for i := 0; i < 5; i++ {
				scanner.Scan()
				rawRow := scanner.Text()
				splittedRawRow := strings.Split(rawRow, " ")
				row := []int{}
				for _, k := range splittedRawRow {
					if k != "" {
						v, err := strconv.Atoi(strings.TrimSpace(k))
						if err != nil {
							log.Fatal(err)
						}
						row = append(row, v)
					}
				}
				board = append(board, row)
			}
			boards = append(boards, board)
		}
	}

	minTurn := math.MaxInt32
	winningBoard := Board{}
	losingBoard := Board{}
	maxTurn := 0

	for _, b := range boards {
		win, turns := b.isWinning(draw)
		if win {
			if turns < minTurn {
				minTurn = turns
				winningBoard = b
			}
			if maxTurn < turns {
				maxTurn = turns
				losingBoard = b
			}
		}
	}

	log.Println("The winning board won in", minTurn+1, " turns, winning number", draw[minTurn], " and it's this board: ", winningBoard)
	log.Println("All those numbers were out:", draw[0:minTurn+1])
	numbersNotFound := RemoveOccurences(winningBoard.Flatten(), draw[0:minTurn+1])

	res := 0

	for _, o := range numbersNotFound {
		res += o
	}

	log.Println("Part 1 Winning score: ", res*draw[minTurn])

	log.Println("Part 2 ----------")
	log.Println("Loosing board:", losingBoard, "Lost in", maxTurn+1, "winningNumber", draw[maxTurn])
	log.Println()
	loosingNumbersNotFound := RemoveOccurences(losingBoard.Flatten(), draw[0:maxTurn+1])
	res = 0

	for _, o := range loosingNumbersNotFound {
		res += o
	}

	log.Println(" Loosing score: ", res*draw[maxTurn])
	elapsed := time.Since(start)
	log.Printf("Elapsed %s\n", elapsed)

}
