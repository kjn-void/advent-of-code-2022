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
	sum := sumOfDuplicateItems(rs)
	if sum != 157 {
		t.Fatalf("Sum is %d, should be 157", sum)
	}
}

func TestDay3_2(t *testing.T) {
	rs := parseRucksacks(input3)
	sum := sumOfGroupItems(rs)
	if sum != 70 {
		t.Fatalf("Group sum is %d, should be 70", sum)
	}
}
