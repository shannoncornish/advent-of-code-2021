package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateFuel(crabsByPosition map[int]int, destination int) int {
	var fuel int
	for position, crabs := range crabsByPosition {
		if position == destination {
			continue
		}

		difference := max(position, destination) - min(position, destination)
		fuel += difference * crabs
	}

	return fuel
}

func main() {
	positions := parseInput()

	crabsByPosition := make(map[int]int)
	for _, position := range positions {
		crabsByPosition[position]++
	}

	var mostPopulatedPosition int
	for position, crabs := range crabsByPosition {
		if crabsByPosition[mostPopulatedPosition] < crabs {
			mostPopulatedPosition = position
		}
	}

	position := mostPopulatedPosition
	for {
		fuelForLeft := calculateFuel(crabsByPosition, position-1)
		fuelForCenter := calculateFuel(crabsByPosition, position)
		fuelForRight := calculateFuel(crabsByPosition, position+1)

		if fuelForCenter < fuelForLeft {
			if fuelForCenter < fuelForRight {
				fmt.Printf("Fuel required for position: %d (position %d)\n", fuelForCenter, position)
				break
			} else {
				position++
			}
		} else {
			position--
		}
	}

	// Answer: 333755
}

func parseInput() []int {
	name := "input.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Input file: %s\n", name)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Split(string(bytes), ",")

	values := make([]int, len(numbers))
	for i, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		values[i] = int(value)
	}

	return values
}
