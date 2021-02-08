package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func solvePart(input []string, checkPasswd func(int, int, string, string) bool) (count int) {
	policyPattern := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	for _, policy := range input {
		if checkPolicy(policy, policyPattern, checkPasswd) {
			count++
		}
	}
	return count
}

func checkPolicy(policy string, pattern *regexp.Regexp, checkPasswd func(int, int, string, string) bool) bool {
	extracted := pattern.FindStringSubmatch(policy)
	first, err := strconv.Atoi(extracted[1])
	if err != nil {
		log.Fatalf("checkPolicy: Parsing %s out of %s failed\n", extracted[1], policy)
	}
	second, err := strconv.Atoi(extracted[2])
	if err != nil {
		log.Fatalf("checkPolicy: Parsing %s out of %s failed\n", extracted[2], policy)
	}
	character, password := extracted[3], extracted[4]
	return checkPasswd(first, second, character, password)
}

func checkPasswordPart1(min, max int, character, password string) bool {
	count := strings.Count(password, character)
	return min <= count && count <= max
}

func part1(input []string) int {
	return solvePart(input, checkPasswordPart1)
}

func main() {
	blob, _ := ioutil.ReadFile("input.txt")
	passwordList := strings.Split(strings.Trim(string(blob), "\n"), "\n")
	// passwordList := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	fmt.Printf("Part 1: %v\n", part1(passwordList))
}
