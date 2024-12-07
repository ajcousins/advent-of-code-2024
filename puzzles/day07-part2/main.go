package main

import (
	"fmt"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day07.txt"
	rows := utils.GetFileLines(filename)
	sumOfValidRows := 0
	for _, row := range rows {
		args := strings.Split(row, ": ")
		target := utils.StringToInt(args[0])
		values := utils.MapSlice(strings.Split(args[1], " "), utils.StringToInt)

		if isValidRow(target, values) {
			sumOfValidRows += target
		}
	}

	println("Answer:", sumOfValidRows)
}

func isValidRow(target int, values []int) bool {
	combination := 0
	firstPass := true
	numberOfOperators := len(values) - 1
	for {
		bin := utils.TrimmedBaseNNumber(numberOfOperators, combination, 3)

		if utils.BaseNToDecimal(bin, 3) == 0 {
			if firstPass {
				firstPass = false
			} else {
				/*
					This is the second time we've seen zero in ternary,
					which means we've checked all combinations.
				*/
				break
			}
		}

		if evaluateRow(values, bin) == target {
			return true
		}
		combination++
	}

	return false
}

func evaluateRow(values []int, bin string) int {
	operations := strings.Split(bin, "")

	total := 0
	for i, val := range values {
		if i == 0 {
			total = val
			continue
		}
		if operations[i-1] == "0" {
			total += values[i]
		}
		if operations[i-1] == "1" {
			total *= values[i]
		}
		if operations[i-1] == "2" {
			total = utils.StringToInt(fmt.Sprintf("%v%v", total, values[i]))
		}
	}

	return total
}
