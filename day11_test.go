package main

import "testing"

var input11 = []string{
	"Monkey 0:",
	"  Starting items: 79, 98",
	"  Operation: new = old * 19",
	"  Test: divisible by 23",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 3",
	"",
	"Monkey 1:",
	"  Starting items: 54, 65, 75, 74",
	"  Operation: new = old + 6",
	"  Test: divisible by 19",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 0",
	"",
	"Monkey 2:",
	"  Starting items: 79, 60, 97",
	"  Operation: new = old * old",
	"  Test: divisible by 13",
	"    If true: throw to monkey 1",
	"    If false: throw to monkey 3",
	"",
	"Monkey 3:",
	"  Starting items: 74",
	"  Operation: new = old + 3",
	"  Test: divisible by 17",
	"    If true: throw to monkey 0",
	"    If false: throw to monkey 1",
}

func TestDay11_1(t *testing.T) {
	monkeys := parseMonkeys(input11)
	if len(monkeys) != 4 {
		t.Fatalf("Expected 4 monkeys, got %d", len(monkeys))
	}
}

func TestDay11_2(t *testing.T) {
	monkey := parseMonkeys(input11)[0]
	if monkey.divisor != 23 {
		t.Fatal("Failed to parse monkey divisor")
	}
	if len(monkey.items) != 2 || monkey.items[0] != 79 || monkey.items[1] != 98 {
		t.Fatal("Failed to parse monkey items")
	}
	if monkey.recipients[0] != 2 || monkey.recipients[1] != 3 {
		t.Fatal("Failed to parse monkey items")
	}
	if monkey.inspect(10) != 19*10 {
		t.Fatal("Failed to parse monkey inspect")
	}
}

func TestDay11_3(t *testing.T) {
	monkeys := parseMonkeys(input11)
	monkeys.doRound(3)
	monkey0 := monkeys[0]
	if len(monkey0.items) != 4 {
		t.Fatalf("Expected 4 items after one round, got %d", len(monkey0.items))
	}
	if monkey0.items[0] != 20 || monkey0.items[1] != 23 || monkey0.items[2] != 27 || monkey0.items[3] != 26 {
		t.Fatal("Invalid worry levels")
	}
}

func TestDay11_4(t *testing.T) {
	monkeys := parseMonkeys(input11)
	mb := monkeyBusinessAfter(monkeys, 20, 3)
	if mb != 10605 {
		t.Fatalf("Monkey business is %d, expected 10605", mb)
	}
}

func TestDay11_5(t *testing.T) {
	monkeys := parseMonkeys(input11)
	monkeys.doNRounds(1, reliefValue(monkeys))
	if monkeys[0].inspections != 2 || monkeys[1].inspections != 4 || monkeys[2].inspections != 3 || monkeys[3].inspections != 6 {
		t.Fatalf("Invalid number of inspections after 1 rounds, should be [2 4 3 6] is %v", inspections(monkeys))
	}
}

func TestDay11_6(t *testing.T) {
	monkeys := parseMonkeys(input11)
	monkeys.doNRounds(20, reliefValue(monkeys))
	if monkeys[0].inspections != 99 || monkeys[1].inspections != 97 || monkeys[2].inspections != 8 || monkeys[3].inspections != 103 {
		t.Fatalf("Invalid number of inspections after 20 rounds, should be [99 97 8 103] is %v", inspections(monkeys))
	}
}

func TestDay11_7(t *testing.T) {
	monkeys := parseMonkeys(input11)
	monkeys.doNRounds(10000, reliefValue(monkeys))
	if monkeys[0].inspections != 52166 || monkeys[1].inspections != 47830 || monkeys[2].inspections != 1938 || monkeys[3].inspections != 52013 {
		t.Fatalf("Invalid number of inspections after 10000 rounds, should be [52166 47830 1938 52013] is %v", inspections(monkeys))
	}
}

func inspections(monkeys Monkeys) []uint {
	inspections := []uint{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	return inspections
}

func BenchmarkDay11_parsing(b *testing.B) {
	input := inputAsString(11)
	for n := 0; n < b.N; n++ {
		parseMonkeys(input)
	}
}

func BenchmarkDay11_part1(b *testing.B) {
	monkeys := parseMonkeys(inputAsString(11))
	for n := 0; n < b.N; n++ {
		monkeyBusinessAfter(monkeys.clone(), 20, 3)
	}
}

func BenchmarkDay11_part2(b *testing.B) {
	monkeys := parseMonkeys(inputAsString(11))
	for n := 0; n < b.N; n++ {
		monkeyBusinessAfter(monkeys.clone(), 10000, reliefValue(monkeys))
	}
}
