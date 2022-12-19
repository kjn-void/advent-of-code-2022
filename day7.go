package main

import "fmt"

type Filesystem map[string]*Directory

type Directory struct {
	Dirs  Filesystem
	Files []File
	Size  int
}

type File struct {
	Size int
}

func makeDir() *Directory {
	return &Directory{Filesystem{}, []File{}, 0}
}

func fillDirSize(d *Directory) {
	dirSz := 0
	for _, f := range d.Files {
		dirSz += f.Size
	}
	for _, d := range d.Dirs {
		fillDirSize(d)
		dirSz += d.Size
	}
	d.Size = dirSz
}

func parseFilesystem(input []string) *Directory {
	rootFs := makeDir()
	i := 1
	parseDir(rootFs, input, &i)
	fillDirSize(rootFs)
	return rootFs
}

func parseDir(d *Directory, input []string, row *int) {
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
			d.Dirs[name] = newDir
			parseDir(newDir, input, row)
		} else if _, err := fmt.Sscanf(cmd, "%d %s", &sz, &name); err == nil {
			(*d).Files = append(d.Files, File{sz})
		} else {
			panic("Failed to parse input")
		}
	}
}

func visitDirs(d Directory, visitor func(Directory)) {
	for _, subDir := range d.Dirs {
		visitDirs(*subDir, visitor)
	}
	visitor(d)
}

func (d Directory) SumSizeWithLimit(maxSize int) int {
	sum := 0
	visitDirs(d, func(d Directory) {
		if d.Size <= maxSize {
			sum += d.Size
		}
	})
	return sum
}

func (d Directory) BestFitForCap(total, required int) int {
	used := d.Size
	free := total - used
	needToFreeUp := required - free
	bestFit := used
	visitDirs(d, func(d Directory) {
		if d.Size >= needToFreeUp && bestFit > d.Size {
			bestFit = d.Size
		}
	})
	return bestFit
}

func day7(input []string) {
	fs := parseFilesystem(input)
	fmt.Println(fs.SumSizeWithLimit(100000))
	fmt.Println(fs.BestFitForCap(70000000, 30000000))
}

func init() {
	Solutions[7] = day7
}
