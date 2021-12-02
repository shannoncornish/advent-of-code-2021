package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type command struct {
	direction string
	unit      int
}

func main() {
	commands := parseInput()

	var horizontal int
	var depth int

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontal += command.unit
		case "down":
			depth += command.unit
		case "up":
			depth -= command.unit
		default:
			log.Fatalf("Unsupported direction: %s\n", command.direction)
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d, Horizontal * Depth: %d\n", horizontal, depth, horizontal*depth)

	// Answer: 2120749
}

func parseInput() []command {
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

	var commands []command
	for scanner.Scan() {
		var command command
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &command.direction, &command.unit)
		if err != nil {
			log.Fatal(err)
		}

		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
