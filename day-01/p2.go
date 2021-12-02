package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	depths := parseInput()

	var increasesFromPreviousCount int
	var previous int
	for i := range depths[:len(depths)-2] {
		window := depths[i : i+3]

		var sum int
		for _, depth := range window {
			sum += depth
		}

		if previous > 0 && sum > previous {
			increasesFromPreviousCount++
		}

		previous = sum
	}

	fmt.Printf("%d increases from the previous sum\n", increasesFromPreviousCount)

	// Answer: 1627
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

	scanner := bufio.NewScanner(f)

	var results []int
	for scanner.Scan() {
		result, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, int(result))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
