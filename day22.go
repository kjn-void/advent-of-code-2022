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

func (f Facing) DirectionVector() Vec2D {
	return [4]Vec2D{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}[f]
}

func (v Vec2D) add(o Vec2D) Vec2D {
	return Vec2D{v.X + o.X, v.Y + o.Y}
}

func (b Board) tile(pos Vec2D) TileId {
	return b.Tiles[pos.X+pos.Y*b.Width]
}

func (b Board) nextTile(from Vec2D, facing Facing) (TileId, Vec2D) {
	dir := facing.DirectionVector()
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

func day22(input []string) {
	board, me := parseBoard(input)
	fmt.Println(FinalPassword(board, me))
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
