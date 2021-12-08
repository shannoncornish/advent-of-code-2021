package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var uniqueSegmentsToDigits = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

type entry struct {
	uniqueSignalPatterns, outputValue []string
}

func main() {
	entries := parseInput()

	var count int
	for _, entry := range entries {
		for _, outputValue := range entry.outputValue {
			_, ok := uniqueSegmentsToDigits[len(outputValue)]
			if !ok {
				continue
			}

			count++
		}
	}

	fmt.Printf("1, 4, 7, or 8 appear %d many times in the output values\n", count)

	// Answer: 239
}

func parseInput() []entry {
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

	var entries []entry
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		entries = append(entries, entry{
			uniqueSignalPatterns: s[0:10],
			outputValue:          s[11:],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}
