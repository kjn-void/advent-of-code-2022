package main

import (
	"fmt"
	"math"
	"runtime"
)

type Pos struct {
	X, Y int
}

type Dir struct {
	pos Pos
	dir int
}

type Path struct {
	pos   Pos
	steps int
}

type PathQueue []Path

type Heightmap struct {
	Zs     []byte
	Height int
	Width  int
}

func (q *PathQueue) enqueue(path Path) {
	*q = append(*q, path)
}

func (q *PathQueue) dequeue() Path {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (hm Heightmap) getHeight(pos Pos) byte {
	return hm.Zs[pos.Y*hm.Width+pos.X]
}

func (hm *Heightmap) setHeight(pos Pos, z byte) {
	hm.Zs[pos.Y*hm.Width+pos.X] = z
}

func (hm Heightmap) reachableFrom(pos Pos) []Dir {
	reachablePositions := []Dir{}
	h := hm.getHeight(pos)
	for x, d := range [...]Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		adjacentPos := Pos{pos.X + d.X, pos.Y + d.Y}
		if isReachable(h, hm.getHeight(adjacentPos)) {
			reachablePositions = append(reachablePositions, Dir{adjacentPos, x})
		}
	}
	return reachablePositions
}

func (hm Heightmap) StepsToHighestPoint(start, end Pos) int {
	visited := Heightmap{make([]byte, hm.Height*hm.Width), hm.Height, hm.Width}
	visited.setHeight(start, 5)
	posQ := PathQueue{}
	path := Path{start, 0}
	for path.pos != end {
		for _, nextPos := range hm.reachableFrom(path.pos) {
			if visited.getHeight(nextPos.pos) == 0 {
				posQ.enqueue(Path{nextPos.pos, path.steps + 1})
				visited.setHeight(nextPos.pos, byte(nextPos.dir+1))
			}
		}
		if len(posQ) == 0 {
			return math.MaxInt
		}
		path = posQ.dequeue()
	}
	return path.steps
}

func (hm Heightmap) StepsToHighestPointAnyLowPoint(end Pos) int {
	ch := make(chan int, runtime.NumCPU())
	for y := 1; y < hm.Height-1; y++ {
		go func(yy int) {
			shortest := math.MaxInt
			for x := 1; x < hm.Width-1; x++ {
				start := Pos{x, yy}
				if hm.getHeight(start) == 'a' {
					steps := hm.StepsToHighestPoint(start, end)
					if shortest > steps {
						shortest = steps
					}
				}
			}
			ch <- shortest
		}(y)
	}
	shortest := math.MaxInt
	for y := 1; y < hm.Height-1; y++ {
		steps := <-ch
		if steps < shortest {
			shortest = steps
		}
	}
	return shortest
}

func isReachable(from, to byte) bool {
	if to == 0 {
		return false
	}
	if to > from {
		return to-from <= 1
	}
	return true
}

func parseHeightmap(input []string) (Heightmap, Pos, Pos) {
	hm := Heightmap{}
	hm.Height = len(input) + 2
	hm.Width = len(input[0]) + 2
	hm.Zs = make([]byte, hm.Height*hm.Width)
	var start, end Pos
	for r, row := range input {
		for c, height := range []byte(row) {
			pos := Pos{c + 1, r + 1}
			if height == 'S' {
				start = pos
				height = 'a'
			} else if height == 'E' {
				end = pos
				height = 'z'
			}
			hm.setHeight(pos, height)
		}
	}
	return hm, start, end
}

func day12(input []string) {
	hm, start, end := parseHeightmap(input)
	fmt.Println(hm.StepsToHighestPoint(start, end))
	fmt.Println(hm.StepsToHighestPointAnyLowPoint(end))
}

func init() {
	Solutions[12] = day12
}
