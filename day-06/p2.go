package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func group(numbers []int) map[int]int {
	m := make(map[int]int)
	for _, number := range numbers {
		m[number]++
	}

	return m
}

func main() {
	state := parseInput()

	existingTimers := group(state)

	for day := 0; day < 256; day++ {
		newTimers := make(map[int]int)
		for k, v := range existingTimers {
			newTimers[k-1] = v
		}
		spawning := newTimers[-1]
		delete(newTimers, -1)
		newTimers[6] += spawning
		newTimers[8] += spawning

		existingTimers = newTimers
	}

	var count int
	for _, v := range existingTimers {
		count += v
	}

	fmt.Printf("Number of laternfish after 256 days: %d\n", count)

	// Answer: 1572358335990
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
