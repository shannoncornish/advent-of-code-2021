package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type heightmap [][]int

type point struct {
	x, y int
}

func main() {
	hm := parseInput()

	lenY := len(hm)
	lenX := len(hm[0])

	var sum int
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			value := hm[y][x]
			if value == 9 {
				continue
			}

			adjacent := [4]int{-1, -1, -1, -1}

			if y > 0 {
				adjacent[0] = hm[y-1][x]
			}
			if y < lenY-1 {
				adjacent[1] = hm[y+1][x]
			}
			if x > 0 {
				adjacent[2] = hm[y][x-1]
			}
			if x < lenX-1 {
				adjacent[3] = hm[y][x+1]
			}

			if (adjacent[0] == -1 || value < adjacent[0]) &&
				(adjacent[1] == -1 || value < adjacent[1]) &&
				(adjacent[2] == -1 || value < adjacent[2]) &&
				(adjacent[3] == -1 || value < adjacent[3]) {
				sum += value + 1
			}
		}
	}

	fmt.Printf("The sum of the risk levels for all low points: %d\n", sum)

	// Answer: 539
}

func parseInput() heightmap {
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

	var lines heightmap
	for scanner.Scan() {
		text := scanner.Text()

		line := make([]int, len(text))
		for i, r := range text {
			line[i] = int(r - '0')
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
