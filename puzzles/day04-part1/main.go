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
			if wordsearch.Get(location) == "X" {
				// For every X, check in every direction for the correct chars.
				matches := getNumberOfMatches(wordsearch, location)
				count += matches
			}
		}
	}

	println("Answer:", count)
}

func getNumberOfMatches(wordsearch grid.Grid, location grid.Vector) int {
	dirs := [8]grid.Vector{
		{Y: -1, X: -1}, {Y: -1, X: 0}, {Y: -1, X: 1},
		{Y: 0, X: -1} /*          */, {Y: 0, X: 1},
		{Y: 1, X: -1}, {Y: 1, X: 0}, {Y: 1, X: 1},
	}

	chars := [3]string{"M", "A", "S"}

	numberOfMatches := 0

	for _, dir := range dirs {
		for i, char := range chars {
			next := grid.AddVectors(grid.MultiplyVector(dir, i+1), location)
			if !wordsearch.IsWithinBounds(next) {
				// Not worth checking the next chars if one of them is already out of bounds.
				break
			}
			if char != wordsearch.Get(next) {
				// Not worth checking the next chars if this char is not expected.
				break
			}
			if i == len(chars)-1 {
				// If last char and still hasn't broken out, then it's a complete match.
				numberOfMatches++
			}
		}
	}

	return numberOfMatches
}
