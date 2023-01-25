package main

import "testing"

var input1 = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestDay1_1(t *testing.T) {
	elfs := parseElfTrain(input1)
	mostCalories := elfs.CarryingMostCalories()
	if mostCalories != 24000 {
		t.Fatalf("Expected most calories to be 24000, got %d", mostCalories)
	}
}

func TestDay1_2(t *testing.T) {
	elfs := parseElfTrain(input1)
	topThree := elfs.TopThreeCarrying()
	if topThree != 45000 {
		t.Fatalf("Expected top 3 to carry 45000 calories, got %d", topThree)
	}
}

func BenchmarkDay1_parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseElfTrain(inputAsString(1))
	}
}

func BenchmarkDay1_part1(b *testing.B) {
	elfs := parseElfTrain(inputAsString(1))
	for i := 0; i < b.N; i++ {
		elfs.CarryingMostCalories()
	}
}

func BenchmarkDay1_part2(b *testing.B) {
	elfs := parseElfTrain(inputAsString(1))
	for i := 0; i < b.N; i++ {
		elfs.TopThreeCarrying()
	}
}
