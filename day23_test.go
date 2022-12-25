package main

import (
	"testing"
)

var input23 = []string{
	"..............",
	"..............",
	".......#......",
	".....###.#....",
	"...#...#.#....",
	"....#...##....",
	"...#.###......",
	"...##.#.##....",
	"....#..#......",
	"..............",
	"..............",
	"..............",
}

func TestDay23_2(t *testing.T) {
	elfMap := parseElfMap(input23)
	cnt := EmptyGroundTiles(elfMap, 10)
	if cnt != 110 {
		t.Fatalf("Expected 110 empty ground tiles, got %d", cnt)
	}
}

func TestDay23_3(t *testing.T) {
	elfMap := parseElfMap(input23)
	round := FirstRoundWhereNoElfMoves(elfMap)
	if round != 20 {
		t.Fatalf("Expected 20 to be the first round no elf moves, got %d", round)
	}
}

func BenchmarkDay23_parse(b *testing.B) {
	input := inputAsString(23)
	for n := 0; n < b.N; n++ {
		parseElfMap(input)
	}
}

func BenchmarkDay23_part1(b *testing.B) {
	input := inputAsString(23)
	for n := 0; n < b.N; n++ {
		EmptyGroundTiles(parseElfMap(input), 10)
	}
}

func BenchmarkDay23_part2(b *testing.B) {
	input := inputAsString(23)
	for n := 0; n < b.N; n++ {
		FirstRoundWhereNoElfMoves(parseElfMap(input))
	}
}
