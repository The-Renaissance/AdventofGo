package main

import "testing"

func TestVM(t *testing.T) {
	cases := []struct {
		name     string
		assembly []string
		want     int
	}{
		{
			"Example case",
			[]string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			5,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v := new(VM)
			v.LoadAssembly(tt.assembly)
			if got := v.Run(); got != tt.want {
				t.Errorf("VM returned %v, want %v\n", got, tt.want)
			}
		})
	}
}
