package main

import (
	"fmt"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

const (
	loss = 0
	draw = 3
	win  = 6
)

func tournamentScorePart1(rounds []string) int {
	roundScores := map[string]int{
		// opponent: rock
		"A X": draw + rock,
		"A Y": win + paper,
		"A Z": loss + scissors,
		// opponent: paper
		"B X": loss + rock,
		"B Y": draw + paper,
		"B Z": win + scissors,
		// opponent: scissors
		"C X": win + rock,
		"C Y": loss + paper,
		"C Z": draw + scissors,
	}
	score := 0
	for _, round := range rounds {
		score += roundScores[round]
	}
	return score
}

func tournamentScorePart2(rounds []string) int {
	roundScores := map[string]int{
		// opponent: rock
		"A X": loss + scissors,
		"A Y": draw + rock,
		"A Z": win + paper,
		// opponent: paper
		"B X": loss + rock,
		"B Y": draw + paper,
		"B Z": win + scissors,
		// opponent: scissors
		"C X": loss + paper,
		"C Y": draw + scissors,
		"C Z": win + rock,
	}
	score := 0
	for _, round := range rounds {
		score += roundScores[round]
	}
	return score
}

func day2(input []string) {
	fmt.Println(tournamentScorePart1(input))
	fmt.Println(tournamentScorePart2(input))
}

func init() {
	Solutions[2] = day2
}
