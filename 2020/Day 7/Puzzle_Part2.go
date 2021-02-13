package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type bagAllowance map[string]uint

type bagRules map[string]bagAllowance

func (b bagRules) parseRuleFromLine(line string) {
	tokens := strings.Split(line, " ")
	var bag string
	for i, token := range tokens {
		switch token {
		case "bags":
			bag = strings.Join(tokens[i-2:i], " ")
			b[bag] = make(bagAllowance)
		case "bags,":
			fallthrough
		case "bag,":
			fallthrough
		case "bag.":
			b[bag][strings.Join(tokens[i-2:i], " ")] = uint(tokens[i-3][0] - '0')
		case "bags.":
			words := strings.Join(tokens[i-2:i], " ")
			if words != "no other" {
				b[bag][words] = uint(tokens[i-3][0] - '0')
			}
		}
	}
}

func (b bagRules) countBags(bagStyle string) uint {
	allowance := b[bagStyle]
	if allowance == nil {
		return 0
	}
	count := uint(0)
	for style, allowanceCount := range allowance {
		count += b.countBags(style)*allowanceCount + allowanceCount
	}
	return count
}

func parseInputLines(lines []string) (rules bagRules) {
	rules = make(bagRules, len(lines))
	for _, line := range lines {
		rules.parseRuleFromLine(line)
	}
	return rules
}

func grabInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func main() {
	r := parseInputLines(grabInput("input.txt"))
	fmt.Printf("Part 2: %v\n", r.countBags("shiny gold"))
}
