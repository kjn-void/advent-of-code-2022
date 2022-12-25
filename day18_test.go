package main

import "testing"

var input18 = []string{
	"2,2,2",
	"1,2,2",
	"3,2,2",
	"2,1,2",
	"2,3,2",
	"2,2,1",
	"2,2,3",
	"2,2,4",
	"2,2,6",
	"1,2,5",
	"3,2,5",
	"2,1,5",
	"2,3,5",
}

func TestDay18_1(t *testing.T) {
	world := parseCubeWorld([]string{"1,1,1", "2,1,1"})
	numFaces := CountFreeFaces(world.Droplets)
	if numFaces != 10 {
		t.Fatalf("Expected 10 free faces, got %d", numFaces)
	}
}

func TestDay18_2(t *testing.T) {
	world := parseCubeWorld(input18)
	numFaces := CountFreeFaces(world.Droplets)
	if numFaces != 64 {
		t.Fatalf("Expected 64 free faces, got %d", numFaces)
	}
}

func TestDay18_3(t *testing.T) {
	world := parseCubeWorld([]string{"0,0,0"})
	numFaces := world.CountExteriorFreeFaces(CountFreeFaces(world.Droplets))
	if numFaces != 6 {
		t.Fatalf("Expected 6 free exterior faces after filling with water, got %d", numFaces)
	}
}

func TestDay18_4(t *testing.T) {
	world := parseCubeWorld(input18)
	numFaces := world.CountExteriorFreeFaces(CountFreeFaces(world.Droplets))
	if numFaces != 58 {
		t.Fatalf("Expected 58 free exterior faces after filling with water, got %d", numFaces)
	}
}

func BenchmarkDay18_parsing(b *testing.B) {
	input := inputAsString(18)
	for n := 0; n < b.N; n++ {
		parseCubeWorld(input)
	}
}

func BenchmarkDay18_part1(b *testing.B) {
	world := parseCubeWorld(inputAsString(18))
	for n := 0; n < b.N; n++ {
		CountFreeFaces(world.Droplets)
	}
}

func BenchmarkDay18_part2(b *testing.B) {
	world := parseCubeWorld(inputAsString(18))
	for n := 0; n < b.N; n++ {
		world.CountExteriorFreeFaces(CountFreeFaces(world.Droplets))
	}
}
