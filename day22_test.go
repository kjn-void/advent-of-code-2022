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
	if fp != 6032 {
		t.Fatalf("Expected final password of 6032, got %d", fp)
	}
}
