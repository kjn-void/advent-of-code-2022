package main

import (
	"fmt"
)

const (
	XY_FACE FaceId = iota
	XZ_FACE
	YZ_FACE
)

const (
	BITS_PER_COORD = 5
	POINT_STRIDE   = BITS_PER_COORD * 3
)

type FaceId byte

type Droplet struct {
	X, Y, Z int8
}

type Droplets struct {
	Droplets []Droplet
	// No cube is at 0 and dimensions is one larger than any cube position
	Width  int8
	Height int8
	Depth  int8
}

type Face uint64

func (face *Face) set(vertex Droplet, index int) {
	mask := Face((1 << POINT_STRIDE) - 1)
	mask <<= POINT_STRIDE * index
	*face &= ^mask

	x := Face(vertex.X)
	y := Face(vertex.Y)
	z := Face(vertex.Z)
	pt := x + (y << BITS_PER_COORD) + (z << (2 * BITS_PER_COORD))
	pt <<= POINT_STRIDE * index
	*face |= pt
}

func (cube Droplet) makeFace(whichFace FaceId, near bool) Face {
	dF := int8(1)
	if near {
		dF = int8(0)
	}
	var a, b *int8
	switch whichFace {
	case XY_FACE:
		a = &cube.X
		b = &cube.Y
		cube.Z += dF
	case XZ_FACE:
		a = &cube.X
		b = &cube.Z
		cube.Y += dF
	case YZ_FACE:
		a = &cube.Y
		b = &cube.Z
		cube.X += dF
	}
	face := Face(0)
	for i := 0; i < 4; i++ {
		switch i {
		case 1:
			*b += 1
		case 2:
			*a += 1
		case 3:
			*b -= 1
		}
		face.set(cube, i)
	}
	return face
}

func (cube Droplet) faces() [6]Face {
	return [6]Face{
		cube.makeFace(XY_FACE, true),
		cube.makeFace(XY_FACE, false),
		cube.makeFace(XZ_FACE, true),
		cube.makeFace(XZ_FACE, false),
		cube.makeFace(YZ_FACE, true),
		cube.makeFace(YZ_FACE, false),
	}
}

func (world Droplets) tryFillWithWater(cubes *[]Droplet, visited []bool, pos Droplet) {
	if pos.X < 0 || pos.X >= world.Width ||
		pos.Y < 0 || pos.Y >= world.Height ||
		pos.Z < 0 || pos.Z >= world.Depth {
		return
	}

	offset := world.visitedOffset(pos)
	if visited[offset] {
		return
	}
	visited[offset] = true
	*cubes = append(*cubes, pos)

	for _, d := range [6]Droplet{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {-1, 0, 0}, {0, -1, 0}, {0, 0, -1}} {
		newPos := Droplet{pos.X + d.X, pos.Y + d.Y, pos.Z + d.Z}
		world.tryFillWithWater(cubes, visited, newPos)
	}
}

func (world Droplets) visitedOffset(cube Droplet) int {
	x := int(cube.X)
	y := int(cube.Y)
	z := int(cube.Z)
	yOffset := int(world.Width)
	zOffset := yOffset * int(world.Height)
	return x + y*yOffset + z*zOffset
}

func (world Droplets) fillWithWater() []Droplet {
	visited := make([]bool, int(world.Width)*int(world.Height)*int(world.Depth))
	cubes := append([]Droplet{}, world.Droplets...)
	for _, cube := range cubes {
		visited[world.visitedOffset(cube)] = true
	}
	world.tryFillWithWater(&cubes, visited, Droplet{})
	return cubes
}

func (world Droplets) CountExteriorFreeFaces(totalFreeFaces int) int {
	surfaceAndInteriorFree := CountFreeFaces(world.fillWithWater())
	width := int(world.Width)
	height := int(world.Height)
	depth := int(world.Depth)
	surfaceFaces := 2 * (width*height + width*depth + height*depth)
	internalFaces := surfaceAndInteriorFree - surfaceFaces
	return totalFreeFaces - internalFaces
}

func CountFreeFaces(cubes []Droplet) int {
	faces := map[Face]int{}
	for _, cube := range cubes {
		for _, face := range cube.faces() {
			faces[face]++
		}
	}
	numFreeFaces := 0
	for _, cnt := range faces {
		if cnt == 1 {
			numFreeFaces++
		}
	}
	return numFreeFaces
}

func day18(input []string) {
	world := parseCubeWorld(input)
	freeFaces := CountFreeFaces(world.Droplets)
	fmt.Println(freeFaces)
	fmt.Println(world.CountExteriorFreeFaces(freeFaces))
}

func init() {
	Solutions[18] = day18
}

func parseCubeWorld(input []string) Droplets {
	world := Droplets{}
	for _, desc := range input {
		var cube Droplet
		if _, err := fmt.Sscanf(desc, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z); err != nil {
			panic("Failed to parse cube")
		}
		cube.X++
		cube.Y++
		cube.Z++
		world.Droplets = append(world.Droplets, cube)
		if world.Width < cube.X+2 {
			world.Width = cube.X + 2
		}
		if world.Height < cube.Y+2 {
			world.Height = cube.Y + 2
		}
		if world.Depth < cube.Z+2 {
			world.Depth = cube.Z + 2
		}
	}
	return world
}

// Debug...
func (face Face) String() string {
	s := "Face { "
	for i := 0; i < 4; i++ {
		pt := face >> (POINT_STRIDE * i) & ((1 << POINT_STRIDE) - 1)
		x := pt & ((1 << BITS_PER_COORD) - 1)
		y := (pt >> BITS_PER_COORD) & ((1 << BITS_PER_COORD) - 1)
		z := (pt >> (2 * BITS_PER_COORD)) & ((1 << BITS_PER_COORD) - 1)
		s += fmt.Sprintf("{ %d, %d, %d } ", x, y, z)
	}
	s += "}"
	return s
}
