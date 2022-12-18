package main

import (
	"fmt"
	"strings"
)

const CAVE_WIDTH_ORDER = 10

const (
	TILE_AIR CaveTile = iota
	TILE_ROCK
	TILE_SAND
)

type CaveTile byte

type CavePos struct {
	X, Y int
}

type Cave struct {
	Tiles      []CaveTile
	FloorDepth int
}

func normalize(i int) int {
	switch {
	case i == 0:
		return 0
	case i < 0:
		return -1
	default:
		return 1
	}
}

func add(p1, p2 CavePos) CavePos {
	return CavePos{p1.X + p2.X, p1.Y + p2.Y}
}

func tileOffset(pos CavePos) int {
	return pos.X + pos.Y<<CAVE_WIDTH_ORDER
}

func (cave Cave) get(pos CavePos) CaveTile {
	if pos.Y == cave.FloorDepth+2 {
		return TILE_ROCK
	}
	return cave.Tiles[tileOffset(pos)]
}

func (cave *Cave) set(pos CavePos, tile CaveTile) {
	cave.Tiles[tileOffset(pos)] = tile
}

func (cave *Cave) setPath(p1, p2 CavePos) {
	step := CavePos{normalize(p2.X - p1.X), normalize(p2.Y - p1.Y)}
	pos := p1
	end := add(p2, step)
	for pos != end {
		cave.set(pos, TILE_ROCK)
		pos = add(pos, step)
	}
}

func (cave *Cave) pourSand(isPart2 bool) bool {
	sandPos := CavePos{500, 0}
	for sandPos.Y < cave.FloorDepth || isPart2 {
		cameToRest := true
		newSandPos := CavePos{sandPos.X, sandPos.Y + 1}
		for _, d := range [3]int{0, -1, 2} {
			newSandPos.X += d
			if cave.get(newSandPos) == TILE_AIR {
				cameToRest = false
				sandPos = newSandPos
				break
			}
		}
		if cameToRest {
			cave.set(sandPos, TILE_SAND)
			return sandPos.Y > 0
		}
	}
	return false
}

func (cave Cave) UnitsOfSandToRest(isPart2 bool) int {
	unitsOfSand := 0
	for cave.pourSand(isPart2) {
		unitsOfSand++
	}
	if isPart2 {
		unitsOfSand++
	}
	return unitsOfSand
}

func (cave Cave) print() {
	for y := 0; y < cave.FloorDepth+1; y++ {
		for x := 490; x <= 510; x++ {
			pos := CavePos{x, y}
			switch {
			case cave.get(pos) == TILE_AIR:
				fmt.Print(".")
			case cave.get(pos) == TILE_ROCK:
				fmt.Print("#")
			case cave.get(pos) == TILE_SAND:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func day14(input []string) {
	cave := parseCave(input)
	unitsOfSand := cave.UnitsOfSandToRest(false)
	fmt.Println(unitsOfSand)
	fmt.Println(unitsOfSand + cave.UnitsOfSandToRest(true))
}

func init() {
	Solutions[14] = day14
}

func parseCave(input []string) Cave {
	cave := Cave{}
	rocks := [][]CavePos{}
	for _, line := range input {
		rocks = append(rocks, parseRockPath(line, &cave.FloorDepth))
	}
	cave.Tiles = make([]CaveTile, (cave.FloorDepth+3)<<CAVE_WIDTH_ORDER)
	for _, path := range rocks {
		for i := 0; i < len(path)-1; i++ {
			cave.setPath(path[i], path[i+1])
		}
	}
	return cave
}

func parseRockPath(line string, maxDepth *int) []CavePos {
	path := []CavePos{}
	for _, edge := range strings.Split(line, " -> ") {
		var x, y int
		if _, err := fmt.Sscanf(edge, "%d,%d", &x, &y); err != nil {
			panic("Failed to parse edge")
		}
		path = append(path, CavePos{x, y})
		if *maxDepth < y {
			*maxDepth = y
		}
	}
	return path
}
