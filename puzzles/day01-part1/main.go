package main

import (
	"math"
	"sort"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day01.txt"
	rows := utils.GetFileLines(filename)

	var left, right []int
	for _, row := range rows {
		vals := strings.Split(row, "   ")
		left = append(left, utils.StringToInt(vals[0]))
		right = append(right, utils.StringToInt(vals[1]))
	}
	sort.Ints(left)
	sort.Ints(right)

	count := 0
	for i := range left {
		difference := int(math.Abs(float64(left[i]) - float64(right[i])))
		count += difference
	}

	println("Answer:", count)
}
