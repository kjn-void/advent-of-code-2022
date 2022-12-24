package main

import "testing"

var input24Small = []string{
	"#.#####",
	"#.....#",
	"#>....#",
	"#.....#",
	"#...v.#",
	"#.....#",
	"#####.#",
}

var input24Small2 = []string{
	"#.#####",
	"#.^...#",
	"#<....#",
	"#####.#",
}

var input24 = []string{
	"#.######",
	"#>>.<^<#",
	"#.<..<<#",
	"#>v.><>#",
	"#<^v^^>#",
	"######.#",
}

func TestDay24_1(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	if valley.Blizzard.isActive(ValleyPos{1, 1}, 0) {
		t.Fatal("Should not be an active blizzard at 1,1")
	}
	if !valley.Blizzard.isActive(ValleyPos{1, 2}, 0) {
		t.Fatal("Should be an active blizzard at 1,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{4, 4}, 0) {
		t.Fatal("Should be an active blizzard at 4,4")
	}
}

func TestDay24_2(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	if valley.Blizzard.isActive(ValleyPos{1, 2}, 1) {
		t.Fatal("Should not be an active blizzard at 1,2")
	}
	if valley.Blizzard.isActive(ValleyPos{4, 4}, 1) {
		t.Fatal("Should not be an active blizzard at 4,4")
	}
	if !valley.Blizzard.isActive(ValleyPos{2, 2}, 1) {
		t.Fatal("Should be an active blizzard at 2,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{4, 5}, 1) {
		t.Fatal("Should be an active blizzard at 4,5")
	}
}

func TestDay24_3(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	if !valley.Blizzard.isActive(ValleyPos{3, 2}, 2) {
		t.Fatal("Should be an active blizzard at 3,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{4, 1}, 2) {
		t.Fatal("Should be an active blizzard at 4,1")
	}
}

func TestDay24_4(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	for y := 1; y < valley.Height-2; y++ {
		for x := 1; x < valley.Width-2; x++ {
			active := valley.Blizzard.isActive(ValleyPos{x, y}, 3)
			if active {
				if x != 4 || y != 2 {
					t.Fatal("Should only be active at 4,2")
				}
			} else {
				if x == 4 && y == 2 {
					t.Fatal("Should be active at 4,2")
				}
			}
		}
	}
}

func TestDay24_5(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	if !valley.Blizzard.isActive(ValleyPos{5, 2}, 4) {
		t.Fatal("Should be an active blizzard at 5,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{4, 3}, 4) {
		t.Fatal("Should be an active blizzard at 4,2")
	}
}

func TestDay24_6(t *testing.T) {
	valley := parseBlizzardValley(input24Small)
	if !valley.Blizzard.isActive(ValleyPos{1, 2}, 5) {
		t.Fatal("Should be an active blizzard at 1,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{4, 4}, 5) {
		t.Fatal("Should be an active blizzard at 4,4")
	}
}

func TestDay24_7(t *testing.T) {
	valley := parseBlizzardValley(input24Small2)
	if !valley.Blizzard.isActive(ValleyPos{2, 1}, 0) {
		t.Fatal("Should be an active blizzard at 2,1")
	}
	if !valley.Blizzard.isActive(ValleyPos{1, 2}, 0) {
		t.Fatal("Should  be an active blizzard at 1,2")
	}
}

func TestDay24_8(t *testing.T) {
	valley := parseBlizzardValley(input24Small2)
	if !valley.Blizzard.isActive(ValleyPos{2, 2}, 1) {
		t.Fatal("Should be an active blizzard at 2,2")
	}
	if !valley.Blizzard.isActive(ValleyPos{5, 2}, 1) {
		t.Fatal("Should  be an active blizzard at 1,5")
	}
}

func TestDay24_9(t *testing.T) {
	valley := parseBlizzardValley(input24)
	steps := MoveThroughValley(valley)
	if steps != 18 {
		t.Fatalf("Should take 18 steps to move through valley, not %d", steps)
	}
}

func TestDay24_10(t *testing.T) {
	valley := parseBlizzardValley(input24)
	time := MoveThroughValley(valley)
	time = FetchTheSnacks(valley, time)
	if time != 54 {
		t.Fatalf("Should take 18 steps fetch the snacks, not %d", time)
	}
}

func BenchmarkDay24_parsing(b *testing.B) {
	input := inputAsString(24)
	for n := 0; n < b.N; n++ {
		parseBlizzardValley(input)
	}

}

func BenchmarkDay24_part1(b *testing.B) {
	valley := parseBlizzardValley(inputAsString(24))
	for n := 0; n < b.N; n++ {
		MoveThroughValley(valley)
	}
}

func BenchmarkDay24_part2(b *testing.B) {
	valley := parseBlizzardValley(inputAsString(24))
	for n := 0; n < b.N; n++ {
		FetchTheSnacks(valley, MoveThroughValley(valley))
	}
}
