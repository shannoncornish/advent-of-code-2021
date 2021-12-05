package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type line [2]point

type point struct {
	x, y int
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func points(l line) []point {
	var points []point

	p1 := l[0]
	p2 := l[1]

	if p1.x == p2.x {
		for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
			points = append(points, point{p1.x, y})
		}
	}

	if p1.y == p2.y {
		for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
			points = append(points, point{x, p1.y})
		}
	}

	return points
}

func main() {
	lines, size := parseInput()

	field := make([][]int, size.y+1)
	for y := 0; y <= size.y; y++ {
		field[y] = make([]int, size.x+1)
	}

	for _, line := range lines {
		points := points(line)
		for _, point := range points {
			field[point.y][point.x]++
		}
	}

	var count int
	for y := 0; y <= size.y; y++ {
		for x := 0; x <= size.x; x++ {
			if field[y][x] >= 2 {
				count++
			}
		}
	}

	fmt.Printf("The number of points where at least two lines overlap: %d\n", count)
}

func parseInput() ([]line, point) {
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

	var size point

	var lines []line
	for scanner.Scan() {
		var p1, p2 point
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &p1.x, &p1.y, &p2.x, &p2.y)
		if err != nil {
			log.Fatal(err)
		}

		if p1.x > size.x {
			size.x = p1.x
		}

		if p1.y > size.y {
			size.y = p1.y
		}

		if p2.x > size.x {
			size.x = p2.x
		}

		if p2.y > size.y {
			size.y = p2.y
		}

		lines = append(lines, line{p1, p2})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, size
}
