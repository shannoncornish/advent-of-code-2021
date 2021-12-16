package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

type grid struct {
	size   point
	values [][]int
}

func (g *grid) neighbors(p point) []point {
	var neighbors []point

	for _, offset := range []point{
		{x: 0, y: -1},
		{x: 0, y: +1},
		{x: -1, y: 0},
		{x: +1, y: 0},
	} {
		neighbor := point{x: p.x + offset.x, y: p.y + offset.y}
		if neighbor.x >= 0 &&
			neighbor.y >= 0 &&
			neighbor.x < g.size.x &&
			neighbor.y < g.size.y {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (g *grid) lowestRisk(origin, destination point) int {
	frontier := []point{
		origin,
	}

	risks := map[point]int{
		origin: 0,
	}

	for len(frontier) > 0 {
		current := frontier[0]
		if current == destination {
			break
		}

		frontier = frontier[1:]
		for _, next := range g.neighbors(current) {
			risk := risks[current] + g.values[next.y][next.x]
			if risks[next] == 0 || risk < risks[next] {
				risks[next] = risk
				frontier = append(frontier, next)
			}
		}
	}

	return risks[destination]
}

func main() {
	g := parseInput()

	lowestRisk := g.lowestRisk(point{x: 0, y: 0}, point{x: g.size.x - 1, y: g.size.y - 1})

	fmt.Printf("Lowest risk path: %d\n", lowestRisk)

	// Correct Answer: 435
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

	var values [][]int
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))
		for i, r := range line {
			row[i], _ = strconv.Atoi(string(r))
		}

		values = append(values, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &grid{
		size: point{
			x: len(values[0]),
			y: len(values),
		},
		values: values,
	}
}
