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

type Cube struct {
	X, Y, Z int8
}

type CubeWorld struct {
	cubes []Cube
	// No cube is at 0 and dimensions is one outside any cube
	width  int8
	height int8
	depth  int8
}

type Face uint64

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

func (face *Face) set(pos Cube, index int) {
	mask := Face((1 << POINT_STRIDE) - 1)
	mask <<= POINT_STRIDE * index
	*face &= ^mask

	x := Face(pos.X)
	y := Face(pos.Y)
	z := Face(pos.Z)
	pt := x + (y << BITS_PER_COORD) + (z << (2 * BITS_PER_COORD))
	pt <<= POINT_STRIDE * index
	*face |= pt
}

func (cube Cube) makeFace(whichFace FaceId, near bool) Face {
	dF := int8(1)
	if near {
		dF = int8(0)
	}
	var a, b *int8
	switch {
	case whichFace == XY_FACE:
		a = &cube.X
		b = &cube.Y
		cube.Z += dF
	case whichFace == XZ_FACE:
		a = &cube.X
		b = &cube.Z
		cube.Y += dF
	case whichFace == YZ_FACE:
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

func (cube Cube) faces() [6]Face {
	return [6]Face{
		cube.makeFace(XY_FACE, true),
		cube.makeFace(XY_FACE, false),
		cube.makeFace(XZ_FACE, true),
		cube.makeFace(XZ_FACE, false),
		cube.makeFace(YZ_FACE, true),
		cube.makeFace(YZ_FACE, false),
	}
}

func (world CubeWorld) fillWaterFrom(cubes *[]Cube, visited []bool, pos Cube) {
	if pos.X < 0 || pos.X >= world.width ||
		pos.Y < 0 || pos.Y >= world.height ||
		pos.Z < 0 || pos.Z >= world.depth {
		return
	}

	offset := world.visitedOffset(pos)
	if visited[offset] {
		return
	}

	visited[offset] = true
	*cubes = append(*cubes, pos)

	for _, d := range [6]Cube{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}} {
		world.fillWaterFrom(cubes, visited, Cube{pos.X + d.X, pos.Y + d.Y, pos.Z + d.Z})
	}
}

func (world CubeWorld) visitedOffset(cube Cube) int {
	x := int(cube.X)
	y := int(cube.Y)
	z := int(cube.Z)
	yOffset := int(world.width)
	zOffset := yOffset * int(world.height)
	return x + y*yOffset + z*zOffset
}

func (world CubeWorld) fillWithWater() []Cube {
	cubes := append([]Cube{}, world.cubes...)
	visited := make([]bool, int(world.width)*int(world.height)*int(world.depth))
	for _, cube := range cubes {
		visited[world.visitedOffset(cube)] = true
	}

	world.fillWaterFrom(&cubes, visited, Cube{})

	return cubes
}

func (world CubeWorld) CountExteriorFreeFaces(totalFreeFaces int) int {
	surfaceAndInteriorFree := CountFreeFaces(world.fillWithWater())
	width := int(world.width)
	height := int(world.height)
	depth := int(world.depth)
	surfaceFaces := 2 * (width*height + width*depth + height*depth)
	internalFaces := surfaceAndInteriorFree - surfaceFaces
	return totalFreeFaces - internalFaces
}

func CountFreeFaces(cubes []Cube) int {
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
	freeFaces := CountFreeFaces(world.cubes)
	fmt.Println(freeFaces)
	fmt.Println(world.CountExteriorFreeFaces(freeFaces))
}

func init() {
	Solutions[18] = day18
}

func parseCubeWorld(input []string) CubeWorld {
	world := CubeWorld{}
	for _, desc := range input {
		var cube Cube
		if _, err := fmt.Sscanf(desc, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z); err != nil {
			panic("Failed to parse cube")
		}
		cube.X++
		cube.Y++
		cube.Z++
		world.cubes = append(world.cubes, cube)
		if world.width < cube.X+2 {
			world.width = cube.X + 2
		}
		if world.height < cube.Y+2 {
			world.height = cube.Y + 2
		}
		if world.depth < cube.Z+2 {
			world.depth = cube.Z + 2
		}
	}
	return world
}
