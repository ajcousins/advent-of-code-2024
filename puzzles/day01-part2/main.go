package main

import (
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

	/*
		For each number in left:
			If left value exists in occurence map:
				Get number of occurences and add to running total
			Else:
				Iterate through right slice and count number of occurences
				Add left value to a map with the number of its occurences
				Add the number of occurences to running total
		Print running total
	*/

	occurences := make(map[int]int)
	total := 0

	for _, leftVal := range left {
		if occurences[leftVal] > 0 {
			total += leftVal * occurences[leftVal]
		} else {
			rightCount := getTotal(leftVal, right)
			occurences[leftVal] = rightCount
			total += leftVal * rightCount
		}
	}

	println("Answer:", total)
}

func getTotal(check int, vals []int) int {
	count := 0
	for _, val := range vals {
		if check == val {
			count++
		}
	}
	return count
}
