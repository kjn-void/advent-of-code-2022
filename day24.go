package main

import "fmt"

type BitArray []byte

type Blizzard struct {
	Eastbound  []BitArray
	Westbound  []BitArray
	Northbound []BitArray
	Southbound []BitArray
}

type ValleyPos struct {
	X, Y int
}

type Valley struct {
	Width, Height int
	Blizzard
}

var valleyStart = ValleyPos{1, 0}

func (ba BitArray) isSet(i int) bool {
	idx := i / 8
	bit := i % 8
	return ba[idx]&(1<<bit) != 0
}

func (v Valley) finish() ValleyPos {
	return ValleyPos{v.Width - 2, v.Height - 1}
}

func (b Blizzard) isActive(pos ValleyPos, time int) bool {
	x := pos.X - 1
	y := pos.Y - 1
	westIdx := (x + time) % len(b.Westbound)
	if b.Westbound[westIdx].isSet(y) {
		return true
	}
	eastIdx := (x + len(b.Eastbound) - time%len(b.Eastbound)) % len(b.Eastbound)
	if b.Eastbound[eastIdx].isSet(y) {
		return true
	}
	northIdx := (y + time) % len(b.Northbound)
	if b.Northbound[northIdx].isSet(x) {
		return true
	}
	southIdx := (y + len(b.Southbound) - time%len(b.Southbound)) % len(b.Southbound)
	return b.Southbound[southIdx].isSet(x)
}

func (v Valley) possibleMoves(pos ValleyPos, time int) []ValleyPos {
	steps := [5]ValleyPos{
		{0, 0},  // Stand still
		{0, 1},  // Move south
		{1, 0},  // Move west
		{0, -1}, // Move north
		{-1, 0}, // Move east
	}
	moves := []ValleyPos{}
	valleyFinish := v.finish()
	xMax, yMax := len(v.Blizzard.Eastbound), len(v.Blizzard.Northbound)
	for _, d := range steps {
		p := ValleyPos{pos.X + d.X, pos.Y + d.Y}
		if (p.X > 0 && p.Y > 0 && p.X <= xMax && p.Y <= yMax && !v.Blizzard.isActive(p, time)) ||
			(p.X == valleyStart.X && p.Y == valleyStart.Y) ||
			(p.X == valleyFinish.X && p.Y == valleyFinish.Y) {
			moves = append(moves, p)
		}
	}
	return moves
}

func moveThroughValleyOnce(valley Valley, start, finish ValleyPos, startTime int) int {
	curPaths := &[]ValleyPos{start}
	nxtPaths := &[]ValleyPos{}
	visited := make([]int, valley.Height*valley.Width)

	for time := startTime; ; time++ {
		for _, pos := range *curPaths {
			for _, nextPos := range valley.possibleMoves(pos, time) {
				if nextPos == finish {
					return time
				}
				offset := nextPos.X + nextPos.Y*valley.Width
				if visited[offset] != time {
					visited[offset] = time
					*nxtPaths = append(*nxtPaths, nextPos)
				}
			}
		}
		curPaths, nxtPaths = nxtPaths, curPaths
		*nxtPaths = (*nxtPaths)[:0]
	}
}

func MoveThroughValley(valley Valley) int {
	return moveThroughValleyOnce(valley, valleyStart, valley.finish(), 1)
}

func FetchTheSnacks(valley Valley, firstRunTook int) int {
	time := moveThroughValleyOnce(valley, valley.finish(), valleyStart, firstRunTook)
	return moveThroughValleyOnce(valley, valleyStart, valley.finish(), time)
}

func day24(input []string) {
	valley := parseBlizzardValley(input)
	steps := MoveThroughValley(valley)
	fmt.Println(steps)
	fmt.Println(FetchTheSnacks(valley, steps))
}

func init() {
	Solutions[24] = day24
}

func parseBlizzardValley(input []string) Valley {
	valley := Valley{
		Width:  len(input[0]),
		Height: len(input),
		Blizzard: Blizzard{
			Eastbound:  []BitArray{},
			Westbound:  []BitArray{},
			Northbound: []BitArray{},
			Southbound: []BitArray{},
		},
	}
	bWidth := valley.Width - 2
	bHeight := valley.Height - 2
	valley.Eastbound = make([]BitArray, bWidth)
	valley.Westbound = make([]BitArray, bWidth)
	valley.Northbound = make([]BitArray, bHeight)
	valley.Southbound = make([]BitArray, bHeight)
	for i := 0; i < len(valley.Eastbound); i++ {
		valley.Eastbound[i] = makeBitArray(bHeight)
		valley.Westbound[i] = makeBitArray(bHeight)
	}
	for i := 0; i < len(valley.Northbound); i++ {
		valley.Northbound[i] = makeBitArray(bWidth)
		valley.Southbound[i] = makeBitArray(bWidth)
	}

	for y, row := range input {
		for x, tile := range row {
			valley.Blizzard.set(x-1, y-1, tile)
		}
	}
	return valley
}

func (b *Blizzard) set(x, y int, dir rune) {
	switch dir {
	case '<':
		b.Westbound[x].set(y)
	case '>':
		b.Eastbound[x].set(y)
	case '^':
		b.Northbound[y].set(x)
	case 'v':
		b.Southbound[y].set(x)
	}
}

func makeBitArray(cap int) BitArray {
	return make([]byte, (cap+7)/8)
}

func (ba *BitArray) set(i int) {
	idx := i / 8
	bit := i % 8
	(*ba)[idx] |= 1 << bit
}
