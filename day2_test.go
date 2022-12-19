package main

import "testing"

var input2 = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestDay2_1(t *testing.T) {
	score := TournamentScorePart1(input2)
	if score != 15 {
		t.Fatal("Wrong score")
	}
}

func TestDay2_2(t *testing.T) {
	score := TournamentScorePart2(input2)
	if score != 12 {
		t.Fatal("Wrong score")
	}
}
