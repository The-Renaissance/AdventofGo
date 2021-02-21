package main

import "testing"

func Test_getNumberInSequence(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		sequencenr int
		want       int
	}{
		{
			"Permutation 1",
			[]int{1, 3, 2},
			2020,
			1,
		},
		{
			"Permutation 2",
			[]int{2, 1, 3},
			2020,
			10,
		},
		{
			"Permutation 3",
			[]int{1, 2, 3},
			2020,
			27,
		},
		{
			"Permutation 4",
			[]int{2, 3, 1},
			2020,
			78,
		},
		{
			"Permutation 5",
			[]int{3, 2, 1},
			2020,
			438,
		},
		{
			"Permutation 6",
			[]int{3, 1, 2},
			2020,
			1836,
		},
		{
			"Permutation 7",
			[]int{0, 3, 6},
			2020,
			436,
		},
		{
			"Part 1",
			[]int{0, 20, 7, 16, 1, 18, 15},
			2020,
			1025,
		},
		{
			"Part 2 Permutation 1",
			[]int{1, 3, 2},
			30000000,
			2578,
		},
		{
			"Part 2 Permutation 2",
			[]int{2, 1, 3},
			30000000,
			3544142,
		},
		{
			"Part 2 Permutation 3",
			[]int{1, 2, 3},
			30000000,
			261214,
		},
		{
			"Part 2 Permutation 4",
			[]int{0, 3, 6},
			30000000,
			175594,
		},
		{
			"Part 2 Permutation 5",
			[]int{2, 3, 1},
			30000000,
			6895259,
		},
		{
			"Part 2 Permutation 6",
			[]int{3, 2, 1},
			30000000,
			18,
		},
		{
			"Part 2 Permutation 7",
			[]int{3, 1, 2},
			30000000,
			362,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberInSequence(tt.input, tt.sequencenr); got != tt.want {
				t.Errorf("getNumberInSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
