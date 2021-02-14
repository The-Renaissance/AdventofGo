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
		{
			"Patched Program",
			[]string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"nop -4",
				"acc +6",
			},
			8,
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

func Test_diagnose(t *testing.T) {
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
			7,
		},
		{
			"Change a nop",
			[]string{
				"nop +0",
				"nop +5",
				"acc +99",
				"jmp -1",
				"acc -42",
				"jmp -3",
				"acc +6",
			},
			1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v := new(VM)
			v.LoadAssembly(tt.assembly)
			if got, err := v.DiagnoseAndFix(); err != nil {
				t.Errorf("DiagnoseAndFix() returned an error\n")
			} else if got != tt.want {
				t.Errorf("DiagnoseAndFix() returned %v, want %v\n", got, tt.want)
			}
		})
	}
}
