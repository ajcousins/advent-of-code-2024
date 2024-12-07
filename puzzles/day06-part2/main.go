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
	count := 0
	for y := 0; y < lab.Height; y++ {
		for x := 0; x < lab.Width; x++ {
			// Place an obstacle at each location.
			if isThisConfigurationValid(lab, start, grid.Vector{Y: y, X: x}) {
				count++
			}
		}
	}
	return count
}

func isThisConfigurationValid(lab grid.Grid, start grid.Vector, obstacle grid.Vector) bool {
	/*
		Keep track of each location visited and the direction the
		guard is facing. If the guard is at the same location and
		facing the same direction, then assume she is in a loop.
	*/
	dirs := [4]grid.Vector{
		{Y: -1, X: 0}, // north
		{Y: 0, X: 1},  // east
		{Y: 1, X: 0},  // south
		{Y: 0, X: -1}, // west
	}
	guardLocation := start
	directionIndex := 0
	visited := map[grid.Vector]grid.Vector{}
	/*
		There are single-line loops which happen that the below
		check won't catch. For these check the step limit.
		A better way to do this would be to record all facing-
		directions at this location and not just the most recent...
		I've got other stuff to do today, so won't bother for now.
	*/
	steps := 0
	stepLimit := 100_000

	for lab.IsWithinBounds(guardLocation) {
		if steps > stepLimit {
			/*
				If the guard is still walking after this long,
				then she must be in a loop.
			*/
			return true
		}
		currentDirection := dirs[directionIndex%4]
		visited[guardLocation] = currentDirection
		cellAddressInFront := incrementLocation(guardLocation, currentDirection)
		cellContentInFront := lab.Get(cellAddressInFront)
		if cellContentInFront == "#" || cellAddressInFront == obstacle {
			directionIndex++
			continue
		}
		guardLocation = incrementLocation(guardLocation, currentDirection)

		if visited[guardLocation] == currentDirection {
			return true
		}
		steps++
	}

	return false
}

func incrementLocation(location grid.Vector, direction grid.Vector) grid.Vector {
	return grid.AddVectors(location, direction)
}
