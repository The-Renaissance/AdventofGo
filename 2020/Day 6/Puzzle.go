package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countUniqueAnswers(answers []string) int {
	unique := make(map[rune]bool, 26)
	for _, line := range answers {
		for _, c := range line {
			unique[c] = true
		}
	}
	return len(unique)
}

func splitbyEmptyLine(input []string) [][]string {
	groupedLines := [][]string{}
	currentGroup := []string{}
	for _, line := range input {
		if line == "" {
			groupedLines = append(groupedLines, currentGroup)
			currentGroup = []string{}
			continue
		}
		currentGroup = append(currentGroup, line)
	}
	if len(currentGroup) != 0 {
		groupedLines = append(groupedLines, currentGroup)
	}
	return groupedLines
}

func grabInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1(input []string) int {
	answers := splitbyEmptyLine(input)
	total := 0
	for _, answer := range answers {
		total += countUniqueAnswers(answer)
	}
	return total
}

func main() {
	input := grabInput("input.txt")
	fmt.Println("Part 1:", solvePart1(input))
}
