package main

import (
	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

/*
	Create a map where the keys are the frequencies and the values are slices
	of vectors. Each vector is an address.

	For each frequency, iterate so that each address is paired with every other
	address of the same frequency.
	Each pair of addresses should give the address of an antinode.
	If an antinode address is out of bounds: disregard it.
	Otherwise store these addresses in a slice of Vectors and deduplicate.

	Answer is the length of this slice.
*/

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
		diff := grid.SubtractVectors(thisAddress, addr)
		result := grid.AddVectors(thisAddress, diff)
		if city.IsWithinBounds(result) {
			results = append(results, result)
		}
	}

	return results
}
