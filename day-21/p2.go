package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var combinations = map[int]int{
	3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1,
}

var wins = map[int]int64{
	1: 0, 2: 0,
}

type GameState struct {
	position1, position2, score1, score2 int
}

func roll(state GameState, gameCount int) map[GameState]int {
	interimStates := make(map[GameState]int)
	for roll, rollCount := range combinations {
		newPosition1 := (state.position1-1+roll)%10 + 1
		newScore1 := state.score1 + newPosition1
		if newScore1 >= 21 {
			wins[1] += int64(rollCount * gameCount)
			continue
		}

		interimStates[GameState{
			position1: newPosition1,
			position2: state.position2,
			score1:    newScore1,
			score2:    state.score2,
		}] += (gameCount * rollCount)
	}

	newStates := make(map[GameState]int)
	for interimState, interimGameCount := range interimStates {
		for roll, rollCount := range combinations {
			newPosition2 := (state.position2-1+roll)%10 + 1
			newScore2 := interimState.score2 + newPosition2
			if newScore2 >= 21 {
				wins[2] += int64(rollCount * interimGameCount)
				continue
			}

			newStates[GameState{
				position1: interimState.position1,
				position2: newPosition2,
				score1:    interimState.score1,
				score2:    newScore2,
			}] += (rollCount * interimGameCount)
		}
	}

	return newStates
}

func main() {
	startingPositions := parseInput()

	states := map[GameState]int{
		{
			position1: startingPositions[0],
			position2: startingPositions[1],
		}: 1,
	}

	for {
		newStates := make(map[GameState]int)
		for state, count := range states {
			for newState, newCount := range roll(state, count) {
				newStates[newState] += newCount
			}
		}

		states = newStates
		if len(states) == 0 {
			break
		}
	}

	winner := 1
	if wins[2] > wins[1] {
		winner = 2
	}

	fmt.Printf("Number of universes the winner wins: %d, player %d\n", wins[winner], winner)

	// Answer: 647608359455719
}

func parseInput() []int {
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

	var startingPositions []int
	for scanner.Scan() {
		var player, startingPosition int
		_, err := fmt.Sscanf(scanner.Text(), "Player %d starting position: %d", &player, &startingPosition)
		if err != nil {
			log.Fatal(err)
		}

		startingPositions = append(startingPositions, startingPosition)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return startingPositions
}
