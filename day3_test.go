package main

import "testing"

var input3 = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestDay3_1(t *testing.T) {
	rs := parseRucksacks(input3)
	sum := SumOfDuplicateItems(rs)
	if sum != 157 {
		t.Fatalf("Sum is %d, should be 157", sum)
	}
}

func TestDay3_2(t *testing.T) {
	rs := parseRucksacks(input3)
	sum := SumOfGroupItems(rs)
	if sum != 70 {
		t.Fatalf("Group sum is %d, should be 70", sum)
	}
}

func BenchmarkDay3_parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseRucksacks(inputAsString(3))
	}
}

func BenchmarkDay3_part1(b *testing.B) {
	rs := parseRucksacks(inputAsString(3))
	for i := 0; i < b.N; i++ {
		SumOfDuplicateItems(rs)
	}
}

func BenchmarkDay3_part2(b *testing.B) {
	rs := parseRucksacks(inputAsString(3))
	for i := 0; i < b.N; i++ {
		SumOfGroupItems(rs)
	}
}
