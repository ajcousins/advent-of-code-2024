package main

import (
	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day08.txt"
	content := utils.GetFileContents(filename)
	city := grid.New(content)
	frequencyMap := getFrequencyMap(city)
	antinodeLocations := getAntinodeLocations(frequencyMap, city)

	println("Answer:", len(antinodeLocations))
}

func getFrequencyMap(city grid.Grid) map[string][]grid.Vector {
	frequencyMap := map[string][]grid.Vector{}
	for y := 0; y < city.Height; y++ {
		for x := 0; x < city.Width; x++ {
			thisAddress := grid.Vector{Y: y, X: x}
			freqAtThisAddress := city.Get(thisAddress)
			if freqAtThisAddress == "." {
				continue
			}
			frequencyMap[freqAtThisAddress] = append(frequencyMap[freqAtThisAddress], thisAddress)
		}
	}

	return frequencyMap
}

func getAntinodeLocations(frequencyMap map[string][]grid.Vector, city grid.Grid) []grid.Vector {
	allLocations := []grid.Vector{}
	for freq := range frequencyMap {
		locations := processVectors(frequencyMap[freq], city)
		allLocations = append(allLocations, locations...)
	}
	uniqueLocations := utils.DeduplicateSlice(allLocations)

	return uniqueLocations
}

func processVectors(addresses []grid.Vector, city grid.Grid) []grid.Vector {
	results := []grid.Vector{}
	for _, addr := range addresses {
		antinodes := processVector(addr, addresses, city)
		results = append(results, antinodes...)

	}

	return results
}

func processVector(thisAddress grid.Vector, allAddresses []grid.Vector, city grid.Grid) []grid.Vector {
	results := []grid.Vector{}
	for _, addr := range allAddresses {
		if addr == thisAddress {
			continue
		}
		points := getPoints(addr, thisAddress, city)
		results = append(results, points...)
	}

	return results
}

func getPoints(curAddr, otherAddr grid.Vector, city grid.Grid) []grid.Vector {
	diff := grid.SubtractVectors(curAddr, otherAddr)
	points := []grid.Vector{}
	currentPoint := curAddr
	for {
		if !city.IsWithinBounds(currentPoint) {
			break
		}
		points = append(points, currentPoint)
		currentPoint = grid.AddVectors(currentPoint, diff)
	}

	return points
}
