package main

import (
	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

/*
	There are areas with the same plant that aren't contiguous!
	Contiguous areas with the same plant need to be identified first
	before evaluating. Instead of using the plant type as the key,
	the keys should just be incremented. In which case it could just
	be a slice instead of a map.
*/

type PlantGroup struct {
	cells     []grid.Vector
	area      int
	perimeter int
	// plantType is not unique to a single PlantGroup
	plantType string
}

func main() {
	filename := "../../data/day12.txt"
	content := utils.GetFileContents(filename)
	farm := grid.New(content)
	plantGroups := identifyContiguousAreas(farm)
	totalPrice := getPrice(plantGroups)

	println("Answer:", totalPrice)
}

func identifyContiguousAreas(farm grid.Grid) []PlantGroup {
	visited := map[grid.Vector]string{}
	contiguousAreas := []PlantGroup{}

	for y := 0; y < farm.Height; y++ {
		for x := 0; x < farm.Width; x++ {
			thisAddress := grid.Vector{Y: y, X: x}
			if _, ok := visited[thisAddress]; ok {
				continue
			}
			contiguousArea, newVisited := checkNeighbours(thisAddress, farm, visited)
			contiguousAreas = append(contiguousAreas, contiguousArea)
			visited = newVisited
		}
	}

	return contiguousAreas
}

func checkNeighbours(start grid.Vector, farm grid.Grid, visited map[grid.Vector]string) (PlantGroup, map[grid.Vector]string) {
	targetContent := farm.Get(start)
	queue := []grid.Vector{start}
	dirs := grid.GetDirs()
	plantGroup := PlantGroup{}
	plantGroup.plantType = targetContent

	for len(queue) > 0 {
		currentAddress := queue[0]

		if _, ok := visited[currentAddress]; ok {
			// Already visited.
			queue = queue[1:]
			continue
		}

		visited[currentAddress] = targetContent
		queue = queue[1:]
		plantGroup.cells = append(plantGroup.cells, currentAddress)
		plantGroup.area += 1
		plantGroup.perimeter += getNumberOfEdges(currentAddress, farm)

		// Check neighbouring cells for the same plant
		for _, dir := range dirs {
			neighbourAddress := grid.AddVectors(dir, currentAddress)
			neighbourContent := farm.Get(neighbourAddress)
			if neighbourContent == targetContent {
				queue = append(queue, neighbourAddress)
			}
		}
	}

	return plantGroup, visited
}

func getNumberOfEdges(currentLocation grid.Vector, farm grid.Grid) int {
	dirs := grid.GetDirs()
	count := 0
	thisPlant := farm.Get(currentLocation)

	for _, dir := range dirs {
		neighbourAddress := grid.AddVectors(dir, currentLocation)
		neighbourPlant := farm.Get(neighbourAddress)

		if neighbourPlant != thisPlant {
			count++
		}
	}

	return count
}

func getPrice(plantGroups []PlantGroup) int {
	combinedTotal := 0
	for _, plantType := range plantGroups {
		thisTotal := plantType.area * plantType.perimeter
		combinedTotal += thisTotal
	}

	return combinedTotal
}
