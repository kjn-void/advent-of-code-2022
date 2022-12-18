package main

import "testing"

var input17 = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestDay17_1(t *testing.T) {
	jets := parseJets(input17)
	heightCh := make(chan uint64)
	go func() { DropRocksIntoWell(jets, heightCh) }()
	height := <-heightCh
	if height != 3068 {
		t.Fatalf("Expected a height of 3068, got %d", height)
	}
}

func TestDay17_2(t *testing.T) {
	jets := parseJets(input17)
	heightCh := make(chan uint64)
	go func() { DropRocksIntoWell(jets, heightCh) }()
	<-heightCh
	height := <-heightCh
	if height != 1514285714288 {
		t.Fatalf("Expected a height of 1514285714288, got %d", height)
	}
}

func BenchmarkDay17_parse(b *testing.B) {
	input := inputAsString(17)
	for n := 0; n < b.N; n++ {
		parseJets(input[0])
	}
}

func BenchmarkDay17_part1(b *testing.B) {
	jets := parseJets(inputAsString(17)[0])
	for n := 0; n < b.N; n++ {
		heightCh := make(chan uint64)
		go func() { DropRocksIntoWell(jets, heightCh) }()
		<-heightCh
	}
}

func BenchmarkDay17_part2(b *testing.B) {
	jets := parseJets(inputAsString(17)[0])
	for n := 0; n < b.N; n++ {
		heightCh := make(chan uint64)
		go func() { DropRocksIntoWell(jets, heightCh) }()
		<-heightCh
		<-heightCh
	}
}
