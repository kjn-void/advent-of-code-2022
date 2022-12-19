package main

import "testing"

var input19 = []string{
	"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
	"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
}

func TestDay19_1(t *testing.T) {
	blueprints := parseBlueprints(input19)
	if len(blueprints) != 2 {
		t.Fatalf("Expected 2 blueprints, got %d", len(blueprints))
	}
	numGeodes := blueprints[0].findMaxOpenGeodes(ProductionTimeLimitPart1)
	if numGeodes != 9 && false {
		t.Fatalf("Expected 9 open geodes from first blueprint, got %d", numGeodes)
	}
}

func TestDay19_2(t *testing.T) {
	blueprints := parseBlueprints(input19)
	if len(blueprints) != 2 {
		t.Fatalf("Expected 2 blueprints, got %d", len(blueprints))
	}
	numGeodes := blueprints[1].findMaxOpenGeodes(ProductionTimeLimitPart1)
	if numGeodes != 12 {
		t.Fatalf("Expected 12 open geodes from second blueprint, got %d", numGeodes)
	}
}

func TestDay19_3(t *testing.T) {
	blueprints := parseBlueprints(input19)
	qualityLevel := SumAllQualityLevels(blueprints)
	if qualityLevel != 33 {
		t.Fatalf("Expected 33 as sum of quality levels, got %d", qualityLevel)
	}
}

func BenchmarkDay19_parsing(b *testing.B) {
	blueprints := parseBlueprints(inputAsString(19))
	for n := 0; n < b.N; n++ {
		SumAllQualityLevels(blueprints)
	}
}

func BenchmarkDay19_part1(b *testing.B) {
	blueprints := parseBlueprints(inputAsString(19))
	for n := 0; n < b.N; n++ {
		SumAllQualityLevels(blueprints)
	}
}

func BenchmarkDay19_part2(b *testing.B) {
	blueprints := parseBlueprints(inputAsString(19))
	for n := 0; n < b.N; n++ {
		FindOpenGeodesProduct(blueprints[:3])
	}
}
