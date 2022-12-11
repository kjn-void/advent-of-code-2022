package main

import "testing"

var input8 = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestDay8_1(t *testing.T) {
	forest := parseForest(input8)
	height := forest[TreePos{2, 2}]
	if height != 3 {
		t.Fatalf("Expected height 3, found %d", height)
	}
}

func TestDay8_2(t *testing.T) {
	forest := parseForest(input8)
	visible := forest.isVisible(TreePos{1, 1})
	if !visible {
		t.Fatalf("Top-left should be visible")
	}
}

func TestDay8_3(t *testing.T) {
	forest := parseForest(input8)
	visible := forest.isVisible(TreePos{2, 2})
	if visible {
		t.Fatalf("Center tree should not be visible")
	}
}

func TestDay8_4(t *testing.T) {
	forest := parseForest(input8)
	visible := forest.isVisible(TreePos{0, 0})
	if !visible {
		t.Fatalf("Edge trees should be visible")
	}
}

func TestDay8_5(t *testing.T) {
	forest := parseForest(input8)
	numVisible := forest.numVisibleTrees()
	if numVisible != 21 {
		t.Fatalf("Should be 21 visible trees")
	}
}

func TestDay8_6(t *testing.T) {
	forest := parseForest(input8)
	score := forest.scenicScore(TreePos{2, 1})
	if score != 4 {
		t.Fatalf("Score should be 4, is %d", score)
	}
}

func TestDay8_7(t *testing.T) {
	forest := parseForest(input8)
	score := forest.scenicScore(TreePos{2, 3})
	if score != 8 {
		t.Fatalf("Score should be 8, is %d", score)
	}
}

func TestDay8_8(t *testing.T) {
	forest := parseForest(input8)
	score := forest.bestScenicScore()
	if score != 8 {
		t.Fatalf("Score should be 8, is %d", score)
	}
}

func BenchmarkDay8_parsing(b *testing.B) {
	input := inputAsString(8)
	for n := 0; n < b.N; n++ {
		parseForest(input)
	}
}

func BenchmarkDay8_part1(b *testing.B) {
	forest := parseForest(inputAsString(8))
	for n := 0; n < b.N; n++ {
		forest.numVisibleTrees()
	}
}

func BenchmarkDay8_part2(b *testing.B) {
	forest := parseForest(inputAsString(8))
	for n := 0; n < b.N; n++ {
		forest.bestScenicScore()
	}
}
