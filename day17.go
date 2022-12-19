package main

import "fmt"

const (
	CHAMBER_WIDTH         = 7
	ROCK_STOP_LIMIT       = 1_000_000_000_000
	NUMBER_OF_ROCKS_PART1 = 2022
)

type Jets []bool
type RockLine uint16
type RockShape []RockLine
type Rock struct {
	RockShape
	Depth int
}
type Well []RockLine

type ShapeJetIndices struct {
	RockIndex int
	JetIndex  int
}

type RoundHeightPair struct {
	Round  uint64
	Height uint64
}

var shapes = [5]RockShape{
	{RockLine(0x78)}, // - shape
	{RockLine(0x10), RockLine(0x38), RockLine(0x10)},                 // + shape
	{RockLine(0x38), RockLine(0x20), RockLine(0x20)},                 // ⅃ shape
	{RockLine(0x08), RockLine(0x08), RockLine(0x08), RockLine(0x08)}, // | shape
	{RockLine(0x18), RockLine(0x18)},                                 // ▖ shape
}

func (shape RockShape) clone() RockShape {
	return append([]RockLine{}, shape...)
}

func (rock *Rock) tryFall(well Well) bool {
	for dY, sprite := range rock.RockShape {
		if well[rock.Depth+dY-1]&sprite != 0 {
			return false
		}
	}
	rock.Depth--
	return true
}

func (rock *Rock) tryApplyJet(well Well, pushLeft bool) {
	for dY, sprite := range rock.RockShape {
		if pushLeft {
			sprite >>= 1
		} else {
			sprite <<= 1
		}
		if well[rock.Depth+dY]&sprite != 0 {
			return
		}
	}
	for i := range rock.RockShape {
		if pushLeft {
			rock.RockShape[i] >>= 1
		} else {
			rock.RockShape[i] <<= 1
		}
	}
}

func (well *Well) createFloor() {
	*well = append(*well, RockLine(0x1ff))
}

func (well *Well) grow(highest uint64) {
	for uint64(len(*well)) < highest+8 {
		*well = append(*well, RockLine(0x101))
	}
}

func (well Well) addRock(rock Rock) {
	for dY, sprite := range rock.RockShape {
		well[rock.Depth+dY] |= sprite
	}
}

func (well Well) dropRock(rock Rock, jets Jets, jetIndex int) (int, uint64) {
	for {
		jet := jets[jetIndex]
		jetIndex++
		if jetIndex == len(jets) {
			jetIndex = 0
		}

		rock.tryApplyJet(well, jet)
		if !rock.tryFall(well) {
			well.addRock(rock)
			return jetIndex, uint64(rock.Depth + len(rock.RockShape) - 1)
		}
	}
}

func (well Well) String() string {
	d := []string{}
	for y := 1; y < len(well); y++ {
		s := "|"
		for x := 0; x < CHAMBER_WIDTH; x++ {
			if well[y]&(1<<(x+1)) != 0 {
				s += "#"
			} else {
				s += "."
			}
		}
		d = append(d, s+"|")
	}
	s := ""
	for i := len(d) - 1; i >= 0; i-- {
		s += d[i]
		s += "\n"
	}
	s += "+-------+\n"
	return s
}

func DropRocksIntoWell(jets Jets, result chan uint64) {
	part1Done := false
	part2Height := uint64(0)
	highest := uint64(0)
	shapeIndex := 0
	jetIndex := 0
	seen := map[ShapeJetIndices]RoundHeightPair{}
	well := Well{}
	well.createFloor()

	for n := uint64(0); n <= ROCK_STOP_LIMIT; n++ {
		well.grow(highest)

		// fmt.Println(well)

		rockJetComb := ShapeJetIndices{shapeIndex, jetIndex}
		if p, found := seen[rockJetComb]; found {
			period := n - p.Round
			if n%period == ROCK_STOP_LIMIT%period {
				part2Height = p.Height + (highest-p.Height)*(((ROCK_STOP_LIMIT-n)/period)+1)
				if part1Done {
					result <- part2Height
					return
				}
			}
		} else {
			seen[rockJetComb] = RoundHeightPair{n, highest}
		}

		if n == NUMBER_OF_ROCKS_PART1 {
			result <- highest
			part1Done = true
			if part2Height > 0 {
				result <- part2Height
				return
			}
		}

		shape := shapes[shapeIndex]

		var highestRockPiece uint64
		rock := Rock{shape.clone(), int(highest + 4)}
		// for dY, sprite := range rock.RockShape {
		// 	well[rock.Depth+dY] |= sprite
		// }
		// fmt.Println(well)
		// for dY, sprite := range rock.RockShape {
		// 	well[rock.Depth+dY] &= ^sprite
		// }
		jetIndex, highestRockPiece = well.dropRock(rock, jets, jetIndex)
		if highest < highestRockPiece {
			highest = highestRockPiece
		}

		shapeIndex++
		if shapeIndex == len(shapes) {
			shapeIndex = 0
		}
	}
}

func day17(input []string) {
	jets := parseJets(input[0])
	heightCh := make(chan uint64)
	go func() { DropRocksIntoWell(jets, heightCh) }()
	fmt.Println(<-heightCh)
	fmt.Println(<-heightCh)
}

func init() {
	Solutions[17] = day17
}

func parseJets(desc string) Jets {
	jets := Jets{}
	for _, jet := range desc {
		jets = append(jets, jet == '<')
	}
	return jets
}
