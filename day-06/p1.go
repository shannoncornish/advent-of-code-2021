package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func simulateDay(previousDayFishes []int) []int {
	var newFish []int
	existingFish := make([]int, len(previousDayFishes))

	for i, previousDayTimer := range previousDayFishes {
		newTimer := previousDayTimer - 1
		if newTimer == -1 {
			newTimer = 6
			newFish = append(newFish, 8)
		}
		existingFish[i] = newTimer
	}

	return append(existingFish, newFish...)
}

func main() {
	initialState := parseInput()

	simulations := [][]int{initialState}

	for day := 0; day <= 80; day++ {
		previousDaysState := simulations[day]

		newValues := simulateDay(previousDaysState)
		simulations = append(simulations, newValues)
	}

	previousToLastSimulation := simulations[len(simulations)-2]
	fmt.Printf("length: %d\n", len(previousToLastSimulation))

	// Answer: 346063
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
		value, err := strconv.ParseInt(number, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		values[i] = int(value)
	}

	return values
}
