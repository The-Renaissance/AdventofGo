package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type address uint64

//// Part 1 ////
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

func newDecoder() *decoder {
	return &decoder{mem: make(map[address]uint64)}
}

//// Part 2 ////

type decoderv2 struct {
	setMask     uint64
	xbitindices []int
	mem         map[address]uint64
}

func (d *decoderv2) execute(instruction string) {
	tokens := strings.Split(instruction, " = ")
	lhs, rhs := tokens[0], tokens[1]
	switch lhs[:4] {
	case "mask":
		d.setmask(rhs)
	case "mem[":
		addr, _ := strconv.Atoi(lhs[4 : len(lhs)-1])
		value, _ := strconv.Atoi(rhs)
		d.writetomem(value, address(addr))
	}
}

func (d *decoderv2) setmask(mask string) {
	d.setMask = 0
	d.xbitindices = nil
	for i, bit := range mask {
		if bit == '0' {
			continue
		} else if bit == '1' {
			d.setMask |= 1 << (len(mask) - i - 1)
		} else if bit == 'X' {
			d.xbitindices = append(d.xbitindices, len(mask)-i-1)
		}
	}
}

func (d *decoderv2) writetomem(value int, addr address) {
	a := uint64(addr)
	a |= d.setMask
	addresses := apply(address(a), d.xbitindices)
	for _, ad := range addresses {
		d.mem[ad] = uint64(value)
	}
}

func apply(addr address, xbitindices []int) []address {
	addresses := make([]address, 0)
	apply_impl(addr, xbitindices, &addresses)
	return addresses
}

func apply_impl(addr address, xbitindices []int, addresses *[]address) {
	addr &^= 1 << xbitindices[0]
	if len(xbitindices) == 1 {
		*addresses = append(*addresses, addr)
	} else {
		apply_impl(addr, xbitindices[1:], addresses)
	}
	addr |= 1 << xbitindices[0]
	if len(xbitindices) == 1 {
		*addresses = append(*addresses, addr)
	} else {
		apply_impl(addr, xbitindices[1:], addresses)
	}
}

func (d *decoderv2) sum() uint64 {
	var sum uint64
	for _, value := range d.mem {
		sum += value
	}
	return sum
}

func newDecoderv2() *decoderv2 {
	return &decoderv2{mem: make(map[address]uint64)}
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
	f := newDecoder()
	for _, instruction := range instructions {
		f.execute(instruction)
	}
	fmt.Printf("Part 1: %v\n", f.sum())
}

func solvePart2() {
	instructions := getInput("input.txt")
	d := newDecoderv2()
	for _, instruction := range instructions {
		d.execute(instruction)
	}
	fmt.Printf("Part 2: %v\n", d.sum())
}

func main() {
	solvePart1()
	solvePart2()
}
