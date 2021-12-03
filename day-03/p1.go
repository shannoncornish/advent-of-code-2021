package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://stackoverflow.com/questions/23192262/how-would-you-set-and-clear-a-single-bit-in-go

func main() {
	values, size := parseInput()

	var gamma uint
	var epsilon uint

	for pos := size - 1; pos >= 0; pos-- {
		var count int
		for i := range values {
			value := values[i] & (1 << pos)
			if value > 0 {
				count++
			}
		}

		if count > len(values)/2 {
			gamma |= (1 << pos)
		} else {
			epsilon |= (1 << pos)
		}
	}

	log.Printf("Gamma: %08b\n", gamma)
	log.Printf("Epsilon: %08b\n", epsilon)

	fmt.Printf("Gamma: %d, Epsilon: %d, Power Consumption: %d\n", gamma, epsilon, gamma*epsilon)

	// Answer: 4191876
}

func parseInput() ([]uint, int) {
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

	var size int
	var values []uint
	for scanner.Scan() {
		text := scanner.Text()
		if size == 0 {
			size = len(text)
		}

		value, err := strconv.ParseUint(text, 2, 32)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, uint(value))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values, size
}
