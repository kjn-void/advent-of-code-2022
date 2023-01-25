package main

import "testing"

var input6 = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestDay6_1(t *testing.T) {
	offsets := [...]int{7, 5, 6, 10, 11}
	for i := 0; i < len(input6); i++ {
		markerOffset := StartOfPacketMarker(input6[i])
		expected := offsets[i]
		if markerOffset != expected {
			t.Fatalf("Got %d, expected %d", markerOffset, expected)
		}
	}
}

func TestDay6_7(t *testing.T) {
	offsets := [...]int{19, 23, 23, 29, 26}
	for i := 0; i < len(input6); i++ {
		markerOffset := startOfMessageMarker(input6[i])
		expected := offsets[i]
		if markerOffset != expected {
			t.Fatalf("Got %d, expected %d", markerOffset, expected)
		}
	}
}

func BenchmarkDay6_parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inputAsString(6)
	}
}

func BenchmarkDay6_part1(b *testing.B) {
	signal := inputAsString(6)[0]
	for i := 0; i < b.N; i++ {
		StartOfPacketMarker(signal)
	}
}

func BenchmarkDay6_part2(b *testing.B) {
	signal := inputAsString(6)[0]
	for i := 0; i < b.N; i++ {
		startOfMessageMarker(signal)
	}
}
