package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"strings"
)

//// Part 1 ////

func earliestBus(ts, shuttles string) (id int, minutes int) {
	timestamp, _ := strconv.Atoi(ts)
	for _, busline := range strings.Split(shuttles, ",") {
		if n, err := strconv.Atoi(busline); err == nil {
			if timestamp%n == 0 {
				return n, 0
			}
			if minutes == 0 || n-timestamp%n < minutes {
				minutes = n - timestamp%n
				id = n
			}
		}
	}
	return id, minutes
}

//// Part 2 ////

type congruence struct {
	remainder, divisor int64
}

// Chinese remainder theorem
func reduce(c1, c2 congruence) congruence {
	a1, a2 := c1.remainder, c2.remainder
	n1, n2 := big.NewInt(c1.divisor), big.NewInt(c2.divisor)
	m1, m2, z := new(big.Int), new(big.Int), new(big.Int)
	z.GCD(m1, m2, n1, n2)
	result := congruence{
		remainder: a1*m2.Int64()*n2.Int64() + a2*m1.Int64()*n1.Int64(),
		divisor:   n1.Int64() * n2.Int64(),
	}

	// round to smallest positive integer
	result.remainder %= result.divisor
	if result.remainder < 0 {
		result.remainder += result.divisor
	}
	return result
}

func fold(congruences []congruence, f func(congruence, congruence) congruence) congruence {
	c1, c2 := congruences[0], congruences[1]
	if len(congruences) == 2 {
		return f(c1, c2)
	}
	return fold(append([]congruence{f(c1, c2)}, congruences[2:]...), f)
}

func getShuttlePart2(shuttles string) int {
	congruences := []congruence{}
	for i, IDstring := range strings.Split(shuttles, ",") {
		id, err := strconv.Atoi(IDstring)
		if err != nil {
			continue
		}
		if i == 0 {
			congruences = append(congruences, congruence{
				remainder: 0,
				divisor:   int64(id),
			})
		} else {
			congruences = append(congruences, congruence{
				remainder: int64(id - i%id),
				divisor:   int64(id),
			})
		}
	}
	result := fold(congruences, reduce)
	return int(result.remainder)
}

func getInput(filename string) []string {
	l, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Reading", filename, "failed:", err)
	}
	return strings.Split(strings.Trim(string(l), "\n"), "\n")
}

func solvePart1() {
	input := getInput("input.txt")
	id, minutes := earliestBus(input[0], input[1])
	fmt.Printf("Part 1: %v\n", id*minutes)
}

func solvePart2() {
	input := getInput("input.txt")
	minutes := getShuttlePart2(input[1])
	fmt.Printf("Part 2: %v\n", minutes)
}

func verifyPart2(ts int, input string) bool {
	result := true
	for dt, bus := range strings.Split(input, ",") {
		if bus == "x" {
			continue
		}
		id, _ := strconv.Atoi(bus)
		if (ts+dt)%id != 0 {
			result = false
		}
	}
	return result
}

func main() {
	// solvePart1()
	// solvePart2()
	
	// Due to issues with integer overflow, I was unable to finish the problem with my program alone.
	// I used https://www.dcode.fr/chinese-remainder to get my answer
	result := 626670513163231
	fmt.Printf("%v: %v\n", result, verifyPart2(result, getInput("input.txt")[1]))
}
