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
		markerOffset := startOfPacketMarker(input6[i])
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

var result int

func BenchmarkDay6_startOfPacketMarker(b *testing.B) {
	signal := inputAsString(6)[0]
	total := 0
	for n := 0; n < b.N; n++ {
		result += startOfPacketMarker(signal)
	}
	result = total
}

func BenchmarkDay6_startOfMessageMarker(b *testing.B) {
	signal := inputAsString(6)[0]
	total := 0
	for n := 0; n < b.N; n++ {
		result += startOfMessageMarker(signal)
	}
	result = total
}
