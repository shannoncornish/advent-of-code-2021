package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	template, rules := parseInput()

	polymer := list.New()
	for _, element := range template {
		polymer.PushBack(byte(element))
	}

	for step := 1; step <= 10; step++ {
		for current := polymer.Front(); current.Next() != nil; current = current.Next() {
			next := current.Next()

			pair := string([]byte{
				current.Value.(byte),
				next.Value.(byte),
			})

			element, ok := rules[pair]
			if !ok {
				log.Fatalf("Missing rule for pair: %s", pair)
			}

			current = polymer.InsertAfter(element, current)
		}
	}

	quantities := make(map[byte]int)
	for current := polymer.Front(); current != nil; current = current.Next() {
		quantities[current.Value.(byte)]++
	}

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

	// Answer: 2068
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
