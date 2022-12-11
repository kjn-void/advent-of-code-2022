package main

import "fmt"

type filesystem map[string]*directory

type directory struct {
	dirs  filesystem
	files []file
	size  int
}

type file struct {
	size int
}

func makeDir() *directory {
	return &directory{filesystem{}, []file{}, 0}
}

func fillDirSize(d *directory) {
	dirSz := 0
	for _, f := range d.files {
		dirSz += f.size
	}
	for _, d := range d.dirs {
		fillDirSize(d)
		dirSz += d.size
	}
	d.size = dirSz
}

func parseFilesystem(input []string) *directory {
	rootFs := makeDir()
	i := 1
	parseDir(rootFs, input, &i)
	fillDirSize(rootFs)
	return rootFs
}

func parseDir(d *directory, input []string, row *int) {
	var sz int
	var name string
	for *row < len(input) {
		cmd := input[*row]
		*row++

		if cmd == "$ ls" || cmd[:3] == "dir" {
			continue
		}
		if cmd == "$ cd .." {
			return
		}

		if _, err := fmt.Sscanf(cmd, "$ cd %s", &name); err == nil {
			newDir := makeDir()
			d.dirs[name] = newDir
			parseDir(newDir, input, row)
		} else if _, err := fmt.Sscanf(cmd, "%d %s", &sz, &name); err == nil {
			(*d).files = append(d.files, file{sz})
		} else {
			panic("Failed to parse input")
		}
	}
}

func visitDirs(d directory, visitor func(directory)) {
	for _, subDir := range d.dirs {
		visitDirs(*subDir, visitor)
	}
	visitor(d)
}

func (d directory) sumSizeWithLimit(maxSize int) int {
	sum := 0
	visitDirs(d, func(d directory) {
		if d.size <= maxSize {
			sum += d.size
		}
	})
	return sum
}

func (d directory) bestFitForCap(total, required int) int {
	used := d.size
	free := total - used
	needToFreeUp := required - free
	bestFit := used
	visitDirs(d, func(d directory) {
		if d.size >= needToFreeUp && bestFit > d.size {
			bestFit = d.size
		}
	})
	return bestFit
}

func day7(input []string) {
	fs := parseFilesystem(input)
	fmt.Println(fs.sumSizeWithLimit(100000))
	fmt.Println(fs.bestFitForCap(70000000, 30000000))
}

func init() {
	Solutions[7] = day7
}
