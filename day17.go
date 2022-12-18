package main

import "fmt"

const (
	CHAMBER_WIDTH         = 7
	ROCK_STOP_LIMIT       = 1_000_000_000_000
	NUMBER_OF_ROCKS_PART1 = 2022
)

type Jets []bool

type RockPos struct {
	X, Y int
}

type RockShape []RockPos

type Well map[RockPos]bool

type ShapeJetIndices struct {
	RockIndex int
	JetIndex  int
}

type RoundHeightPair struct {
	Round  uint64
	Height uint64
}

var shapes = [5]RockShape{
	{{2, 0}, {3, 0}, {4, 0}, {5, 0}},         // - shape
	{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}}, // + shape
	{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}}, // ⅃ shape
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}},         // | shape
	{{2, 0}, {3, 0}, {2, 1}, {3, 1}},         // ▖ shape
}

func (rock RockShape) clone() RockShape {
	return append(RockShape{}, rock...)
}

func (rock RockShape) addDepth(depth int) {
	for i := range rock {
		rock[i].Y += depth
	}

}

func (rock RockShape) tryFall(well Well) bool {
	for _, pt := range rock {
		if well[RockPos{pt.X, pt.Y - 1}] {
			return false
		}
	}
	for i := range rock {
		rock[i].Y--
	}
	return true
}

func (rock RockShape) tryApplyJet(well Well, pushLeft bool) {
	dX := 1
	if pushLeft {
		dX = -1
	}
	for _, pt := range rock {
		newPos := RockPos{pt.X + dX, pt.Y}
		if newPos.X < 0 || newPos.X == CHAMBER_WIDTH || well[newPos] {
			return
		}
	}
	for i := range rock {
		rock[i].X += dX
	}
}

func (well Well) createFloor() {
	for x := 0; x < CHAMBER_WIDTH; x++ {
		well[RockPos{x, 0}] = true
	}
}

func (well Well) addRock(rock RockShape) {
	for _, pt := range rock {
		well[pt] = true
	}
}

func (well Well) dropRock(rock RockShape, jets Jets, jetIndex int) (int, uint64) {
	for {
		jet := jets[jetIndex]
		jetIndex++
		if jetIndex == len(jets) {
			jetIndex = 0
		}

		rock.tryApplyJet(well, jet)
		if !rock.tryFall(well) {
			well.addRock(rock)
			return jetIndex, uint64(rock[len(rock)-1].Y)
		}
	}
}

func (well Well) String() string {
	d := []string{}
	done := false
	for y := 1; !done; y++ {
		s := "|"
		done = true
		for x := 0; x < CHAMBER_WIDTH; x++ {
			if well[RockPos{x, y}] {
				s += "#"
				done = false
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
	well := Well{}
	seen := map[ShapeJetIndices]RoundHeightPair{}
	well.createFloor()

	for n := uint64(0); n <= 3500; n++ {
		// fmt.Println(well)

		rockJetComb := ShapeJetIndices{shapeIndex, jetIndex}
		if p, found := seen[rockJetComb]; found {
			period := n - p.Round
			if n%period == ROCK_STOP_LIMIT%period {
				part2Height = p.Height + (highest+1-p.Height)*(((ROCK_STOP_LIMIT-n)/period)+1) - 1
				if part1Done {
					result <- part2Height
					return
				}
			}
		} else {
			seen[rockJetComb] = RoundHeightPair{n, highest + 1}
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
		rock := shape.clone()
		rock.addDepth(int(highest + 4))
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
