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
	v.Reset()
	for !v.visitedInstructions[v.programCounter] && v.programCounter < len(v.memory) {
		v.visitedInstructions[v.programCounter] = true
		v.step()
	}
	return v.accumulator
}

// Reset resets the VM
func (v *VM) Reset() {
	v.accumulator = 0
	v.programCounter = 0
	v.visitedInstructions = make(map[int]bool, len(v.memory))
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
	for _, line := range asm {
		tokens := strings.Split(line, " ")
		op := opcodeFromToken(tokens[0])
		operand, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Panicln(err)
		}
		v.memory = append(v.memory, instruction{op: op, operand: operand})
	}
	v.Reset()
}

func (v *VM) patchMemory(address int) {
	switch v.memory[address].op {
	case jmp:
		v.memory[address].op = nop
	case nop:
		v.memory[address].op = jmp
	case acc:
		log.Panicln("acc instruction at address", address, "should not be patched!")
	}
}

// DiagnoseAndFix finds the address that needs to be patched and patch it, then return the address of the patch and errors if any.
//
// It patches every single nop or jmp until the program runs without infinite loops
func (v *VM) DiagnoseAndFix() (patchAddr int, err error) {
	addrs := v.findAllJmpsandNops()
	for _, addr := range addrs {
		v.patchMemory(addr)
		_ = v.Run()
		if !v.exitedWithInfiniteLoop() {
			return addr, nil
		}
		v.patchMemory(addr)
	}
	return 0, fmt.Errorf("This program cannot be fixed")
}

func (v *VM) findAllJmpsandNops() []int {
	out := make([]int, 0)
	for i, instruction := range v.memory {
		if instruction.op == jmp || instruction.op == nop {
			out = append(out, i)
		}
	}
	return out
}

//InfiniteLoop determines if the program that has been run has exited with an infinite loop
func (v *VM) exitedWithInfiniteLoop() bool {
	return v.programCounter < len(v.memory)
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
	result := v.Run()
	return result
}

func solvePart2(inputFile string) int {
	asm := GetAssembly(inputFile)
	v := new(VM)
	v.LoadAssembly(asm)
	if _, err := v.DiagnoseAndFix(); err != nil {
		log.Panicln(err)
	}
	result := v.Run()
	return result
}

func main() {
	fmt.Printf("Part 1: %v\n", solvePart1("input.txt"))
	fmt.Printf("Part 2: %v\n", solvePart2("input.txt"))
}
