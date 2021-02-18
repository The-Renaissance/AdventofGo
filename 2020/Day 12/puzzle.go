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
	w       Waypoint
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

// NewShip creates a new ship heading east and a waypoint at 10 east, 1 north
func NewShip() *Ship {
	return &Ship{
		heading: 90,
		w: Waypoint{
			ew: 10,
			ns: 1,
		},
	}
}

// Waypoint denotes a waypoint's position relative to the ship
type Waypoint struct {
	ns, ew int
}

// East moves waypoint east
func (w *Waypoint) East(dist int) {
	w.ew += dist
}

// West moves waypoint west
func (w *Waypoint) West(dist int) {
	w.ew -= dist
}

// North moves waypoint north
func (w *Waypoint) North(dist int) {
	w.ns += dist
}

// South moves waypoint south
func (w *Waypoint) South(dist int) {
	w.ns -= dist
}

// Left rotates a waypoint left by deg degrees
func (w *Waypoint) Left(deg int) {
	rad := float64(deg) / 180.0 * math.Pi
	newWaypoint := Waypoint{
		ew: int(math.Round(math.Cos(rad)*float64(w.ew) - math.Sin(rad)*float64(w.ns))),
		ns: int(math.Round(math.Sin(rad)*float64(w.ew) + math.Cos(rad)*float64(w.ns))),
	}
	*w = newWaypoint
}

// Right rotates a waypoint right by deg degrees
func (w *Waypoint) Right(deg int) {
	w.Left(-deg)
}

// MovetoWaypoint moves the ship to the waypoint as many times as needed
func (s *Ship) MovetoWaypoint(times int) {
	for i := 0; i < times; i++ {
		s.ew += s.w.ew
		s.ns += s.w.ns
	}
}

// ExecuteInstructionPart2 executes an instruction according to Part 2
func (s *Ship) ExecuteInstructionPart2(ins string) error {
	dir := ins[0]
	dist, err := strconv.Atoi(ins[1:])
	if err != nil {
		return err
	}
	switch dir {
	case 'E':
		s.w.East(dist)
	case 'W':
		s.w.West(dist)
	case 'N':
		s.w.North(dist)
	case 'S':
		s.w.South(dist)
	case 'L':
		s.w.Left(dist)
	case 'R':
		s.w.Right(dist)
	case 'F':
		s.MovetoWaypoint(dist)
	default:
		return fmt.Errorf("%q is not a valid direction", dir)
	}
	return nil
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

func solvePart2() {
	instructions := getInput("input.txt")
	s := NewShip()
	for _, instruction := range instructions {
		s.ExecuteInstructionPart2(instruction)
	}
	fmt.Printf("Part 2: %v\n", s.GetDistance())
}

func main() {
	solvePart1()
	solvePart2()
}
