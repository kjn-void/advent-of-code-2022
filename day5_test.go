package main

import "testing"

var input5 = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestDay5_1(t *testing.T) {
	cs, _ := parseCrates(input5)
	if len(cs) != 3 {
		t.Fatalf("Wrong crate size")
	}
	if len(cs[0]) != 2 {
		t.Fatalf("Wrong crate size stack 1")
	}
	if len(cs[1]) != 3 {
		t.Fatalf("Wrong crate size stack 2")
	}
	if len(cs[2]) != 1 {
		t.Fatalf("Wrong crate size stack 3")
	}
}

func TestDay5_2(t *testing.T) {
	cs, _ := parseCrates(input5)
	if cs[0].pop() != 'N' || cs[0].pop() != 'Z' {
		t.Fatalf("Wrong crate size")
	}
}

func TestDay5_3(t *testing.T) {
	_, moves := parseCrates(input5)
	if len(moves) != 4 {
		t.Fatalf("Wrong number of moves")
	}
	if moves[0].Cnt != 1 || moves[0].From != 2 || moves[0].To != 1 {
		t.Fatalf("Wrong number of moves")
	}
	if moves[1].Cnt != 3 || moves[1].From != 1 || moves[1].To != 3 {
		t.Fatalf("Wrong number of moves")
	}
}

func TestDay5_4(t *testing.T) {
	cs, moves := parseCrates(input5)
	msg := MoveCrates9000(cs, moves)
	if msg != "CMZ" {
		t.Fatalf("Got message '%s', expected CMZ", msg)
	}
}

func TestDay5_5(t *testing.T) {
	cs, moves := parseCrates(input5)
	msg := MoveCrates9001(cs, moves)
	if msg != "MCD" {
		t.Fatalf("Got message '%s', expected CMZ", msg)
	}
}

func BenchmarkDay5_parsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseCrates(inputAsString(5))
	}
}

func BenchmarkDay5_part1(b *testing.B) {
	storage, moves := parseCrates(inputAsString(5))
	for i := 0; i < b.N; i++ {
		MoveCrates9000(copyCrateStorage(storage), moves)
	}
}

func BenchmarkDay5_part2(b *testing.B) {
	storage, moves := parseCrates(inputAsString(5))
	for i := 0; i < b.N; i++ {
		MoveCrates9001(copyCrateStorage(storage), moves)
	}
}
