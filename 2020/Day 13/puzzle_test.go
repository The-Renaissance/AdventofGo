package main

import "testing"

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
