package main

import "testing"

var input21 = []string{
	"root: pppw + sjmn",
	"dbpl: 5",
	"cczh: sllz + lgvd",
	"zczc: 2",
	"ptdq: humn - dvpt",
	"dvpt: 3",
	"lfqf: 4",
	"humn: 5",
	"ljgn: 2",
	"sjmn: drzm * dbpl",
	"sllz: 4",
	"pppw: cczh / lfqf",
	"lgvd: ljgn * ptdq",
	"drzm: hmdt - zczc",
	"hmdt: 32",
}

func TestDay21_1(t *testing.T) {
	mm, _ := parseMonkeyMatch(input21)
	rootYelling := RootMonkeyYelling(mm)
	if rootYelling != 152 {
		t.Fatalf("Got root yelling %d, expected 152", rootYelling)
	}
}

func TestDay21_2(t *testing.T) {
	mm, rootArgs := parseMonkeyMatch(input21)
	myNumber := MyNumberForRootMonkeyMatch(mm, rootArgs)
	if myNumber != 301 {
		t.Fatalf("I had to yell %d, expected 301", myNumber)
	}
}

func BenchmarkDay21_parse(b *testing.B) {
	input := inputAsString(21)
	for n := 0; n < b.N; n++ {
		parseMonkeyMatch(input)
	}
}

func BenchmarkDay21_part1(b *testing.B) {
	mm, _ := parseMonkeyMatch(inputAsString(21))
	for n := 0; n < b.N; n++ {
		RootMonkeyYelling(mm)
	}
}

func BenchmarkDay21_part2(b *testing.B) {
	mm, rootArgs := parseMonkeyMatch(inputAsString(21))
	for n := 0; n < b.N; n++ {
		MyNumberForRootMonkeyMatch(mm, rootArgs)
	}
}
