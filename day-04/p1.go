package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board [5][5]int
type vector [2]int

func (b *board) findNumber(number int) []vector {
	var vectors []vector

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == number {
				vectors = append(vectors, vector{i, j})
			}
		}
	}

	return vectors
}

func (b *board) findComplete(vectors []vector) (vector, bool) {
	var rows [5]byte
	var columns [5]byte

	for _, vector := range vectors {
		rows[vector[0]] |= (1 << vector[1])
		columns[vector[1]] |= (1 << vector[0])
	}

	for i, row := range rows {
		if row == byte(0b00011111) {
			return vector{i, -1}, true
		}
	}

	for i, column := range columns {
		if column == byte(0b00011111) {
			return vector{-1, i}, true
		}
	}

	return vector{-1, -1}, false
}

func main() {
	numbers, boards := parseInput()

	found := make(map[int][]vector)

	var lastNumberCalled int

	winningBoardIndex := -1
	for _, number := range numbers {
		for b, board := range boards {
			vectors := board.findNumber(number)
			if len(vectors) > 0 {
				found[b] = append(found[b], vectors...)
				if len(found[b]) > 4 {
					if _, ok := board.findComplete(found[b]); ok {
						winningBoardIndex = b
						break
					}
				}
			}
		}

		if winningBoardIndex > -1 {
			lastNumberCalled = number
			break
		}
	}

	winningBoard := boards[winningBoardIndex]
	for _, vector := range found[winningBoardIndex] {
		winningBoard[vector[0]][vector[1]] = 0
	}

	var sum int
	for _, row := range winningBoard {
		for _, value := range row {
			sum += value
		}
	}

	log.Printf("Sum of all unmarked numbers: %d\n", sum)
	log.Printf("Last number called: %d\n", lastNumberCalled)

	fmt.Printf("Final score: %d\n", sum*lastNumberCalled)

	// Answer: 44088
}

func parseInput() ([]int, []board) {
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

	scanner.Scan()
	var numbers []int
	for _, number := range strings.Split(scanner.Text(), ",") {
		i, err := strconv.ParseInt(number, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, int(i))
	}

	var boards []board
	for scanner.Scan() {
		var board board
		for i := 0; i < 5; i++ {
			scanner.Scan()
			_, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d", &board[i][0], &board[i][1], &board[i][2], &board[i][3], &board[i][4])
			if err != nil {
				log.Fatal(err)
			}
		}

		boards = append(boards, board)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers, boards
}
