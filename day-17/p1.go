package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
	x, y int
}

type area struct {
	bottomLeft, topRight point
}

func (a *area) contains(p point) bool {
	return p.x >= a.bottomLeft.x &&
		p.y >= a.bottomLeft.y &&
		p.x <= a.topRight.x &&
		p.y <= a.topRight.y
}

type probe struct {
	position, velocity point
}

func fire(velocity point) *probe {
	return &probe{
		velocity: velocity,
	}
}

func (p *probe) step() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y

	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}

	p.velocity.y--
}

func main() {
	target := parseInput()

	var velocities []point
	for y := -1000; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			velocities = append(velocities, point{x: x, y: y})
		}
	}

	hits := make(map[point]int)
	for _, velocity := range velocities {
		p := fire(velocity)

		var maxY int
		var hit bool
		for {
			p.step()

			if p.position.y > maxY || maxY == 0 {
				maxY = p.position.y
			}

			if target.contains(p.position) {
				hit = true
				break
			}

			if p.position.y < target.bottomLeft.y || p.position.x > target.topRight.x {
				break
			}
		}

		if hit {
			hits[velocity] = maxY
		}
	}

	var maxY int
	for _, y := range hits {
		if y > maxY || maxY == 0 {
			maxY = y
		}
	}

	fmt.Printf("Highest y position: %v\n", maxY)

	// Answer: 8256
}

func parseInput() area {
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

	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var x1, x2, y1, y2 int
	fmt.Sscanf(string(contents), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	return area{
		bottomLeft: point{
			x: x1,
			y: y1,
		},
		topRight: point{
			x: x2,
			y: y2,
		},
	}
}
