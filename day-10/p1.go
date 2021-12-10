package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func indexIllegalCharacter(line string) int {
	var stack []rune
	for i, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
			break
		case ')', ']', '}', '>':
			expected := pairs[stack[len(stack)-1]]
			if expected == r {
				stack = stack[:len(stack)-1]
			} else {
				return i
			}

			break
		}
	}

	return -1
}

func main() {
	lines := parseInput()

	var counts = map[rune]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}

	for _, line := range lines {
		index := indexIllegalCharacter(line)
		if index != -1 {
			counts[rune(line[index])]++
		}
	}

	score := (3 * counts[')']) + (57 * counts[']']) + (1197 * counts['}']) + (25137 * counts['>'])

	fmt.Printf("Total syntax error score: %d\n", score)

	// Answer: 215229
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
