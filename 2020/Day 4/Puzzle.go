package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
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

func solvePart2(input string) int {
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
				pair := strings.Split(scanner.Text(), ":")
				field, value := pair[0], pair[1]
				fieldValidRecord[field] = validateValue(field, value)
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

func validateValue(field string, value string) bool {
	switch field {
	case "byr":
		parsedbyr, err := strconv.Atoi(value)
		return err == nil && 1920 <= parsedbyr && parsedbyr <= 2002
	case "iyr":
		parsediyr, err := strconv.Atoi(value)
		return err == nil && 2010 <= parsediyr && parsediyr <= 2020
	case "eyr":
		parsedeyr, err := strconv.Atoi(value)
		return err == nil && 2020 <= parsedeyr && parsedeyr <= 2030
	case "hgt":
		n := len(value)
		if n < 4 {
			return false
		}
		unit := value[n-2:]
		parsedhgt, err := strconv.Atoi(value[:n-2])
		if err != nil {
			return false
		}
		if unit == "cm" {
			return 150 <= parsedhgt && parsedhgt <= 193
		} else if unit == "in" {
			return 59 <= parsedhgt && parsedhgt <= 76
		} else {
			return false
		}
	case "hcl":
		if n := len(value); n != 7 || value[0] != '#' {
			return false
		}
		valid := true
		for i := 1; i <= 6; i++ {
			if !('0' <= value[i] && value[i] <= '9') &&
				!('a' <= value[i] && value[i] <= 'f') {
				valid = false
				break
			}
		}
		return valid
	case "ecl":
		colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range colors {
			if value == color {
				return true
			}
		}
		return false
	case "pid":
		return len(value) == 9
	case "cid":
		return true
	}
	return false
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
	examplePart1 := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm

	iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	hcl:#cfa07d byr:1929

	hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm

	hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`
	examplePart2 := `eyr:1972 cid:100
	hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

	iyr:2019
	hcl:#602927 eyr:1967 hgt:170cm
	ecl:grn pid:012533040 byr:1946

	hcl:dab227 iyr:2012
	ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

	hgt:59cm ecl:zzz
	eyr:2038 hcl:74454a iyr:2023
	pid:3556412378 byr:2007

	pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
	hcl:#623a2f

	eyr:2029 ecl:blu cid:129 byr:1989
	iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

	hcl:#888785
	hgt:164cm byr:2001 iyr:2015 cid:88
	pid:545766238 ecl:hzl
	eyr:2022

	iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	resultPart1 := solvePart1(examplePart1)
	assert(resultPart1 == 2, fmt.Sprintf("Part 1 example input test failed, got: %v", resultPart1))
	resultPart2 := solvePart2(examplePart2)
	assert(resultPart2 == 4, fmt.Sprintf("Part 2 example input test failed, got: %v", resultPart2))
	fmt.Printf("Part 1: %v\n", solvePart1(passports))
	fmt.Printf("Part 2: %v\n", solvePart2(passports))
}
