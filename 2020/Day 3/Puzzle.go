package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func solvePart1(inputMap []string, right, down int) int {
	count := 0
	width := len(inputMap[0])
	height := len(inputMap)
	for x, y := right, down; y < height; x, y = (x+right)%width, y+down {
		if inputMap[y][x] == '#' {
			count++
		}
	}
	return count
}

func readMap(filename string) []string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read map from %v failed\n", filename)
	}
	return strings.Split(strings.Trim(string(f), "\n"), "\n")
}

func main() {
	inputMap := readMap("input.txt")
	/*
		exampleMap := []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}
	*/
	rightStep := 3
	downStep := 1
	fmt.Printf("Part 1: %v\n", solvePart1(inputMap, rightStep, downStep))
}
