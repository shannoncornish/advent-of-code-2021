package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type heightmap struct {
	points [][]int
	size   point
}

type point struct {
	x, y int
}

type basin map[point]int

func (h *heightmap) mapBasin(p point, b basin) {
	b[p] = h.points[p.y][p.x]

	adjacent := make(map[point]int)
	if p.y > 0 {
		adjacent[point{x: p.x, y: p.y - 1}] = h.points[p.y-1][p.x]
	}
	if p.y < h.size.y-1 {
		adjacent[point{x: p.x, y: p.y + 1}] = h.points[p.y+1][p.x]
	}
	if p.x > 0 {
		adjacent[point{x: p.x - 1, y: p.y}] = h.points[p.y][p.x-1]
	}
	if p.x < h.size.x-1 {
		adjacent[point{x: p.x + 1, y: p.y}] = h.points[p.y][p.x+1]
	}

	for ap, av := range adjacent {
		if av == 9 {
			continue
		}

		if _, ok := b[ap]; !ok {
			h.mapBasin(ap, b)
		}
	}
}

func (h *heightmap) fillBasin(b basin) {
	for p := range b {
		h.points[p.y][p.x] = 9
	}
}

func main() {
	hm := parseInput()

	var sizes []int

	for y := 0; y < hm.size.y; y++ {
		for x := 0; x < hm.size.x; x++ {
			value := hm.points[y][x]
			if value >= 9 {
				continue
			}

			b := make(basin)
			hm.mapBasin(point{x: x, y: y}, b)
			hm.fillBasin(b)

			sizes = append(sizes, len(b))
		}
	}

	sort.Ints(sizes)

	largest := sizes[len(sizes)-3:]

	fmt.Printf("The 3 largest sizes multipled together: %d\n", largest[0]*largest[1]*largest[2])

	// Answer: 736920
}

func parseInput() *heightmap {
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

	var lines [][]int
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

	return &heightmap{
		points: lines,
		size: point{
			x: len(lines[0]),
			y: len(lines),
		},
	}
}
