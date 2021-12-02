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
	for _, depth := range depths {
		if previous > 0 && depth > previous {
			increasesFromPreviousCount++
		}

		previous = depth
	}

	fmt.Printf("%d increases from the previous depth\n", increasesFromPreviousCount)

	// Answer: 1583
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
