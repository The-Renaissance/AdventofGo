package main

import (
	"reflect"
	"testing"
)

func Test_countUniqueAnswers(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"Case 1", []string{"abc"}, 3},
		{"Case 2", []string{"a", "b", "c"}, 3},
		{"Case 3", []string{"ab", "ac"}, 3},
		{"Case 4", []string{"a", "a", "a", "a"}, 1},
		{"Case 5", []string{"b"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueAnswers(tt.input); got != tt.want {
				t.Errorf("countUniqueAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countEveryonesAnswers(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"Case 1", []string{"abc"}, 3},
		{"Case 2", []string{"a", "b", "c"}, 0},
		{"Case 3", []string{"ab", "ac"}, 1},
		{"Case 4", []string{"a", "a", "a", "a"}, 1},
		{"Case 5", []string{"b"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEveryonesAnswers(tt.input); got != tt.want {
				t.Errorf("countEveryonesAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitbyEmptyLine(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			"Case 1",
			[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b"},
			[][]string{
				[]string{"abc"},
				[]string{"a", "b", "c"},
				[]string{"ab", "ac"},
				[]string{"a", "a", "a", "a"},
				[]string{"b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitbyEmptyLine(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitbyEmptyLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"Example case",
			[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b"},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.input); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"Example case",
			[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.input); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
