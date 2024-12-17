package grid

import (
	"errors"
	"strings"
)

type Vector struct {
	Y int
	X int
}

type Grid struct {
	src    string
	Cells  map[Vector]string
	Width  int
	Height int
}

func New(src string) Grid {
	rows := strings.SplitAfter(src, "\n")
	allCells := map[Vector]string{}
	var height, width int

	for y, row := range rows {
		chars := strings.Split(row, "")
		height = len(rows)

		for x, char := range chars {
			address := Vector{y, x}
			allCells[address] = char
			width = len(chars)
		}
	}

	return Grid{
		src:    src,
		Cells:  allCells,
		Width:  width,
		Height: height,
	}
}

func (grid Grid) Get(vector Vector) string {
	return grid.Cells[vector]
}

func MultiplyVector(vector Vector, multiplier int) Vector {
	return Vector{Y: vector.Y * multiplier, X: vector.X * multiplier}
}

func AddVectors(vector1, vector2 Vector) Vector {
	return Vector{Y: vector1.Y + vector2.Y, X: vector1.X + vector2.X}
}

func SubtractVectors(vector1, vector2 Vector) Vector {
	return Vector{Y: vector1.Y - vector2.Y, X: vector1.X - vector2.X}
}

func (grid Grid) IsWithinBounds(location Vector) bool {
	return location.Y >= 0 &&
		location.X >= 0 &&
		location.Y < grid.Height &&
		location.X < grid.Width
}

func (grid Grid) GetAddressFromValue(val string) (Vector, error) {
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			address := Vector{y, x}
			if grid.Get(address) == val {
				return address, nil
			}
		}
	}
	return Vector{-1, -1}, errors.New("not found")
}

func GetDirs() []Vector {
	return []Vector{
		{Y: -1, X: 0}, // north
		{Y: 0, X: 1},  // east
		{Y: 1, X: 0},  // south
		{Y: 0, X: -1}, // west
	}
}

func MoveMap() map[string]Vector {
	return map[string]Vector{
		"^": {Y: -1, X: 0},
		"<": {Y: 0, X: -1},
		">": {Y: 0, X: 1},
		"v": {Y: 1, X: 0},
	}
}

func Includes(target Vector, cells []Vector) bool {
	for _, cell := range cells {
		if target == cell {
			return true
		}
	}
	return false
}

func (g Grid) Swap(aAddress, bAddress Vector) Grid {
	valA := g.Get(aAddress)
	valB := g.Get(bAddress)
	g.Cells[aAddress] = valB
	g.Cells[bAddress] = valA
	return g
}

func (g Grid) Print() {
	rows := []string{}
	for y := 0; y < g.Height; y++ {
		row := []string{}
		for x := 0; x < g.Width; x++ {
			val := g.Get(Vector{y, x})
			row = append(row, val)
		}
		rows = append(rows, strings.Join(row, ""))
	}
	output := strings.Join(rows, "\n")
	println(output)
}
