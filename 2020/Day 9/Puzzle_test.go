package main

import "testing"

func Test_solvePart1(t *testing.T) {
	type args struct {
		xmas        []int
		preambleLen int
	}
	cases := []struct {
		name  string
		input args
		want  int
	}{
		{
			"Example",
			args{
				[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
				5,
			},
			127,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.input.xmas, tt.input.preambleLen); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v\n", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		xmas []int
		want int
	}
	cases := []struct {
		name              string
		input             args
		smallest, largest int
	}{
		{
			"Example",
			args{
				[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
				127,
			},
			15, 47,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if smallest, largest := solvePart2(tt.input.xmas, tt.input.want); smallest != tt.smallest || largest != tt.largest {
				t.Errorf("solvePart2() = (%v,%v), want (%v,%v)\n", smallest, largest, tt.smallest, tt.largest)
			}
		})
	}
}
