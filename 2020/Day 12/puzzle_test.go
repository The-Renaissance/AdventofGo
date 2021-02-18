package main

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		name         string
		instructions []string
		wantdistance int
	}{
		{
			"Example",
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			25,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShip()
			for _, instruction := range tt.instructions {
				s.ExecuteInstruction(instruction)
			}
			if got := s.GetDistance(); got != tt.wantdistance {
				t.Errorf("Want Manhattan Distance %v, got %v", tt.wantdistance, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		name         string
		instructions []string
		wantdistance int
	}{
		{
			"Example",
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			286,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShip()
			for _, instruction := range tt.instructions {
				s.ExecuteInstructionPart2(instruction)
			}
			if got := s.GetDistance(); got != tt.wantdistance {
				t.Errorf("Want Manhattan Distance %v, got %v", tt.wantdistance, got)
			}
		})
	}
}
