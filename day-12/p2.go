package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type graph map[string]*node

type node struct {
	name string

	edges []*node
}

func count(s []*node, n *node) int {
	var c int
	for i := range s {
		if s[i] == n {
			c++
		}
	}

	return c
}

func clone(s []*node) []*node {
	c := make([]*node, len(s))
	for i := range s {
		c[i] = s[i]
	}

	return c
}

func path(visited []*node) string {
	names := make([]string, len(visited))
	for i := range visited {
		names[i] = visited[i].name
	}

	return strings.Join(names, ",")
}

func (n *node) visit(goal *node, visited []*node) []string {
	visited = append(visited, n)

	if n == goal {
		names := make([]string, len(visited))
		for i := range visited {
			names[i] = visited[i].name
		}

		return []string{strings.Join(names, ",")}
	}

	var paths []string
	for _, edge := range n.edges {
		if edge.name == "start" {
			continue
		}

		if isLower(edge.name) {

			visitCounts := make(map[*node]int)
			for _, n := range visited[1:] {
				if isLower(n.name) {
					visitCounts[n]++
				}
			}

			if visitCounts[edge] != 0 {
				var visitedTwice bool
				for _, visitCount := range visitCounts {
					if visitCount >= 2 {
						visitedTwice = true
						break
					}
				}

				if visitedTwice {
					continue
				}
			}
		}

		for _, path := range edge.visit(goal, clone(visited)) {
			paths = append(paths, path)
		}
	}

	return paths
}

func main() {
	graph := parseInput()

	paths := graph["start"].visit(graph["end"], nil)

	fmt.Printf("Paths through the cave system: %d\n", len(paths))

	// Answer: 116692
}

func isLower(s string) bool {
	return s == strings.ToLower(s)
}

func parseInput() graph {
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

	nodes := map[string]*node{}

	for scanner.Scan() {
		names := strings.Split(scanner.Text(), "-")
		leftName, rightName := names[0], names[1]

		leftNode, ok := nodes[leftName]
		if !ok {
			leftNode = &node{
				name: leftName,
			}
			nodes[leftName] = leftNode
		}

		rightNode, ok := nodes[rightName]
		if !ok {
			rightNode = &node{
				name: rightName,
			}
			nodes[rightName] = rightNode
		}

		leftNode.edges = append(leftNode.edges, rightNode)
		rightNode.edges = append(rightNode.edges, leftNode)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes
}
