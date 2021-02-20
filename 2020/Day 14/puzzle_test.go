package main

import "testing"

func TestPart1(t *testing.T) {
	var f decoder
	instructions := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	want := uint64(165)
	for _, ins := range instructions {
		f.execute(ins)
	}
	if got := f.sum(); got != want {
		t.Errorf("Got sum %v, want %v", got, want)
	}
}
