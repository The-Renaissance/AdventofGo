package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// FindNumberofDifferences returns the number of 1-jolt and 3-jolt differences
func FindNumberofDifferences(adapters []int) (onejolts, threejolts int) {
	sort.Ints(adapters)
	for i, rating := range adapters {
		var diff int
		if i == 0 {
			diff = rating - 0
		} else {
			diff = rating - adapters[i-1]
		}
		if diff == 1 {
			onejolts++
		} else if diff == 3 {
			threejolts++
		}
	}
	threejolts++ // difference between largest adapter and target is always 3
	return
}

func parseList(filename string) []int {
	lines := getInput(filename)
	list := make([]int, 0, len(lines))
	for i, line := range lines {
		if num, err := strconv.Atoi(line); err != nil {
			log.Panicf("parsing %v at line %v failed\n", line, i+1)
		} else {
			list = append(list, num)
		}
	}
	return list
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1() {
	adapters := parseList("input.txt")
	onejolts, threejolts := FindNumberofDifferences(adapters)
	fmt.Printf("Part 1: %v\n", onejolts*threejolts)
}

func main() {
	solvePart1()
}
