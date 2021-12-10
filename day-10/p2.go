package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func openChunks(line string) []rune {
	var stack []rune
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
			break
		case ')', ']', '}', '>':
			expected := pairs[stack[len(stack)-1]]
			if expected == r {
				stack = stack[:len(stack)-1]
			} else {
				return []rune{}
			}

			break
		}
	}

	return stack
}

func main() {
	lines := parseInput()

	var scores []int

	for _, line := range lines {
		open := openChunks(line)
		if len(open) > 0 {
			var score int
			for i := len(open) - 1; i >= 0; i-- {
				score *= 5
				switch pairs[open[i]] {
				case ')':
					score += 1
					break
				case ']':
					score += 2
					break
				case '}':
					score += 3
					break
				case '>':
					score += 4
					break
				}
			}

			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	middle := scores[len(scores)/2]

	fmt.Printf("The middle score: %d\n", middle)

	// Answer: 1105996483
}

func parseInput() []string {
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

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
