package main

import (
	"testing"
)

func Test_earliestBus(t *testing.T) {
	type args struct {
		ts       string
		shuttles string
	}
	tests := []struct {
		name        string
		args        args
		wantId      int
		wantMinutes int
	}{
		{
			"Example",
			args{
				ts:       "939",
				shuttles: "7,13,x,x,59,x,31,19",
			},
			59,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotMinutes := earliestBus(tt.args.ts, tt.args.shuttles)
			if gotId != tt.wantId {
				t.Errorf("earliestBus() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotMinutes != tt.wantMinutes {
				t.Errorf("earliestBus() gotMinutes = %v, want %v", gotMinutes, tt.wantMinutes)
			}
		})
	}
}

func Test_getShuttlePart2(t *testing.T) {
	type args struct {
		shuttles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example",
			args{
				"7,13,x,x,59,x,31,19",
			},
			1068781,
		},
		{
			"Example 2",
			args{
				"17,x,13,19",
			},
			3417,
		},
		{
			"Example 3",
			args{
				"67,7,59,61",
			},
			754018,
		},
		{
			"Example 4",
			args{
				"67,x,7,59,61",
			},
			779210,
		},
		{
			"Example 5",
			args{
				"67,7,x,59,61",
			},
			1261476,
		},
		{
			"Example 6",
			args{
				"1789,37,47,1889",
			},
			1202161486,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShuttlePart2(tt.args.shuttles); got != tt.want {
				t.Errorf("getShuttlePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
