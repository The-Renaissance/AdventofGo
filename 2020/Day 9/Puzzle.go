package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func solvePart1(xmas []int, preambleLen int) int {
	result := 0
	preamblelist := make([]int, preambleLen)
	for i, num := range xmas[preambleLen:] {
		copy(preamblelist, xmas[i:i+preambleLen])
		if checkNumber(preamblelist, num) != true {
			result = num
			break
		}
	}
	return result
}

func checkNumber(preambles []int, number int) (found bool) {
	sort.Ints(preambles)
	leftIndex := 0
	rightIndex := len(preambles) - 1
	for leftIndex < rightIndex {
		if a, b := preambles[leftIndex], preambles[rightIndex]; a+b < number {
			leftIndex++
		} else if a+b > number {
			rightIndex--
		} else {
			found = true
			break
		}
	}
	return
}

func parseList(filename string) (xmas []int) {
	lines := getInput(filename)
	for i, line := range lines {
		if num, err := strconv.Atoi(line); err != nil {
			log.Panicf("parsing %v at line %v failed\n", line, i+1)
		} else {
			xmas = append(xmas, num)
		}
	}
	return
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func main() {
	xmas := parseList("input.txt")
	fmt.Printf("Part 1: %v\n", solvePart1(xmas, 25))
}
