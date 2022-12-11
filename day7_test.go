package main

import "testing"

var input7 = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestDay7_1(t *testing.T) {
	fs := parseFilesystem(input7)
	if fs.size != 48381165 {
		t.Fatalf("Filesystem size is %d, expected 48381165", fs.size)
	}
}

func TestDay7_2(t *testing.T) {
	fs := parseFilesystem(input7)
	fsSize := fs.sumSizeWithLimit(100000)
	if fsSize != 95437 {
		t.Fatalf("Filesystem size is %d, expected 95437", fsSize)
	}
}

func TestDay7_4(t *testing.T) {
	fs := parseFilesystem(input7)
	freeUp := fs.bestFitForCap(70000000, 30000000)
	if freeUp != 24933642 {
		t.Fatalf("Selected size is %d, expected 24933642", freeUp)
	}
}

func BenchmarkDay7_parsing(b *testing.B) {
	input := inputAsString(7)
	for n := 0; n < b.N; n++ {
		parseFilesystem(input)
	}
}

func BenchmarkDay7_part1(b *testing.B) {
	fs := parseFilesystem(inputAsString(7))
	for n := 0; n < b.N; n++ {
		fs.sumSizeWithLimit(100000)
	}
}

func BenchmarkDay7_part2(b *testing.B) {
	fs := parseFilesystem(inputAsString(7))
	for n := 0; n < b.N; n++ {
		fs.bestFitForCap(70000000, 30000000)
	}
}
