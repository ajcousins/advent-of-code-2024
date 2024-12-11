package main

import (
	"strconv"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day10.txt"
	content := utils.GetFileContents(filename)
	topographicMap := grid.New(content)
	trailheads := getTrailheads(topographicMap)
	paths := getAllPaths(topographicMap, trailheads)

	println("Answer:", len(paths))
}

func getAllPaths(topoMap grid.Grid, trailheads []grid.Vector) [][]grid.Vector {
	paths := [][]grid.Vector{}
	for _, head := range trailheads {
		initialPath := []grid.Vector{head}
		newPaths := getPaths(0, topoMap, initialPath)
		paths = append(paths, newPaths...)
	}

	return paths
}

func getPaths(currentHeight int, topoMap grid.Grid, path []grid.Vector) [][]grid.Vector {
	if currentHeight == 9 {
		return [][]grid.Vector{path}
	}
	nextHeight := currentHeight + 1
	currentPos := path[len(path)-1]

	dirs := [4]grid.Vector{
		{Y: -1, X: 0}, // north
		{Y: 0, X: 1},  // east
		{Y: 1, X: 0},  // south
		{Y: 0, X: -1}, // west
	}

	allPaths := [][]grid.Vector{}

	for _, dir := range dirs {
		addressToCheck := grid.AddVectors(dir, currentPos)
		if topoMap.Get(addressToCheck) == strconv.Itoa(nextHeight) {
			newPath := make([]grid.Vector, len(path))
			copy(newPath, path)
			newPath = append(newPath, addressToCheck)
			paths := getPaths(nextHeight, topoMap, newPath)
			allPaths = append(allPaths, paths...)
		}
	}

	return allPaths
}

func getTrailheads(topoMap grid.Grid) []grid.Vector {
	trailheads := []grid.Vector{}
	for y := 0; y < topoMap.Height; y++ {
		for x := 0; x < topoMap.Width; x++ {
			here := grid.Vector{Y: y, X: x}
			if topoMap.Get(here) == "0" {
				trailheads = append(trailheads, here)
			}
		}
	}

	return trailheads
}
