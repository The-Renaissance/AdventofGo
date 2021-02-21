package main

import "fmt"

func getNumberInSequence(input []int, sequencenr int) int {
	var prevNumber, currNumber int
	numberTable := make(map[int]int)
	for i := 0; i < sequencenr; i++ {
		if i < len(input) {
			currNumber = input[i]
			numberTable[currNumber] = i
		} else if i == len(input) {
			currNumber = 0
		} else {
			prevNumber = currNumber
			if prevIndex, ok := numberTable[prevNumber]; ok {
				currNumber = i - prevIndex - 1
			} else {
				currNumber = 0
			}
			numberTable[prevNumber] = i - 1
		}
	}
	return currNumber
}

func main() {
	input := []int{0, 20, 7, 16, 1, 18, 15}
	fmt.Printf("Part 1: %v\n", getNumberInSequence(input, 2020))
}
