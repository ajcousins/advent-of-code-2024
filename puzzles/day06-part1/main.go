package main

import (
	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day06.txt"
	content := utils.GetFileContents(filename)
	lab := grid.New(content)
	start, _ := lab.GetAddressFromValue("^")
	count := getNumberOfLocations(lab, start)

	println("Answer:", count)
}

func getNumberOfLocations(lab grid.Grid, start grid.Vector) int {
	dirs := [4]grid.Vector{
		{Y: -1, X: 0}, // north
		{Y: 0, X: 1},  // east
		{Y: 1, X: 0},  // south
		{Y: 0, X: -1}, // west
	}

	guardLocation := start
	directionIndex := 0
	visited := map[grid.Vector]bool{}

	for lab.IsWithinBounds(guardLocation) {
		visited[guardLocation] = true
		currentDirection := dirs[directionIndex%4]
		cellContentInFront := lab.Get(incrementLocation(guardLocation, currentDirection))
		if cellContentInFront == "#" {
			directionIndex++
			continue
		}
		guardLocation = incrementLocation(guardLocation, currentDirection)
	}

	return len(visited)
}

func incrementLocation(location grid.Vector, direction grid.Vector) grid.Vector {
	return grid.AddVectors(location, direction)
}
