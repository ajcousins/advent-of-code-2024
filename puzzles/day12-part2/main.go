package main

import (
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

/*
	Number of corners = Number of sides
	Probably easier to figure out the number of corners?
*/

type PlantGroup struct {
	cells     []grid.Vector
	area      int
	perimeter int
	// plantType is not unique to a single PlantGroup
	plantType string
	corners   int
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
	contiguousAreas = getCorners(contiguousAreas)

	return contiguousAreas
}

func getCorners(contiguousAreas []PlantGroup) []PlantGroup {
	newContiguousAreas := []PlantGroup{}
	for _, group := range contiguousAreas {
		amendedGroup := calculateCorners(group)
		newContiguousAreas = append(newContiguousAreas, amendedGroup)
	}

	return newContiguousAreas
}

func calculateCorners(group PlantGroup) PlantGroup {
	cornerDirs := []grid.Vector{
		{Y: -5, X: -5}, {Y: -5, X: 5},
		{Y: 5, X: -5}, {Y: 5, X: 5},
	}

	halfSpaces := map[grid.Vector]int{}

	for _, cell := range group.cells {
		cellAddress := grid.Vector{Y: cell.Y * 10, X: cell.X * 10}
		for _, dir := range cornerDirs {
			cornerAddress := grid.AddVectors(dir, cellAddress)
			halfSpaces[cornerAddress] = getNumberOfCorners(cornerAddress, group.cells)
		}
	}

	for _, t := range halfSpaces {
		group.corners += t
	}

	return group
}

func getNumberOfCorners(halfAddress grid.Vector, groupCells []grid.Vector) int {
	cornerDirs := []grid.Vector{
		{Y: -5, X: -5}, {Y: -5, X: 5},
		{Y: 5, X: -5}, {Y: 5, X: 5},
	}
	combination := []string{}

	for _, dir := range cornerDirs {
		adjAddress := grid.AddVectors(dir, halfAddress)
		farmAddress := grid.Vector{Y: adjAddress.Y / 10, X: adjAddress.X / 10}
		if grid.Includes(farmAddress, groupCells) {
			combination = append(combination, "1")
		} else {
			combination = append(combination, "0")
		}
	}

	count := 0
	for _, digit := range combination {
		count += utils.StringToInt(digit)
	}
	if count == 1 || count == 3 {
		/*
			One target element:
			..		X.		..		.X
			.X		..		X.		..

			Three target elements:
			XX		XX		.X		X.
			X.		.X		XX		XX

			Both conditions mean 1 corner.
		*/
		return 1
	}

	combinedString := strings.Join(combination, "")
	if combinedString == "0110" || combinedString == "1001" {
		/*
			Two corners:
			.X		X.
			X.		.X

		*/
		return 2
	}

	/*
		Zero corners:
		..		X.		XX		.X		..		XX
		XX		X.		..		.X		..		XX
	*/
	return 0
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
		thisTotal := plantType.area * plantType.corners
		combinedTotal += thisTotal
	}

	return combinedTotal
}
