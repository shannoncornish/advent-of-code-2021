package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type item struct {
	point           point
	priority, index int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

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
	frontier := priorityQueue{
		&item{
			point:    origin,
			priority: 0,
		},
	}
	heap.Init(&frontier)

	risks := map[point]int{
		origin: 0,
	}

	for len(frontier) > 0 {
		current := heap.Pop(&frontier).(*item).point
		if current == destination {
			break
		}

		for _, next := range g.neighbors(current) {
			risk := risks[current] + g.values[next.y][next.x]
			if risks[next] == 0 || risk < risks[next] {
				risks[next] = risk
				heap.Push(&frontier, &item{
					point:    next,
					priority: risk,
				})
			}
		}
	}

	return risks[destination]
}

func (g *grid) enlarge(times int) *grid {
	size := point{
		x: g.size.x * times,
		y: g.size.y * times,
	}

	values := make([][]int, size.y)
	for y := 0; y < size.y; y++ {
		values[y] = make([]int, size.x)
	}

	for y := 0; y < times; y++ {
		for x := 0; x < times; x++ {
			offset := y + x
			for iy := 0; iy < g.size.y; iy++ {
				for ix := 0; ix < g.size.x; ix++ {
					value := offset + g.values[iy][ix]
					if value > 9 {
						value -= 9
					}

					values[(y*g.size.y)+iy][(x*g.size.x)+ix] = value
				}
			}
		}
	}

	return &grid{
		size:   size,
		values: values,
	}
}

func main() {
	g := parseInput()
	g = g.enlarge(5)

	lowestRisk := g.lowestRisk(point{x: 0, y: 0}, point{x: g.size.x - 1, y: g.size.y - 1})

	fmt.Printf("Lowest risk path: %d\n", lowestRisk)

	// Answer: 2842
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
