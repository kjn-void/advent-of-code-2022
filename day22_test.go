package main

import (
	"testing"
)

var input22 = []string{
	"        ...#",
	"        .#..",
	"        #...",
	"        ....",
	"...#.......#",
	"........#...",
	"..#....#....",
	"..........#.",
	"        ...#....",
	"        .....#..",
	"        .#......",
	"        ......#.",
	"",
	"10R5L5R10L4R5L5",
}

func TestDay22_1(t *testing.T) {
	board, me := parseBoard(input22)
	fp := FinalPassword(board, me)
	if fp != 6032 {
		t.Fatalf("Expected final password of 6032, got %d", fp)
	}
}

func TestDay22_2(t *testing.T) {
	board, me := parseBoard(input22)
	fp := FinalPasswordCube(board, 4, me.Pos, me.Actions)
	if fp != 5031 {
		t.Fatalf("Expected final password of 5031, got %d", fp)
	}
}

func BenchmarkDay22_parsing(b *testing.B) {
	input := inputAsString(22)
	for n := 0; n < b.N; n++ {
		parseBoard(input)
	}
}

func BenchmarkDay22_part1(b *testing.B) {
	board, me := parseBoard(inputAsString(22))
	for n := 0; n < b.N; n++ {
		FinalPassword(board, me)
	}
}

func BenchmarkDay22_part2(b *testing.B) {
	board, me := parseBoard(inputAsString(22))
	for n := 0; n < b.N; n++ {
		FinalPasswordCube(board, 50, me.Pos, me.Actions)
	}
}
