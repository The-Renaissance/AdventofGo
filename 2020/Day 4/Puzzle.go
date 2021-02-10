package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func solvePart1(input string) int {
	validFields := []string{"byr", "eyr", "iyr", "hgt", "hcl", "ecl", "pid"}
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	validPassportCount := 0
	fieldValidRecord := make(map[string]bool)
	for _, line := range lines {
		if line == "" {
			if validatePassport(validFields, fieldValidRecord) {
				validPassportCount++
			}
			fieldValidRecord = make(map[string]bool)
		} else {
			buf := bytes.NewBufferString(line)
			scanner := bufio.NewScanner(buf)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				pair := scanner.Text()
				fieldValidRecord[strings.Split(pair, ":")[0]] = true
			}
		}
	}
	if validatePassport(validFields, fieldValidRecord) {
		validPassportCount++
	}
	return validPassportCount
}

func validatePassport(validFields []string, fieldRecord map[string]bool) bool {
	valid := true
	for _, field := range validFields {
		if exist := fieldRecord[field]; !exist {
			valid = false
			break
		}
	}
	return valid
}

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)
	assert(err == nil, "Error reading file")
	return string(content)
}

func assert(cond bool, errstr string) {
	if !cond {
		panic(errstr)
	}
}

func main() {
	passports := readInput("input.txt")
	example := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm

	iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	hcl:#cfa07d byr:1929

	hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm

	hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`
	testOut := solvePart1(example)
	assert(testOut == 2, fmt.Sprintf("Example input test failed, got: %v", testOut))
	fmt.Printf("Part 1: %v", solvePart1(passports))
}
