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
				t.Errorf("solvePart1 = %v, want %v\n", got, tt.want)
			}
		})
	}
}
