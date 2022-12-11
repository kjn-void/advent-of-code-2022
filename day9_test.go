package main

import "testing"

var input9 = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

var input9_larger = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestDay9_1(t *testing.T) {
	moves := parseMoves(input9)
	if moves[0].Steps != 4 || moves[0].Direction.X != 1 || moves[0].Direction.Y != 0 {
		t.Fatal("Error in parsing")
	}
}

func TestDay9_2(t *testing.T) {
	moves := parseMoves(input9)
	visited := visitedPositions(moves, 2)[0]
	if len(visited) != 13 {
		t.Fatalf("Visited is %d, should be 13", len(visited))
	}
}

func TestDay9_3(t *testing.T) {
	moves := parseMoves(input9)
	visited := visitedPositions(moves, 10)[1]
	if len(visited) != 1 {
		t.Fatalf("Visited is %d, should be 1", len(visited))
	}
}

func TestDay9_4(t *testing.T) {
	moves := parseMoves(input9)
	visited := visitedPositions(moves, 10)[1]
	if len(visited) != 1 {
		t.Fatalf("Visited is %d, should be 1", len(visited))
	}
}

func TestDay9_5(t *testing.T) {
	moves := parseMoves(input9_larger)
	visited := visitedPositions(moves, 10)[1]
	if len(visited) != 36 {
		t.Fatalf("Visited is %d, should be 36", len(visited))
	}
}

func BenchmarkDay9_parsing(b *testing.B) {
	input := inputAsString(9)
	for n := 0; n < b.N; n++ {
		parseMoves(input)
	}
}

func BenchmarkDay9_part1(b *testing.B) {
	moves := parseMoves(inputAsString(9))
	for n := 0; n < b.N; n++ {
		visitedPositions(moves, 2)
	}
}

func BenchmarkDay9_part2(b *testing.B) {
	moves := parseMoves(inputAsString(9))
	for n := 0; n < b.N; n++ {
		visitedPositions(moves, 10)
	}
}
