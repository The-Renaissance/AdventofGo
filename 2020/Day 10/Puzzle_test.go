package main

import "testing"

func TestFindNumberofDifferences(t *testing.T) {
	type result struct {
		onejolts, threejolts int
	}
	cases := []struct {
		name     string
		adapters []int
		want     result
	}{
		{
			"Example",
			[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			result{7, 5},
		},
		{
			"Larger Example",
			[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			result{22, 10},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if onejolts, threejolts := FindNumberofDifferences(tt.adapters); onejolts != tt.want.onejolts || threejolts != tt.want.threejolts {
				t.Errorf("FindNumberofDifferences() = %v,%v want %v,%v", onejolts, threejolts, tt.want.onejolts, tt.want.threejolts)
			}
		})
	}
}

func TestFindArrangements(t *testing.T) {
	cases := []struct {
		name     string
		adapters []int
		want     int
	}{
		{
			"Example",
			[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			8,
		},
		{
			"Larger Example",
			[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			19208,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindArrangements(tt.adapters); got != tt.want {
				t.Errorf("FindArrangements() = %v want %v", got, tt.want)
			}
		})
	}
}
