package main

import (
	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day04.txt"
	content := utils.GetFileContents(filename)
	wordsearch := grid.New(content)

	count := 0
	for y := 0; y < wordsearch.Height; y++ {
		for x := 0; x < wordsearch.Width; x++ {
			location := grid.Vector{Y: y, X: x}
			if wordsearch.Get(location) == "A" && isValidMatch(wordsearch, location) {
				count++
			}
		}
	}

	println("Answer:", count)
}

func isValidMatch(wordsearch grid.Grid, location grid.Vector) bool {
	corners := [2]grid.Vector{
		{Y: -1, X: -1}, // top-left
		{Y: -1, X: 1},  // top-right
	}

	for _, corner := range corners {
		cornerAddress := grid.AddVectors(corner, location)
		cornerChar := wordsearch.Get(cornerAddress)

		if !wordsearch.IsWithinBounds(cornerAddress) ||
			(cornerChar != "M" && cornerChar != "S") ||
			!isCorrectOpposite(wordsearch, location, cornerChar, corner) {
			return false
		}
	}

	return true
}

func isCorrectOpposite(wordsearch grid.Grid, currentLocation grid.Vector, char string, corner grid.Vector) bool {
	oppositeCorner := grid.MultiplyVector(corner, -1) // [-1, 1] => [1, -1]
	oppositeLocation := grid.AddVectors(currentLocation, oppositeCorner)
	oppositeChar := wordsearch.Get(oppositeLocation)

	return (char == "M" && oppositeChar == "S") || (char == "S" && oppositeChar == "M")
}
