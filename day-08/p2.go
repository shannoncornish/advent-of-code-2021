package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://stackoverflow.com/questions/19599364/how-to-convert-array-of-integers-into-an-integer-in-c

var uniqueSegmentsToDigits = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

type entry struct {
	uniqueSignalPatterns, outputValue []string
}

func intersecting(a, b string) int {
	var count int
	for _, r := range a {
		if strings.ContainsRune(b, r) {
			count++
		}
	}

	return count
}

func mappingForUniqueSignalPatterns(uniqueSignalPatterns []string) map[int]string {
	digitPatterns := make(map[int]string)

	var patterns []string
	for _, pattern := range uniqueSignalPatterns {
		if digit, ok := uniqueSegmentsToDigits[len(pattern)]; ok {
			digitPatterns[digit] = pattern
			continue
		}

		patterns = append(patterns, pattern)
	}

	// Known patterns: 1, 4, 7, 8

	// Pattern length: 5
	// 2 (1:1, 4:2, 7:2, 8:5)
	// 3 (1:2, 4:3, 7:3, 8:5)
	// 5 (1:1, 4:3, 7:2, 8:5)

	// Pattern length: 6
	// 0 (1:2, 4:3, 7:2, 8:6)
	// 6 (1:1, 4:3, 7:2, 8:6)
	// 9 (1:2, 4:4, 7:3, 8:6)

	for _, pattern := range patterns {
		switch len(pattern) {
		case 5:
			switch intersecting(pattern, digitPatterns[1]) {
			case 1:
				switch intersecting(pattern, digitPatterns[4]) {
				case 2:
					digitPatterns[2] = pattern
					break
				case 3:
					digitPatterns[5] = pattern
					break
				}
				break
			case 2:
				digitPatterns[3] = pattern
				break
			}
			break
		case 6:
			switch intersecting(pattern, digitPatterns[1]) {
			case 1:
				digitPatterns[6] = pattern
				break
			case 2:
				switch intersecting(pattern, digitPatterns[4]) {
				case 3:
					digitPatterns[0] = pattern
					break
				case 4:
					digitPatterns[9] = pattern
					break
				}
			}

			break
		}
	}

	return digitPatterns
}

func main() {
	entries := parseInput()

	var sum int
	for _, entry := range entries {
		var value int

		mapping := mappingForUniqueSignalPatterns(entry.uniqueSignalPatterns)
		for _, outputValue := range entry.outputValue {
			for digit, pattern := range mapping {
				if len(outputValue) == len(pattern) && len(pattern) == intersecting(pattern, outputValue) {
					value = 10*value + digit
				}
			}
		}

		sum += value
	}

	fmt.Printf("The sum of all the output values is %d\n", sum)

	// Answer: 946346
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
