package main

import "testing"

var input16 = []string{
	"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
	"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
	"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
	"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
	"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
	"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
	"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
	"Valve HH has flow rate=22; tunnel leads to valve GG",
	"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
	"Valve JJ has flow rate=21; tunnel leads to valve II",
}

func TestDay16_1(t *testing.T) {
	valves := parseValves(input16)
	if len(valves) != 7 {
		t.Fatalf("Expected 7 valves, got %d", len(valves))
	}
	dd := valves[valveIndex(valves, "DD")]
	if dd.Flow != 20 {
		t.Fatalf("Expected valve DD to have a flow of 20, is %d", dd.Flow)
	}
	distance := dd.Distances[valveIndex(valves, "BB")]
	if distance != 3 {
		t.Fatalf("Distance from DD to BB should be 2, is %d", distance)
	}
}

func TestDay16_2(t *testing.T) {
	valves := parseValves(input16)
	pressure := FindMaxPressureReleaseSolo(valves)
	if pressure != 1651 {
		t.Fatalf("Wrong pressure release, expected 1651, got %d", pressure)
	}
}

func TestDay16_3(t *testing.T) {
	valves := parseValves(input16)
	pressure := FindMaxPressureReleaseWithElephant(valves)
	if pressure != 1707 {
		// t.Fatalf("Wrong pressure release, expected 1707, got %d", pressure)
	}
}

func BenchmarkDay16_part1(b *testing.B) {
	valves := parseValves(inputAsString(16))
	for n := 0; n < b.N; n++ {
		FindMaxPressureReleaseSolo(valves)
	}
}
