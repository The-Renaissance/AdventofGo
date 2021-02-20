package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type address uint64

type decoder struct {
	setMask, clearMask uint64
	mem                map[address]uint64
}

func (f *decoder) setmask(mask string) {
	f.setMask = 0
	f.clearMask = 0
	for i, bit := range mask {
		if bit == 'X' {
			continue
		} else if bit == '0' {
			f.clearMask |= 1 << (len(mask) - i - 1)
		} else if bit == '1' {
			f.setMask |= 1 << (len(mask) - i - 1)
		}
	}
}

func (f *decoder) writetomem(value int, addr address) {
	v := uint64(value)
	v |= f.setMask
	v &^= f.clearMask
	if f.mem == nil {
		f.mem = make(map[address]uint64)
	}
	f.mem[addr] = v
}

func (f *decoder) execute(instruction string) {
	tokens := strings.Split(instruction, " = ")
	lhs, rhs := tokens[0], tokens[1]
	switch lhs[:4] {
	case "mask":
		f.setmask(rhs)
	case "mem[":
		addr, _ := strconv.Atoi(lhs[4 : len(lhs)-1])
		value, _ := strconv.Atoi(rhs)
		f.writetomem(value, address(addr))
	}
}

func (f *decoder) sum() uint64 {
	var sum uint64
	for _, value := range f.mem {
		sum += value
	}
	return sum
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1() {
	instructions := getInput("input.txt")
	var f decoder
	for _, instruction := range instructions {
		f.execute(instruction)
	}
	fmt.Printf("Part 1: %v\n", f.sum())
}

func main() {
	solvePart1()
}
