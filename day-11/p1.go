package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type grid [10][10]int

type point struct {
	x, y int
}

func (g *grid) adjacent(p point) []point {
	points := []point{
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
	}

	var adjacent []point
	for _, point := range points {
		if point.x >= 0 && point.x <= 9 &&
			point.y >= 0 && point.y <= 9 {
			adjacent = append(adjacent, point)
		}
	}

	return adjacent
}

func (g *grid) flash(p point, flashes map[point]struct{}) {
	flashes[p] = struct{}{}

	stack := []point{p}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		for _, next := range g.adjacent(current) {
			g[next.y][next.x]++
			if g[next.y][next.x] > 9 {
				if _, ok := flashes[next]; !ok {
					stack = append(stack, next)
					flashes[next] = struct{}{}
				}

			}
		}
	}
}

func main() {
	grid := parseInput()

	var sum int
	for i := 0; i < 100; i++ {

		flashes := make(map[point]struct{})

		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				grid[y][x]++

				if grid[y][x] > 9 {
					point := point{x, y}
					if _, ok := flashes[point]; !ok {
						grid.flash(point, flashes)
					}
				}
			}
		}

		for flash := range flashes {
			grid[flash.y][flash.x] = 0
		}

		sum += len(flashes)
	}

	fmt.Printf("Total flashes after 100 steps: %d\n", sum)

	// Answer: 1683
}

func parseInput() *grid {
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

	var grid grid

	var row int
	for scanner.Scan() {
		for i, r := range scanner.Text() {
			grid[row][i] = int(r - '0')
		}
		row += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &grid
}
