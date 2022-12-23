package main

import (
	"fmt"
	"math"
)

type ElfPos struct {
	X, Y int
}

type ElfMap map[ElfPos]bool

type ElfPurposedMap map[ElfPos]*[]ElfPos

func purposedMove(elfMap ElfMap, pos ElfPos, round int, didMove *bool) ElfPos {
	adjacent := [4][3]ElfPos{
		{{-1, -1}, {1, -1}, {0, -1}}, // North
		{{-1, 1}, {1, 1}, {0, 1}},    // South
		{{-1, -1}, {-1, 1}, {-1, 0}}, // West
		{{1, -1}, {1, 1}, {1, 0}},    // East
	}
	purposedPos := pos
	usePurposedPos := false
	for i := 0; i < 4; i++ {
		for j, d := range adjacent[(i+round)%len(adjacent)] {
			checkPos := ElfPos{pos.X + d.X, pos.Y + d.Y}
			if elfMap[checkPos] {
				usePurposedPos = true
				break
			}
			if j == 2 && purposedPos == pos {
				purposedPos = checkPos
			}
		}
	}
	if usePurposedPos && purposedPos != pos {
		*didMove = true
		return purposedPos
	}
	return pos
}

func boundingBox(elfMap ElfMap) (int, int, int, int) {
	xMin, yMin := math.MaxInt, math.MaxInt
	xMax, yMax := math.MinInt, math.MinInt
	for pos := range elfMap {
		if xMin > pos.X {
			xMin = pos.X
		}
		if yMin > pos.Y {
			yMin = pos.Y
		}
		if xMax < pos.X {
			xMax = pos.X
		}
		if yMax < pos.Y {
			yMax = pos.Y
		}
	}
	return xMin, yMin, xMax, yMax
}

func numberOfEmptyGroundTiles(elfMap ElfMap) int {
	xMin, yMin, xMax, yMax := boundingBox(elfMap)
	return (xMax-xMin+1)*(yMax-yMin+1) - len(elfMap)
}

func moveElfs(elfMap ElfMap, round int) (ElfMap, bool) {
	purposed := ElfPurposedMap{}
	atLeastOneElfMoved := false
	// First half
	for pos := range elfMap {
		tryPos := purposedMove(elfMap, pos, round, &atLeastOneElfMoved)
		if collision, found := purposed[tryPos]; found {
			*collision = append(*collision, pos)
		} else {
			purposed[tryPos] = &[]ElfPos{pos}
		}
	}
	// Second half
	elfMap = ElfMap{}
	for purposedPos, oldPosistions := range purposed {
		if len(*oldPosistions) == 1 {
			// Can move
			elfMap[purposedPos] = true
		} else {
			// Collision, keep old position
			for _, oldPos := range *oldPosistions {
				elfMap[oldPos] = true
			}
		}
	}
	return elfMap, atLeastOneElfMoved
}

func EmptyGroundTiles(elfMap ElfMap, rounds int) int {
	for r := 0; r < rounds; r++ {
		elfMap, _ = moveElfs(elfMap, r)
	}
	return numberOfEmptyGroundTiles(elfMap)
}

func FirstRoundWhereNoElfMoves(elfMap ElfMap) int {
	for round := 1; ; round++ {
		var didMove bool
		elfMap, didMove = moveElfs(elfMap, round-1)
		if !didMove {
			return round
		}
	}
}

func day23(input []string) {
	elfMap := parseElfMap(input)
	fmt.Println(EmptyGroundTiles(elfMap, 10))
	fmt.Println(FirstRoundWhereNoElfMoves(elfMap))
}

func init() {
	Solutions[23] = day23
}

func parseElfMap(input []string) ElfMap {
	elfMap := ElfMap{}
	for y, row := range input {
		for x, tile := range row {
			if tile == '#' {
				elfMap[ElfPos{x, y}] = true
			}
		}
	}
	return elfMap
}

// Debug
func (em ElfMap) String() string {
	xMin, yMin, xMax, yMax := boundingBox(em)
	s := ""
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if em[ElfPos{x, y}] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}
