package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) (count int) {
	policyPattern := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	for _, policy := range input {
		if checkPolicy(policy, policyPattern) {
			count++
		}
	}
	return count
}

func checkPolicy(policy string, pattern *regexp.Regexp) bool {
	extracted := pattern.FindStringSubmatch(policy)
	min, err := strconv.Atoi(extracted[1])
	if err != nil {
		log.Fatalf("checkPolicy: Parsing %s out of %s failed\n", extracted[1], policy)
	}
	max, err := strconv.Atoi(extracted[2])
	if err != nil {
		log.Fatalf("checkPolicy: Parsing %s out of %s failed\n", extracted[2], policy)
	}
	character, password := extracted[3], extracted[4]
	count := strings.Count(password, character)
	return min <= count && count <= max
}

func main() {
	blob, _ := ioutil.ReadFile("input.txt")
	passwordList := strings.Split(strings.Trim(string(blob), "\n"), "\n")
	// passwordList := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	fmt.Printf("Part 1: %v\n", part1(passwordList))
}
