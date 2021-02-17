package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// SeatMap stores seat state information
type SeatMap [][]rune

// PPrint pretty prints a SeatMap
func (m SeatMap) PPrint() string {
	rows := make([]string, len(m))
	for i := range rows {
		rows[i] = string(m[i])
	}
	return strings.Join(rows, "\n")
}

// FloorState stores the previous and next states of a seat floor
type FloorState struct {
	previous SeatMap
	current  SeatMap
}

// Position encodes a seat's row and column in the map
type Position struct {
	row, column int
}

// AdjacentFunc is a type of function that accepts position and seat state and returns adjacent positions
type AdjacentFunc func(*FloorState, Position) []Position

// AdvanceState runs one step of the simulation
func (s *FloorState) AdvanceState(occupiedSeatQuota int, adjacentFunc AdjacentFunc) {
	if s.previous == nil {
		s.previous = make([][]rune, len(s.current))
	}
	for i := range s.previous {
		s.previous[i] = make([]rune, len(s.current[0]))
		copy(s.previous[i], s.current[i])
	}
	for i, row := range s.previous {
		for j, seat := range row {
			adjacents := adjacentFunc(s, Position{i, j})
			if seat == 'L' {
				noOccupied := true
				for _, adjseat := range adjacents {
					if s.previous[adjseat.row][adjseat.column] == '#' {
						noOccupied = false
						break
					}
				}
				if noOccupied {
					s.current[i][j] = '#'
				}
			} else if seat == '#' {
				occupiedSeatCount := 0
				for _, adjseat := range adjacents {
					if s.previous[adjseat.row][adjseat.column] == '#' {
						occupiedSeatCount++
					}
				}
				if occupiedSeatCount >= occupiedSeatQuota {
					s.current[i][j] = 'L'
				}
			}
		}
	}
}

// RunandCount runs the simulation up until previous == current then counts occupied seats
func (s *FloorState) RunandCount(quota int, adjacentFunc AdjacentFunc) int {
	for !reflect.DeepEqual(s.previous, s.current) {
		s.AdvanceState(quota, adjacentFunc)
	}
	occupiedSeatCount := 0
	for _, row := range s.current {
		for _, seat := range row {
			if seat == '#' {
				occupiedSeatCount++
			}
		}
	}
	return occupiedSeatCount
}

// GetImmediateAdjacents returns the adjacent seats immediately surrounding position p
func GetImmediateAdjacents(s *FloorState, p Position) []Position {
	rows, columns := len(s.current), len(s.current[0])
	i, j := p.row, p.column
	var pos []Position
	if i > 0 && j > 0 {
		pos = append(pos, Position{i - 1, j - 1})
	}
	if i > 0 {
		pos = append(pos, Position{i - 1, j})
	}
	if i > 0 && j+1 < columns {
		pos = append(pos, Position{i - 1, j + 1})
	}
	if j > 0 {
		pos = append(pos, Position{i, j - 1})
	}
	if j+1 < columns {
		pos = append(pos, Position{i, j + 1})
	}
	if i+1 < rows && j > 0 {
		pos = append(pos, Position{i + 1, j - 1})
	}
	if i+1 < rows {
		pos = append(pos, Position{i + 1, j})
	}
	if i+1 < rows && j+1 < columns {
		pos = append(pos, Position{i + 1, j + 1})
	}
	return pos
}

// FromInitialState returns a FloorState struct from an initial state
func FromInitialState(init []string) *FloorState {
	fs := FloorState{
		current: make([][]rune, 0, len(init)),
	}
	for _, row := range init {
		fs.current = append(fs.current, []rune(row))
	}
	return &fs
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1() {
	const occupiedSeatQuota = 4
	init := getInput("input.txt")
	fs := FromInitialState(init)
	fmt.Printf("Part 1: %v\n", fs.RunandCount(occupiedSeatQuota, GetImmediateAdjacents))
}

func main() {
	solvePart1()
}
