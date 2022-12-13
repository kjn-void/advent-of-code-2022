package main

import (
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type WorryLevel uint64
type Items []WorryLevel
type WorryLevelAdjustment func(WorryLevel) WorryLevel
type Monkey struct {
	items       Items
	inspect     WorryLevelAdjustment
	divisor     WorryLevel
	recipients  [2]int
	inspections uint
}
type Monkeys []Monkey

// Item receiver functions

func (items Items) empty() bool {
	return len(items) == 0
}

func (items *Items) enqueue(item WorryLevel) {
	*items = append(*items, item)
}

func (items *Items) dequeue() WorryLevel {
	first := (*items)[0]
	*items = (*items)[1:]
	return first
}

// Monkeys receiver functions

func (monkeys *Monkeys) doRound(reliefValue WorryLevel) {
	for m := 0; m < len(*monkeys); m++ {
		monkey := &(*monkeys)[m]
		for !monkey.items.empty() {
			newWorryLevel := monkey.inspect(monkey.items.dequeue())
			if reliefValue == 3 {
				newWorryLevel /= reliefValue
			} else {
				newWorryLevel %= reliefValue
			}
			recipient := monkey.recipients[1]
			if newWorryLevel%monkey.divisor == 0 {
				recipient = monkey.recipients[0]
			}
			(*monkeys)[recipient].items.enqueue(newWorryLevel)
			monkey.inspections++
		}
	}
}

func (monkeys *Monkeys) doNRounds(n int, reliefDivisor WorryLevel) {
	for round := 0; round < n; round++ {
		monkeys.doRound(reliefDivisor)
	}
}

func (monkeys Monkeys) clone() Monkeys {
	clones := Monkeys{}
	for _, m := range monkeys {
		clone := m
		m.items = append(Items{}, m.items...)
		clones = append(clones, clone)
	}
	return clones
}

// Functions

func monkeyBusinessAfter(monkeys Monkeys, rounds int, reliefDivisor WorryLevel) WorryLevel {
	monkeys.doNRounds(rounds, reliefDivisor)
	return monkeyBusiness(monkeys)
}

func monkeyBusiness(monkeys Monkeys) WorryLevel {
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspections > monkeys[j].inspections })
	return WorryLevel(monkeys[0].inspections) * WorryLevel(monkeys[1].inspections)
}

func reliefValue(monkeys Monkeys) WorryLevel {
	rd := WorryLevel(1)
	for _, monkey := range monkeys {
		rd *= monkey.divisor
	}
	return rd
}

func day11(input []string) {
	monkeys := parseMonkeys(input)
	fmt.Println(monkeyBusinessAfter(monkeys.clone(), 20, 3))
	fmt.Println(monkeyBusinessAfter(monkeys, 10000, reliefValue(monkeys)))
}

func init() {
	Solutions[11] = day11
}

// Parsing of input

func parseWorryLevel(s string) WorryLevel {
	if wl, err := strconv.Atoi(s); err == nil {
		return WorryLevel(wl)
	} else {
		panic("Failed to parse worry level")
	}
}

func parseMonkeyItems(line string) Items {
	items := Items{}
	parts := strings.Split(line, ": ")
	r := csv.NewReader(strings.NewReader(parts[1]))
	r.TrimLeadingSpace = true
	record, err := r.Read()
	if err != nil {
		panic("Failed to parse monkey items")
	}
	for _, worryLevel := range record {
		items.enqueue(parseWorryLevel(worryLevel))
	}
	return items
}

func parseWorryLevelAdjustment(line string) WorryLevelAdjustment {
	parts := strings.Fields(line)
	secondArg := parts[5]
	if secondArg == "old" {
		return func(old WorryLevel) WorryLevel { return old * old }
	}
	arg := parseWorryLevel(secondArg)
	operator := parts[4]
	if operator == "*" {
		return func(old WorryLevel) WorryLevel { return old * arg }
	}
	return func(old WorryLevel) WorryLevel { return old + arg }
}

func parseDivisor(line string) WorryLevel {
	return parseWorryLevel(strings.Fields(line)[3])
}

func parseRecipient(line string) int {
	return int(parseWorryLevel(strings.Fields(line)[5]))
}

func parseMonkey(input []string, row *int) Monkey {
	monkey := Monkey{}
	monkey.items = parseMonkeyItems(input[*row+1])
	monkey.inspect = parseWorryLevelAdjustment(input[*row+2])
	monkey.divisor = parseDivisor(input[*row+3])
	monkey.recipients[0] = parseRecipient(input[*row+4])
	monkey.recipients[1] = parseRecipient(input[*row+5])
	*row += 7
	return monkey
}

func parseMonkeys(input []string) Monkeys {
	monkeys := Monkeys{}
	row := 0
	for row < len(input) {
		monkeys = append(monkeys, parseMonkey(input, &row))
	}
	return monkeys
}
