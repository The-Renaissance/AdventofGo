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

func solvePart2(inputMap []string) int {
	multResult := 1
	type slope struct {
		right, down int
	}
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, slope := range slopes {
		multResult *= solvePart1(inputMap, slope.right, slope.down)
	}
	return multResult
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
	fmt.Printf("Part 2: %v\n", solvePart2(inputMap))
}
