package main

import (
	"reflect"
	"testing"
)

func TestFloorState_AdvanceState(t *testing.T) {
	tests := []struct {
		name         string
		s            *FloorState
		quota        int
		adjacentFunc AdjacentFunc
		next         SeatMap
	}{
		{
			"Case 1",
			&FloorState{
				current: SeatMap{
					[]rune("L.LL.LL.LL"),
					[]rune("LLLLLLL.LL"),
					[]rune("L.L.L..L.."),
					[]rune("LLLL.LL.LL"),
					[]rune("L.LL.LL.LL"),
					[]rune("L.LLLLL.LL"),
					[]rune("..L.L....."),
					[]rune("LLLLLLLLLL"),
					[]rune("L.LLLLLL.L"),
					[]rune("L.LLLLL.LL"),
				},
			},
			4,
			GetImmediateAdjacents,
			SeatMap{
				[]rune("#.##.##.##"),
				[]rune("#######.##"),
				[]rune("#.#.#..#.."),
				[]rune("####.##.##"),
				[]rune("#.##.##.##"),
				[]rune("#.#####.##"),
				[]rune("..#.#....."),
				[]rune("##########"),
				[]rune("#.######.#"),
				[]rune("#.#####.##"),
			},
		},
		{
			"Case 2",
			&FloorState{
				current: SeatMap{
					[]rune("#.#L.L#.##"),
					[]rune("#LLL#LL.L#"),
					[]rune("L.L.L..#.."),
					[]rune("#LLL.##.L#"),
					[]rune("#.LL.LL.LL"),
					[]rune("#.LL#L#.##"),
					[]rune("..L.L....."),
					[]rune("#L#LLLL#L#"),
					[]rune("#.LLLLLL.L"),
					[]rune("#.#L#L#.##"),
				},
			},
			4,
			GetImmediateAdjacents,
			SeatMap{
				[]rune("#.#L.L#.##"),
				[]rune("#LLL#LL.L#"),
				[]rune("L.#.L..#.."),
				[]rune("#L##.##.L#"),
				[]rune("#.#L.LL.LL"),
				[]rune("#.#L#L#.##"),
				[]rune("..L.L....."),
				[]rune("#L#L##L#L#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.#L#L#.##"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing state: \n%s", tt.s.current.PPrint())
			tt.s.AdvanceState(tt.quota, tt.adjacentFunc)
			if !reflect.DeepEqual(tt.s.current, tt.next) {
				t.Errorf("Want: \n%s\n\n, got \n\n%s", tt.next.PPrint(), tt.s.current.PPrint())
			}
		})
	}
}

func TestFloorState_RunandCount(t *testing.T) {
	tests := []struct {
		name         string
		s            *FloorState
		quota        int
		adjacentFunc AdjacentFunc
		want         int
	}{
		{
			"Example",
			FromInitialState(
				[]string{
					"L.LL.LL.LL",
					"LLLLLLL.LL",
					"L.L.L..L..",
					"LLLL.LL.LL",
					"L.LL.LL.LL",
					"L.LLLLL.LL",
					"..L.L.....",
					"LLLLLLLLLL",
					"L.LLLLLL.L",
					"L.LLLLL.LL",
				},
			),
			4,
			GetImmediateAdjacents,
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RunandCount(tt.quota, tt.adjacentFunc); got != tt.want {
				t.Errorf("FloorState.RunandCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
