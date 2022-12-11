package main

import "testing"

var input10Small = []string{
	"noop",
	"addx 3",
	"addx -5",
	"noop",
}

var input10 = []string{
	"addx 15",
	"addx -11",
	"addx 6",
	"addx -3",
	"addx 5",
	"addx -1",
	"addx -8",
	"addx 13",
	"addx 4",
	"noop",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx -35",
	"addx 1",
	"addx 24",
	"addx -19",
	"addx 1",
	"addx 16",
	"addx -11",
	"noop",
	"noop",
	"addx 21",
	"addx -15",
	"noop",
	"noop",
	"addx -3",
	"addx 9",
	"addx 1",
	"addx -3",
	"addx 8",
	"addx 1",
	"addx 5",
	"noop",
	"noop",
	"noop",
	"noop",
	"noop",
	"addx -36",
	"noop",
	"addx 1",
	"addx 7",
	"noop",
	"noop",
	"noop",
	"addx 2",
	"addx 6",
	"noop",
	"noop",
	"noop",
	"noop",
	"noop",
	"addx 1",
	"noop",
	"noop",
	"addx 7",
	"addx 1",
	"noop",
	"addx -13",
	"addx 13",
	"addx 7",
	"noop",
	"addx 1",
	"addx -33",
	"noop",
	"noop",
	"noop",
	"addx 2",
	"noop",
	"noop",
	"noop",
	"addx 8",
	"noop",
	"addx -1",
	"addx 2",
	"addx 1",
	"noop",
	"addx 17",
	"addx -9",
	"addx 1",
	"addx 1",
	"addx -3",
	"addx 11",
	"noop",
	"noop",
	"addx 1",
	"noop",
	"addx 1",
	"noop",
	"noop",
	"addx -13",
	"addx -19",
	"addx 1",
	"addx 3",
	"addx 26",
	"addx -30",
	"addx 12",
	"addx -1",
	"addx 3",
	"addx 1",
	"noop",
	"noop",
	"noop",
	"addx -9",
	"addx 18",
	"addx 1",
	"addx 2",
	"noop",
	"noop",
	"addx 9",
	"noop",
	"noop",
	"noop",
	"addx -1",
	"addx 2",
	"addx -37",
	"addx 1",
	"addx 3",
	"noop",
	"addx 15",
	"addx -21",
	"addx 22",
	"addx -6",
	"addx 1",
	"noop",
	"addx 2",
	"addx 1",
	"noop",
	"addx -10",
	"noop",
	"noop",
	"addx 20",
	"addx 1",
	"addx 2",
	"addx 2",
	"addx -6",
	"addx -11",
	"noop",
	"noop",
	"noop",
}

func TestDay10_1(t *testing.T) {
	xs := parseCrtXs(input10Small)
	expectedXs := []int{1, 1, 1, 4, 4, -1}
	for i := 0; i < len(expectedXs); i++ {
		if xs[i] != expectedXs[i] {
			t.Fatalf("At cycle %d, expected %d but got %d", i, expectedXs[i], xs[i])
		}
	}
}

func TestDay10_2(t *testing.T) {
	xs := parseCrtXs(input10)
	signalStrength := signalStrengthSum(xs)
	if signalStrength != 13140 {
		t.Fatalf("Invalid signal strength, is %d should be 13140", signalStrength)
	}
}

func compareImages(img1 CrtImage, img2 []string) bool {
	if len(img1) != len(img2) {
		return false
	}
	for i := 0; i < len(img1); i++ {
		if string(img1[i][:]) != img2[i] {
			return false
		}
	}
	return true
}

func TestDay10_3(t *testing.T) {
	expectedImage := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	xs := parseCrtXs(input10)
	image := renderCrtImage(xs)
	if !compareImages(image, expectedImage) {
		t.Fatalf("Expected image %v, got %v", expectedImage, image)
	}
}

func BenchmarkDay10_parsing(b *testing.B) {
	input := inputAsString(10)
	for n := 0; n < b.N; n++ {
		parseCrtXs(input)
	}
}

func BenchmarkDay10_part1(b *testing.B) {
	xs := parseCrtXs(inputAsString(10))
	for n := 0; n < b.N; n++ {
		signalStrengthSum(xs)
	}
}

func BenchmarkDay10_part2(b *testing.B) {
	xs := parseCrtXs(inputAsString(10))
	for n := 0; n < b.N; n++ {
		renderCrtImage(xs)
	}
}
