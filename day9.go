package main

import "fmt"

type KnotPos struct {
	X, Y int
}

type KnotMove struct {
	Direction KnotPos
	Steps     int
}

type KnotVisited map[KnotPos]bool

var directions = map[byte]KnotPos{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

func (a KnotPos) adjacent(b KnotPos) bool {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return (dx >= -1 && dx <= 1) && (dy >= -1 && dy <= 1)
}

func (a *KnotPos) add(b KnotPos) {
	a.X += b.X
	a.Y += b.Y
}

func dStep(d int) int {
	switch {
	case d > 0:
		return 1
	case d < 0:
		return -1
	}
	return 0
}

func (a KnotPos) delta(b KnotPos) KnotPos {
	if a.adjacent(b) {
		return KnotPos{}
	}
	return KnotPos{dStep(b.X - a.X), dStep(b.Y - a.Y)}
}

func updateVisited(visited [2]KnotVisited, pos []KnotPos) {
	visited[0][pos[1]] = true
	visited[1][pos[len(pos)-1]] = true
}

func visitedPositions(moves []KnotMove, numKnots int) [2]KnotVisited {
	visited := [2]KnotVisited{{}, {}}
	pos := make([]KnotPos, numKnots)
	updateVisited(visited, pos)

	for _, move := range moves {
		for step := 0; step < move.Steps; step++ {
			pos[0].add(move.Direction)
			for k := 1; k < numKnots; k++ {
				pos[k].add(pos[k].delta(pos[k-1]))
			}
			updateVisited(visited, pos)
		}
	}

	return visited
}

func parseMoves(input []string) []KnotMove {
	moves := []KnotMove{}
	var dir string
	var steps int
	for _, row := range input {
		if _, err := fmt.Sscanf(row, "%v %d", &dir, &steps); err != nil {
			panic("Failed to parse move")
		}
		moves = append(moves, KnotMove{directions[byte(dir[0])], steps})
	}
	return moves
}

func day9(input []string) {
	moves := parseMoves(input)
	visited := visitedPositions(moves, 10)
	fmt.Println(len(visited[0]))
	fmt.Println(len(visited[1]))
}

func init() {
	Solutions[9] = day9
}
