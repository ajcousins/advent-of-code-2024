package main

import (
	"math"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day02.txt"
	rows := utils.GetFileLines(filename)

	count := 0
	for _, row := range rows {
		levels := strings.Split(row, " ")
		levelVals := utils.MapSlice(levels, utils.StringToInt)
		if isSafe(levelVals) {
			count++
		} else if itCanBeSafe(levelVals) {
			count++
		}
	}

	println("Answer:", count)
}

func isSafe(levels []int) bool {
	var isIncreasing bool

	for i := 0; i < len(levels)-1; i++ {
		difference := levels[i+1] - levels[i]
		if difference == 0 || math.Abs(float64(difference)) > 3 {
			return false
		}
		if i == 0 {
			isIncreasing = difference > 0
			continue
		}
		if isIncreasing != (difference > 0) {
			return false
		}
	}

	return true
}

func itCanBeSafe(levels []int) bool {
	for i := range levels {
		copy := append([]int{}, levels...)
		newLevels := utils.RemoveNthElement(copy, i)
		if isSafe(newLevels) {
			return true
		}
	}
	return false
}
