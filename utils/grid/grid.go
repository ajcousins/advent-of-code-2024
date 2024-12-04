package grid

import (
	"strings"
)

type Vector struct {
	Y int
	X int
}

type Grid struct {
	src    string
	cells  map[Vector]string
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
		cells:  allCells,
		Width:  width,
		Height: height,
	}
}

func (grid Grid) Get(vector Vector) string {
	return grid.cells[vector]
}

func MultiplyVector(vector Vector, multiplier int) Vector {
	return Vector{Y: vector.Y * multiplier, X: vector.X * multiplier}
}

func AddVectors(vector1, vector2 Vector) Vector {
	return Vector{Y: vector1.Y + vector2.Y, X: vector1.X + vector2.X}
}

func (grid Grid) IsWithinBounds(location Vector) bool {
	return location.Y >= 0 &&
		location.X >= 0 &&
		location.Y < grid.Height &&
		location.X < grid.Width
}
