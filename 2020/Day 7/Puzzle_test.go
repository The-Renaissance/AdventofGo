package main

import (
	"reflect"
	"testing"
)

func Test_parseRule(t *testing.T) {
	bagMap := map[string]*bag{}
	rules := []rule{
		rule{
			bagColor: "light red",
			contains: []string{"bright white", "muted yellow"},
		},
		rule{
			bagColor: "dark orange",
			contains: []string{"bright white", "muted yellow"},
		},
		rule{
			bagColor: "bright white",
			contains: []string{"shiny gold"},
		},
		rule{
			bagColor: "muted yellow",
			contains: []string{"shiny gold", "faded blue"},
		},
		rule{
			bagColor: "shiny gold",
			contains: []string{"dark olive", "vibrant plum"},
		},
		rule{
			bagColor: "dark olive",
			contains: []string{"faded blue", "dotted black"},
		},
		rule{
			bagColor: "vibrant plum",
			contains: []string{"faded blue", "dotted black"},
		},
		rule{
			bagColor: "faded blue",
			contains: nil,
		},
		rule{
			bagColor: "dotted black",
			contains: nil,
		},
	}
	for _, r := range rules {
		parseRule(&r, bagMap)
	}
	expected := map[string][]string{
		"faded blue":   []string{"muted yellow", "dark olive", "vibrant plum"},
		"shiny gold":   []string{"muted yellow", "bright white"},
		"muted yellow": []string{"light red", "dark orange"},
		"bright white": []string{"light red", "dark orange"},
		"light red":    nil,
		"dark orange":  nil,
		"dark olive":   []string{"shiny gold"},
		"vibrant plum": []string{"shiny gold"},
		"dotted black": []string{"dark olive", "vibrant plum"},
	}
	for bag, expectedParents := range expected {
		if bagMap[bag] == nil {
			t.Errorf("Rule for bag color %v is not parsed!", bag)
			continue
		}
		if expectedParents == nil && bagMap[bag].parents != nil {
			t.Errorf("%v should not have parent but has %v", bag, bagMap[bag].parents)
			continue
		}
		if len(expectedParents) != len(bagMap[bag].parents) {
			t.Errorf("Expect %v parents, got %v parents for %v", len(expectedParents), len(bagMap[bag].parents), bag)
			continue
		}
		for _, parent := range expectedParents {
			found := false
			for _, parentinMap := range bagMap[bag].parents {
				if parentinMap.color == parent {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected to find color '%v' as a parent, got no such parent", parent)
			}
		}
	}
}

func Test_parseRuleFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *rule
	}{
		{"Case 1", args{"light red bags contain 1 bright white bag, 2 muted yellow bags."}, &rule{
			bagColor: "light red",
			contains: []string{"bright white", "muted yellow"},
		}},
		{"Case 2", args{"dark orange bags contain 3 bright white bags, 4 muted yellow bags."}, &rule{
			bagColor: "dark orange",
			contains: []string{"bright white", "muted yellow"},
		}},
		{"Case 3", args{"bright white bags contain 1 shiny gold bag."}, &rule{
			bagColor: "bright white",
			contains: []string{"shiny gold"},
		}},
		{"No other bags", args{"dotted black bags contain no other bags."}, &rule{
			bagColor: "dotted black",
			contains: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRuleFromLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRuleFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodes(t *testing.T) {
	// node structure

	lightRed := &bag{
		color:   "light red",
		parents: nil,
	}

	darkOrange := &bag{
		color:   "dark orange",
		parents: nil,
	}

	mutedYellow := &bag{
		color: "muted yellow",
		parents: []*bag{
			lightRed,
			darkOrange,
		},
	}

	brightWhite := &bag{
		color: "bright white",
		parents: []*bag{
			lightRed,
			darkOrange,
		},
	}

	shinyGold := &bag{
		color: "shiny gold",
		parents: []*bag{
			mutedYellow,
			brightWhite,
		},
	}

	want := 4
	if got := countNodes(shinyGold); got != want {
		t.Errorf("countNodes() = %v, want %v", got, want)
	}
}
