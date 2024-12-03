package main

import (
	"regexp"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day03.txt"
	content := utils.GetFileContents(filename)

	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	cmds := pattern.FindAllString(content, -1)

	total := 0
	for _, cmd := range cmds {
		res := getResult(cmd)
		total += res
	}
	println("Answer:", total)

}

func getResult(input string) int {
	digitPattern := regexp.MustCompile(`\d{1,3}`)
	nums := digitPattern.FindAllString(input, -1)
	ints := utils.MapSlice(nums, utils.StringToInt)

	return ints[0] * ints[1]
}
