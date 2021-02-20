package main

import "testing"

func TestPart1(t *testing.T) {
	f := newDecoder()
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

func TestPart2(t *testing.T) {
	d := newDecoderv2()
	instructions := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	want := uint64(208)
	for _, ins := range instructions {
		d.execute(ins)
	}
	if got := d.sum(); got != want {
		t.Errorf("Got sum %v, want %v", got, want)
	}
}
