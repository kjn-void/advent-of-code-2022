// 193140

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	Void TileId = iota
	Open
	Wall
)

const (
	Right Facing = iota
	Down
	Left
	Up
)

type TileId byte

type Facing int8

type Vec2D struct {
	X int
	Y int
}

type Vec3D struct {
	X int
	Y int
	Z int
}

type Face3D struct {
	Pos    Vec3D
	Normal Vec3D
}

type Board struct {
	Width  int
	Height int
	Tiles  []TileId
}

type Action struct {
	IsRotation bool
	Steps      int // -1 is rotate counter-clockwise and 1 is clockwise
}

type Me struct {
	Facing
	Pos     Vec2D
	Actions []Action
}

func (f Facing) directionVector() Vec2D {
	return [4]Vec2D{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}[f]
}

func (a Vec2D) add(b Vec2D) Vec2D {
	return Vec2D{a.X + b.X, a.Y + b.Y}
}

func (a Vec3D) add(b Vec3D) Vec3D {
	return Vec3D{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3D) cross(b Vec3D) Vec3D {
	return Vec3D{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (v Vec3D) scalarMul(i int) Vec3D {
	return Vec3D{v.X * i, v.Y * i, v.Z * i}
}

func (b Board) tile(pos Vec2D) TileId {
	return b.Tiles[pos.X+pos.Y*b.Width]
}

func (b Board) safeTile(pos Vec2D) TileId {
	if pos.X < 0 || pos.Y < 0 || pos.X >= b.Width || pos.Y >= b.Height {
		return Void
	}
	return b.tile(pos)
}

func (b Board) nextTile(from Vec2D, facing Facing) (TileId, Vec2D) {
	dir := facing.directionVector()
	pos := from.add(dir)
	tileId := b.tile(pos)
	if tileId == Void {
		switch facing {
		case Up:
			pos.Y = b.Height - 1
		case Down:
			pos.Y = 0
		case Left:
			pos.X = b.Width - 1
		case Right:
			pos.X = 0
		}
		for tileId == Void {
			pos = pos.add(dir)
			tileId = b.tile(pos)
		}
	}
	return tileId, pos
}

func FinalPassword(board Board, me Me) int {
	for _, action := range me.Actions {
		if action.IsRotation {
			me.Facing += Facing(action.Steps)
			if me.Facing == -1 {
				me.Facing = Up
			} else if me.Facing == 4 {
				me.Facing = Right
			}
		} else {
			pos := me.Pos
			for s := 0; s < action.Steps; s++ {
				tile, nextPos := board.nextTile(pos, me.Facing)
				if tile == Wall {
					break
				}
				pos = nextPos
			}
			me.Pos = pos
		}
	}
	return me.Pos.Y*1000 + me.Pos.X*4 + int(me.Facing)
}

func addFace(cube map[Face3D]Vec2D, sz int, board Board, tl2d Vec2D, tl3d, i, j, n Vec3D) {
	if tile := board.safeTile(tl2d); tile == Void {
		return
	}
	if _, exists := cube[Face3D{tl3d, n}]; exists {
		return
	}
	fmt.Println(tl2d, tl3d, n)
	for y := 0; y < sz; y++ {
		for x := 0; x < sz; x++ {
			p3d := tl3d.add(i.scalarMul(x).add(j.scalarMul(y)))
			cube[Face3D{p3d, n}] = tl2d.add(Vec2D{x, y})
		}
	}
	addFace(cube, sz, board, tl2d.add(Vec2D{-sz, 0}), tl3d.add(i.cross(j).scalarMul(1-sz)), i.cross(j), j, n.cross(j))
	addFace(cube, sz, board, tl2d.add(Vec2D{sz, 0}), tl3d.add(i.scalarMul(sz-1)), i.cross(j.scalarMul(-1)), j, n.cross(j.scalarMul(-1)))
	addFace(cube, sz, board, tl2d.add(Vec2D{0, -sz}), tl3d.add(j.cross(i.scalarMul(-1)).scalarMul(1-sz)), i, j.cross(i.scalarMul(-1)), n.cross(i.scalarMul(-1)))
	addFace(cube, sz, board, tl2d.add(Vec2D{0, sz}), tl3d.add(j.scalarMul(sz-1)), i, j.cross(i), n.cross(i))
}

func foldCube(board Board, sz int, topLeft Vec2D) map[Face3D]Vec2D {
	cube := map[Face3D]Vec2D{}
	addFace(cube, sz, board, topLeft, Vec3D{}, Vec3D{1, 0, 0}, Vec3D{0, 1, 0}, Vec3D{0, 0, 1})
	return cube
}

func FinalPasswordCube(board Board, sz int, topLeft Vec2D, actions []Action) int {
	cube := foldCube(board, sz, topLeft)
	for k, v := range cube {
		fmt.Println(k, v)
	}
	return 0
}

func day22(input []string) {
	board, me := parseBoard(input)
	fmt.Println(FinalPassword(board, me))
	fmt.Println(FinalPasswordCube(board, 50, me.Pos, me.Actions))
}

func init() {
	Solutions[22] = day22
}

func parseBoard(input []string) (Board, Me) {
	board := Board{}
	var actionRow int
	for i, row := range input {
		if len(row) == 0 {
			// Create a frame of Void tiles
			board.Height = i + 2
			board.Width += 2
			actionRow = i + 1
			break
		}
		if board.Width < len(row) {
			board.Width = len(row)
		}
	}
	board.Tiles = make([]TileId, board.Width*board.Height)
	for y := 1; y < actionRow; y++ {
		for j, ch := range input[y-1] {
			x := j + 1
			board.setTile(ch, x, y)
		}
	}

	me := Me{Pos: Vec2D{1, 1}, Actions: parseActions(input[actionRow])}
	for board.tile(me.Pos) != Open {
		me.Pos.X++
	}
	return board, me
}

func (b *Board) setTile(tileCh rune, x, y int) {
	var tileId TileId
	switch tileCh {
	case '.':
		tileId = Open
	case '#':
		tileId = Wall
	case ' ':
		tileId = Void
	default:
		panic(fmt.Sprintf("Invalid tile character: %c", tileCh))
	}
	b.Tiles[x+y*b.Width] = tileId
}

func parseActions(input string) []Action {
	actions := []Action{}
	for start, end, isRotation := 0, 0, false; end < len(input); start, isRotation = end, !isRotation {
		var steps int
		if isRotation {
			end++
			switch input[start:end] {
			case "R":
				steps = 1
			case "L":
				steps = -1
			default:
				panic("Invalid rotation: " + input[start:end])
			}
		} else {
			for end < len(input) && unicode.IsDigit(rune(input[end])) {
				end++
			}
			if n, err := strconv.Atoi(input[start:end]); err == nil {
				steps = n
			} else {
				panic("Cannot parse step length: " + input[start:end])
			}
		}
		actions = append(actions, Action{isRotation, steps})
	}
	return actions
}
