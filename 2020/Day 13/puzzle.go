package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func earliestBus(ts, shuttles string) (id int, minutes int) {
	timestamp, _ := strconv.Atoi(ts)
	for _, busline := range strings.Split(shuttles, ",") {
		if n, err := strconv.Atoi(busline); err == nil {
			if timestamp%n == 0 {
				return n, 0
			}
			if minutes == 0 || n-timestamp%n < minutes {
				minutes = n - timestamp%n
				id = n
			}
		}
	}
	return id, minutes
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1() {
	input := getInput("input.txt")
	id, minutes := earliestBus(input[0], input[1])
	fmt.Printf("Part 1: %v\n", id*minutes)
}

func main() {
	solvePart1()
}
