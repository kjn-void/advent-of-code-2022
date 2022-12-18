package main

import "testing"

var input15 = []string{
	"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
	"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
	"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
	"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
	"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
	"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
	"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
	"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
	"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
	"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
	"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
	"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
	"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
	"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
}

func TestDay15_1(t *testing.T) {
	sensorMap := parseSensorMap(input15)
	cnt := sensorMap.CoveredPositions(10)
	if cnt != 26 {
		t.Fatalf("Expected 26 positions to be covered on line 10, got %d", cnt)
	}
}

func TestDay15_2(t *testing.T) {
	sensorMap := parseSensorMap(input15)
	tf := sensorMap.TuningFrequency(20)
	if tf != 56000011 {
		t.Fatalf("Expected tuning frequency to be 56000011, got %d", tf)
	}
}

func BenchmarkDay15_parse(b *testing.B) {
	input := inputAsString(15)
	for n := 0; n < b.N; n++ {
		parseSensorMap(input)
	}
}

func BenchmarkDay15_part1(b *testing.B) {
	sensorMap := parseSensorMap(inputAsString(15))
	for n := 0; n < b.N; n++ {
		sensorMap.CoveredPositions(2000000)
	}
}

func BenchmarkDay15_part2(b *testing.B) {
	sensorMap := parseSensorMap(inputAsString(15))
	for n := 0; n < b.N; n++ {
		sensorMap.TuningFrequency(4000000)
	}
}
