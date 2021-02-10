package main

import (
	"testing"
)

func Test_grabRowAndColumn(t *testing.T) {
	type args struct {
		boardingPass string
	}
	tests := []struct {
		name       string
		args       args
		wantRow    int
		wantColumn int
	}{
		{"Case 1", args{"BFFFBBFRRR"}, 70, 7},
		{"Case 2", args{"FFFBBBFRRR"}, 14, 7},
		{"Case 3", args{"BBFFBBFRLL"}, 102, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotColumn := grabRowAndColumn(tt.args.boardingPass)
			if gotRow != tt.wantRow {
				t.Errorf("grabRowAndColumn() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotColumn != tt.wantColumn {
				t.Errorf("grabRowAndColumn() gotColumn = %v, want %v", gotColumn, tt.wantColumn)
			}
		})
	}
}

func Test_getSeatID(t *testing.T) {
	type args struct {
		boardingPass string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"BFFFBBFRRR"}, 567},
		{"Test 2", args{"FFFBBBFRRR"}, 119},
		{"Test 3", args{"BBFFBBFRLL"}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeatID(tt.args.boardingPass); got != tt.want {
				t.Errorf("getSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
