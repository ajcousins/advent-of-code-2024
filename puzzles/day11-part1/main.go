package main

import (
	"strconv"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

/*
	Observations:
	- There will be a lot of repeats; Could chunks just
		be given a multiplier to avoid redundancy?
	- The numbers don't need to be in the order given;
		The order shouldn't make any difference to the
		blinks or final count.
*/

func main() {
	filename := "../../data/day11.txt"
	content := utils.GetFileContents(filename)
	length := blinkAll(content, 25)

	println("Answer:", length)
}

func blinkAll(content string, repeat int) int {
	values := utils.MapSlice(strings.Split(content, " "), utils.StringToInt)
	for i := 0; i < repeat; i++ {
		values = blink(values)
	}

	return len(values)
}

func blink(values []int) []int {
	newValues := []int{}
	for _, val := range values {
		if val == 0 {
			newValues = append(newValues, 1)
			continue
		}
		if len(strconv.Itoa(val))%2 == 0 {
			newValues = append(newValues, bisectNumber(strconv.Itoa(val))...)
			continue
		}
		newValues = append(newValues, val*2024)
	}

	return newValues
}

func bisectNumber(textNumber string) []int {
	midIndex := len(textNumber) / 2
	first := utils.StringToInt(textNumber[:midIndex])
	second := utils.StringToInt(textNumber[midIndex:])

	return []int{first, second}
}
