package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Player struct {
	Number   int
	Score    int
	Position int
}

func (p *Player) Move(n int) {
	p.Position = (p.Position-1+n)%10 + 1
	p.Score += p.Position
}

func NewPlayer(number, startingPosition int) *Player {
	player := Player{
		Number:   number,
		Position: startingPosition,
	}

	return &player
}

var deterministicDice, rollCount int

func RollDice() int {
	rollCount++
	deterministicDice++

	if deterministicDice > 100 {
		deterministicDice = 1
	}

	return deterministicDice
}

func main() {
	players := parseInput()
	for {
		var over bool
		for _, player := range players {
			var n int
			for roll := 0; roll < 3; roll++ {
				n += RollDice()
			}

			player.Move(n)

			if player.Score >= 1000 {
				over = true
				break
			}
		}

		if over {
			break
		}
	}

	var loser *Player
	for _, player := range players {
		if loser == nil || player.Score < loser.Score {
			loser = player
		}
	}

	fmt.Printf("The score of the losing player multipled by roll count: %d\n", loser.Score*rollCount)

	// Answer: 888735
}

func parseInput() []*Player {
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

	var players []*Player
	for scanner.Scan() {
		var number, startingPosition int
		_, err := fmt.Sscanf(scanner.Text(), "Player %d starting position: %d", &number, &startingPosition)
		if err != nil {
			log.Fatal(err)
		}

		players = append(players, NewPlayer(number, startingPosition))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return players
}
