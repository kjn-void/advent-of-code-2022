package main

import "testing"

var input14 = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestDay14_1(t *testing.T) {
	cave := parseCave(input14)
	if cave.FloorDepth != 9 {
		t.Fatalf("Expected floor at depth 9, got %d", cave.FloorDepth)
	}
	tile := cave.get(CavePos{498, 4})
	if tile != TILE_ROCK {
		t.Fatalf("Expected pos 498,4 to be rock(1), is %d", tile)
	}
	tile = cave.get(CavePos{500, 9})
	if tile != TILE_ROCK {
		t.Fatalf("Expected pos 500,9 to be rock(1), is %d", tile)
	}
	tile = cave.get(CavePos{499, 1})
	if tile != TILE_AIR {
		t.Fatalf("Expected pos 499,1 to be air(0), is %d", tile)
	}
}

func TestDay14_2(t *testing.T) {
	cave := parseCave(input14)
	unitsOfSand := cave.UnitsOfSandToRest(false)
	if unitsOfSand != 24 {
		t.Fatalf("Expected 24 units of sand to come to rest, got %d", unitsOfSand)
	}
}

func TestDay14_3(t *testing.T) {
	cave := parseCave(input14)
	unitsOfSand := cave.UnitsOfSandToRest(true)
	if unitsOfSand != 93 {
		t.Fatalf("Expected 93 units of sand to come to rest, got %d", unitsOfSand)
	}
}

func BenchmarkDay14_parsing(b *testing.B) {
	input := inputAsString(14)
	for n := 0; n < b.N; n++ {
		parseCave(input)
	}
}

func BenchmarkDay14_part1(b *testing.B) {
	cave := parseCave(inputAsString(14))
	for n := 0; n < b.N; n++ {
		caveMap := append([]CaveTile{}, cave.Tiles...)
		cave.UnitsOfSandToRest(false)
		cave.Tiles = caveMap
	}
}

func BenchmarkDay14_part2(b *testing.B) {
	cave := parseCave(inputAsString(14))
	for n := 0; n < b.N; n++ {
		caveMap := append([]CaveTile{}, cave.Tiles...)
		cave.UnitsOfSandToRest(true)
		cave.Tiles = caveMap
	}
}
