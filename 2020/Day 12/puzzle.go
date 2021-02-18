package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// Ship denotes a ship's position
type Ship struct {
	ew, ns  int
	heading int
}

// East moves ship east
func (s *Ship) East(dist int) {
	s.ew += dist
}

// West moves ship west
func (s *Ship) West(dist int) {
	s.ew -= dist
}

// North moves ship north
func (s *Ship) North(dist int) {
	s.ns += dist
}

// South moves ship south
func (s *Ship) South(dist int) {
	s.ns -= dist
}

// GetDistance calculates the Manhattan Distance between current and the starting location
func (s *Ship) GetDistance() int {
	return int(math.Abs(float64(s.ns)) + math.Abs(float64(s.ew)))
}

// Left turns ship left by deg degrees
func (s *Ship) Left(deg int) {
	s.heading = (s.heading - deg) % 360
	if s.heading < 0 {
		s.heading += 360
	}
}

// Right turns ship right by deg degrees
func (s *Ship) Right(deg int) {
	s.heading = (s.heading + deg) % 360
}

// Forward moves the ship forward along its heading
// The answer to the puzzle is in integers. This means the ship's heading must be a multiple of 90.
func (s *Ship) Forward(dist int) {
	switch s.heading {
	case 0:
		s.ns += dist
	case 90:
		s.ew += dist
	case 180:
		s.ns -= dist
	case 270:
		s.ew -= dist
	default:
		log.Panicf("Ship's heading is %v\n", s.heading)
	}
}

// ExecuteInstruction executes an instruction in the form of "F30"
func (s *Ship) ExecuteInstruction(ins string) error {
	dir := ins[0]
	dist, err := strconv.Atoi(ins[1:])
	if err != nil {
		return err
	}
	switch dir {
	case 'E':
		s.East(dist)
	case 'W':
		s.West(dist)
	case 'N':
		s.North(dist)
	case 'S':
		s.South(dist)
	case 'L':
		s.Left(dist)
	case 'R':
		s.Right(dist)
	case 'F':
		s.Forward(dist)
	default:
		return fmt.Errorf("%q is not a valid direction", dir)
	}
	return nil
}

// NewShip creates a new ship heading east
func NewShip() *Ship {
	return &Ship{
		heading: 90,
	}
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
	s := NewShip()
	for _, instruction := range instructions {
		s.ExecuteInstruction(instruction)
	}
	fmt.Printf("Part 1: %v\n", s.GetDistance())
}

func main() {
	solvePart1()
}
