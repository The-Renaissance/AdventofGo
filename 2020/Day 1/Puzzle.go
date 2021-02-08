package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sortList(input []int) {
	sort.Ints(input)
}

func solvePuzzle1(sortedList []int) int {
	topIndex := 0
	bottomIndex := len(sortedList) - 1
	result := 0
	for topIndex < bottomIndex {
		if a, b := sortedList[topIndex], sortedList[bottomIndex]; a+b < 2020 {
			topIndex++
		} else if a+b > 2020 {
			bottomIndex--
		} else {
			result = a * b
			break
		}
	}
	return result
}

func solvePuzzle2(sortedList []int) int {
	topIndex := 0
	outerIndex := len(sortedList) - 1
	innerIndex := outerIndex - 1
	result := 0
	for topIndex < innerIndex && innerIndex < outerIndex {
		if a, b, c := sortedList[topIndex], sortedList[innerIndex], sortedList[outerIndex]; a+b+c < 2020 {
			if topIndex < innerIndex-1 {
				topIndex++
			} else {
				outerIndex--
				innerIndex = outerIndex - 1
				topIndex = 0
			}
		} else if a+b+c > 2020 {
			if innerIndex > topIndex+1 {
				innerIndex--
			} else {
				outerIndex--
				innerIndex = outerIndex - 1
				topIndex = 0
			}
		} else {
			result = a * b * c
			break
		}
	}
	return result
}

func grabInput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open file %v\n", filename)
		os.Exit(1)
	}
	expenses := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			num, _ := strconv.Atoi(text)
			expenses = append(expenses, num)
		}
	}
	return expenses
}

func main() {
	expenses := grabInput("input.txt")
	// expenses := []int{1721, 979, 366, 299, 675, 1456}
	// expenses := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 21, 800, 811, 822, 833, 844, 1009, 1010, 1209}
	sortList(expenses)
	fmt.Printf("Puzzle 1 answer: %v\n", solvePuzzle1(expenses))
	fmt.Printf("Puzzle 2 answer: %v\n", solvePuzzle2(expenses))
}
