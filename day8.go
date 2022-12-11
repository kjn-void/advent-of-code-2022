package main

import "fmt"

type TreePos struct {
	x, y int
}
type Forest map[TreePos]int

var allDirections = [4]TreePos{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func parseForest(input []string) Forest {
	forest := Forest{}
	for y, row := range input {
		for x, height := range row {
			pos := TreePos{x, y}
			forest[pos] = int(height - '0')
		}
	}
	return forest
}

func (forest Forest) moveAt(pos TreePos, dx, dy int, visitor func(Forest, TreePos) bool) {
	for {
		pos.x += dx
		pos.y += dy
		if !visitor(forest, pos) {
			break
		}
	}
}

func (forest Forest) isVisible(pos TreePos) bool {
	startHeight := forest[pos]
	visible := false
	for i := 0; !visible && i < len(allDirections); i++ {
		d := allDirections[i]
		forest.moveAt(pos, d.x, d.y, func(f Forest, p TreePos) bool {
			if height, inForest := f[p]; !inForest {
				visible = true
				return false
			} else {
				return height < startHeight
			}
		})
	}
	return visible
}

func (forest Forest) numVisibleTrees() int {
	visibleTrees := 0
	for pos := range forest {
		if forest.isVisible(pos) {
			visibleTrees++
		}
	}
	return visibleTrees
}

func (forest Forest) scenicScore(pos TreePos) int {
	startHeight := forest[pos]
	score := 1
	for _, d := range allDirections {
		steps := 0
		forest.moveAt(pos, d.x, d.y, func(f Forest, p TreePos) bool {
			if height, inForest := f[p]; !inForest {
				return false
			} else {
				steps += 1
				return height < startHeight
			}
		})
		score *= steps
	}
	return score
}

func (forest Forest) bestScenicScore() int {
	bestScore := 0
	for pos := range forest {
		score := forest.scenicScore(pos)
		if bestScore < score {
			bestScore = score
		}
	}
	return bestScore
}

func day8(input []string) {
	forest := parseForest(input)
	fmt.Println(forest.numVisibleTrees())
	fmt.Println(forest.bestScenicScore())
}

func init() {
	Solutions[8] = day8
}
