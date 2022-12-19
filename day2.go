package main

import (
	"fmt"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

const (
	LOSS = 0
	DRAW = 3
	WIN  = 6
)

func TournamentScorePart1(rounds []string) int {
	roundScores := map[string]int{
		// opponent: rock
		"A X": DRAW + ROCK,
		"A Y": WIN + PAPER,
		"A Z": LOSS + SCISSORS,
		// opponent: paper
		"B X": LOSS + ROCK,
		"B Y": DRAW + PAPER,
		"B Z": WIN + SCISSORS,
		// opponent: scissors
		"C X": WIN + ROCK,
		"C Y": LOSS + PAPER,
		"C Z": DRAW + SCISSORS,
	}
	score := 0
	for _, round := range rounds {
		score += roundScores[round]
	}
	return score
}

func TournamentScorePart2(rounds []string) int {
	roundScores := map[string]int{
		// opponent: rock
		"A X": LOSS + SCISSORS,
		"A Y": DRAW + ROCK,
		"A Z": WIN + PAPER,
		// opponent: paper
		"B X": LOSS + ROCK,
		"B Y": DRAW + PAPER,
		"B Z": WIN + SCISSORS,
		// opponent: scissors
		"C X": LOSS + PAPER,
		"C Y": DRAW + SCISSORS,
		"C Z": WIN + ROCK,
	}
	score := 0
	for _, round := range rounds {
		score += roundScores[round]
	}
	return score
}

func day2(input []string) {
	fmt.Println(TournamentScorePart1(input))
	fmt.Println(TournamentScorePart2(input))
}

func init() {
	Solutions[2] = day2
}
