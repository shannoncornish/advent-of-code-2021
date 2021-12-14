package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	template, rules := parseInput()

	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for step := 1; step <= 40; step++ {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			element := rules[pair]

			newPairs[string([]byte{pair[0], element})] += count
			newPairs[string([]byte{element, pair[1]})] += count
		}

		pairs = newPairs
	}

	quantities := make(map[byte]int)
	for pair, count := range pairs {
		quantities[pair[0]] += count
	}

	quantities[template[len(template)-1]] += 1

	var most, least int
	for _, count := range quantities {
		if count > most {
			most = count
		}
		if least == 0 || count < least {
			least = count
		}
	}

	fmt.Printf("Quantity of the most common element subtracted the quantity of the least common element: %d\n", most-least)

	// Answer: 2158894777814
}

func parseInput() (string, map[string]byte) {
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

	scanner.Scan()
	template := scanner.Text()

	scanner.Scan()

	rules := make(map[string]byte)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " -> ")
		rules[elements[0]] = elements[1][0]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return template, rules
}
