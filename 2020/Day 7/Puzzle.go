package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type bag struct {
	color   string
	parents []*bag // types of bags that directly contain this bag
}

type rule struct {
	bagColor string
	contains []string
}

func parseRule(r *rule, bagMap map[string]*bag) {
	if bagMap[r.bagColor] == nil {
		bagMap[r.bagColor] = &bag{
			color: r.bagColor,
		}
	}
	for _, color := range r.contains {
		if bagMap[color] == nil {
			bagMap[color] = &bag{
				color:   color,
				parents: []*bag{bagMap[r.bagColor]},
			}
		} else {
			bagMap[color].parents = append(bagMap[color].parents, bagMap[r.bagColor])
		}
	}
}

func parseRuleFromLine(line string) *rule {
	tokens := strings.Split(line, " ")
	r := rule{}
	for i, token := range tokens {
		switch token {
		case "bags":
			r.bagColor = strings.Join(tokens[i-2:i], " ")
		case "bags,":
			fallthrough
		case "bag,":
			fallthrough
		case "bag.":
			r.contains = append(r.contains, strings.Join(tokens[i-2:i], " "))
		case "bags.":
			words := strings.Join(tokens[i-2:i], " ")
			if words != "no other" {
				r.contains = append(r.contains, words)
			}
		}
	}
	return &r
}

func countNodes(root *bag) int {
	visitedNodes := make(map[*bag]bool)
	if root == nil || root.parents == nil {
		return 0
	}
	for _, parent := range root.parents {
		countNodesImpl(parent, visitedNodes)
	}
	return len(visitedNodes)
}

func countNodesImpl(node *bag, visitedNodes map[*bag]bool) {
	visitedNodes[node] = true
	if node.parents == nil {
		return
	}
	for _, parent := range node.parents {
		if !visitedNodes[parent] {
			countNodesImpl(parent, visitedNodes)
		}
	}
}

func grabInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1(lines []string) int {
	bagMap := make(map[string]*bag, len(lines))
	for _, line := range lines {
		rule := parseRuleFromLine(line)
		parseRule(rule, bagMap)
	}
	return countNodes(bagMap["shiny gold"])
}

func main() {
	lines := grabInput("input.txt")
	fmt.Printf("Part 1: %v\n", solvePart1(lines))
}
