package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold point

type grid struct {
	size   point
	values [][]int
}

func (g *grid) print() {
	for y := 0; y <= g.size.y; y++ {
		for x := 0; x <= g.size.x; x++ {
			if g.values[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (g *grid) visible() int {
	var visible int
	for y := 0; y <= g.size.y; y++ {
		for x := 0; x <= g.size.x; x++ {
			if g.values[y][x] > 0 {
				visible++
			}
		}
	}

	return visible
}

func (g *grid) fold(fold fold) {
	original, new := g.size, g.size

	var hidden point
	if fold.x > 0 {
		hidden.x = fold.x + 1
		new.x = fold.x - 1
	} else {
		hidden.y = fold.y + 1
		new.y = fold.y - 1
	}

	for y := hidden.y; y <= original.y; y++ {
		for x := hidden.x; x <= original.x; x++ {
			if g.values[y][x] > 0 {
				p := point{x, y}
				if fold.x > 0 {
					p.x = fold.x - (x - fold.x)
				} else {
					p.y = fold.y - (y - fold.y)
				}

				g.values[p.y][p.x] = g.values[y][x]
			}
		}
	}

	g.size = new
}

func newGrid(dots []point) *grid {
	var size point
	for _, dot := range dots {
		if dot.x > size.x {
			size.x = dot.x
		}

		if dot.y > size.y {
			size.y = dot.y
		}
	}

	values := make([][]int, size.y+1)
	for y := 0; y <= size.y; y++ {
		values[y] = make([]int, size.x+1)
	}

	for _, dot := range dots {
		values[dot.y][dot.x] = 1
	}

	return &grid{
		size:   size,
		values: values,
	}
}

func main() {
	dots, folds := parseInput()

	grid := newGrid(dots)
	for _, fold := range folds {
		grid.fold(fold)
	}

	grid.print()

	// Answer: BCZRCEAB
}

func parseInput() ([]point, []fold) {
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

	var dots []point
	var folds []fold

	var instructions bool
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			instructions = true
			continue
		}

		if instructions {
			axis := line[11]
			value, _ := strconv.Atoi(line[13:])

			var fold fold
			switch rune(axis) {
			case 'y':
				fold.y = value
				break
			case 'x':
				fold.x = value
				break
			}

			folds = append(folds, fold)
		} else {
			values := strings.Split(line, ",")

			var dot point
			dot.x, _ = strconv.Atoi(values[0])
			dot.y, _ = strconv.Atoi(values[1])

			dots = append(dots, dot)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dots, folds
}
