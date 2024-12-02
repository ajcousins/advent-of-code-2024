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
		if isSafe(levels) {
			count++
		}
	}

	println("Answer:", count)
}

func isSafe(levels []string) bool {
	vals := utils.MapSlice(levels, utils.StringToInt)
	var isIncreasing bool

	for i := 0; i < len(vals)-1; i++ {
		difference := vals[i+1] - vals[i]
		if difference == 0 || math.Abs(float64(difference)) > 3 {
			return false
		}
		if i == 0 {
			isIncreasing = checkIfIsIncreasing(difference)
		} else {
			if isIncreasing != checkIfIsIncreasing(difference) {
				return false
			}
		}
	}

	return true
}

func checkIfIsIncreasing(diff int) bool {
	return diff > 0
}
