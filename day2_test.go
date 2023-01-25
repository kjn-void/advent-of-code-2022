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

func BenchmarkDay2_parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inputAsString(2)
	}
}

func BenchmarkDay2_part1(b *testing.B) {
	input := inputAsString(2)
	for i := 0; i < b.N; i++ {
		TournamentScorePart1(input)
	}
}

func BenchmarkDay2_part2(b *testing.B) {
	input := inputAsString(2)
	for i := 0; i < b.N; i++ {
		TournamentScorePart2(input)
	}
}
