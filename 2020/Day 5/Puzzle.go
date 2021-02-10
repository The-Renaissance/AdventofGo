package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func getSeatID(boardingPass string) int {
	row, column := grabRowAndColumn(boardingPass)
	return 8*row + column
}

func grabRowAndColumn(boardingPass string) (row int, column int) {
	const numRows = 128
	const numColumns = 8
	rowL, rowR := 0, numRows-1
	columnL, columnR := 0, numColumns-1
	i := 0
	for rowL != rowR {
		if boardingPass[i] == 'F' {
			rowR = (rowL + rowR) / 2
		} else if boardingPass[i] == 'B' {
			rowL = (rowL+rowR)/2 + 1
		} else {
			log.Fatalln("Boarding pass at", i, "position is", boardingPass[i])
		}
		i++
	}
	row = rowL
	for columnL != columnR {
		if boardingPass[i] == 'L' {
			columnR = (columnL + columnR) / 2
		} else if boardingPass[i] == 'R' {
			columnL = (columnL+columnR)/2 + 1
		} else {
			log.Fatalln("Boarding pass at", i, "position is", boardingPass[i])
		}
		i++
	}
	column = columnL
	if i != len(boardingPass) {
		log.Panicln("Length of boarding pass is too long")
	}
	return row, column
}

func grabBoardingPasses(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed")
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func getSortedSeatlist(boardingPasses []string) []int {
	seatList := make([]int, 0, len(boardingPasses))
	for _, bp := range boardingPasses {
		seatList = append(seatList, getSeatID(bp))
	}
	sort.Ints(seatList)
	return seatList
}

func solvePart1(boardingPasses []string) int {
	seatList := getSortedSeatlist(boardingPasses)
	return seatList[len(seatList)-1]
}

func solvePart2(boardingPasses []string) int {
	seatList := getSortedSeatlist(boardingPasses)
	var vacantSeat int
	for i, seat := range seatList {
		if seatList[i+1]-seat > 1 {
			vacantSeat = seat + 1
			break
		}
	}
	return vacantSeat
}

func main() {
	boardingPasses := grabBoardingPasses("input.txt")
	fmt.Printf("Part 1: %v\n", solvePart1(boardingPasses))
	fmt.Printf("Part 2: %v\n", solvePart2(boardingPasses))
}
