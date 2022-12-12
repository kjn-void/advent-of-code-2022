package main

import "testing"

var input12 = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestDay12_1(t *testing.T) {
	hm, start, end := parseHeightmap(input12)
	steps := hm.StepsToHighestPoint(start, end)
	if steps != 31 {
		t.Fatalf("Got %d steps, expected 31", steps)
	}
}

func BenchmarkDay12_parsing(b *testing.B) {
	input := inputAsString(12)
	for n := 0; n < b.N; n++ {
		parseHeightmap(input)
	}
}

func BenchmarkDay12_part1(b *testing.B) {
	hm, start, end := parseHeightmap(inputAsString(12))
	for n := 0; n < b.N; n++ {
		hm.StepsToHighestPoint(start, end)
	}
}

func BenchmarkDay12_part2(b *testing.B) {
	hm, _, end := parseHeightmap(inputAsString(12))
	for n := 0; n < b.N; n++ {
		hm.StepsToHighestPointAnyLowPoint(end)
	}
}
