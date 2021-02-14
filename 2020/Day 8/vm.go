package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type opcode uint8

const (
	acc opcode = iota
	jmp
	nop
)

type instruction struct {
	op      opcode
	operand int
}

// VM implements a virtual machine
type VM struct {
	accumulator         int
	programCounter      int
	memory              []instruction
	visitedInstructions map[int]bool
}

func (v *VM) step() {
	instruction := v.memory[v.programCounter]
	switch instruction.op {
	case acc:
		v.accumulator += instruction.operand
	case nop:
	case jmp:
		v.programCounter += instruction.operand
		return
	}
	v.programCounter++
}

// Run runs the program stored in the VM
func (v *VM) Run() int {
	for !v.visitedInstructions[v.programCounter] {
		v.visitedInstructions[v.programCounter] = true
		v.step()
	}
	return v.accumulator
}

func opcodeFromToken(token string) (op opcode) {
	switch token {
	case "acc":
		op = acc
	case "nop":
		op = nop
	case "jmp":
		op = jmp
	default:
		log.Panicf("Got token %q, which cannot be parsed!\n", token)
	}
	return
}

// LoadAssembly loads assembly instructions into the memory and resets the VM
func (v *VM) LoadAssembly(asm []string) {
	v.memory = make([]instruction, 0, len(asm))
	v.visitedInstructions = make(map[int]bool, len(asm))
	v.accumulator = 0
	v.programCounter = 0
	for _, line := range asm {
		tokens := strings.Split(line, " ")
		op := opcodeFromToken(tokens[0])
		operand, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Panicln(err)
		}
		v.memory = append(v.memory, instruction{op: op, operand: operand})
	}
}

// GetAssembly gets assembly instructions from a file
func GetAssembly(filename string) []string {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicln(err)
	}
	return strings.Split(strings.Trim(string(blob), "\n"), "\n")
}

func solvePart1(inputFile string) int {
	asm := GetAssembly(inputFile)
	v := new(VM)
	v.LoadAssembly(asm)
	return v.Run()
}

func main() {
	fmt.Printf("Part 1: %v\n", solvePart1("input.txt"))
}
