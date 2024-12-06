package main

import (
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day05.txt"
	content := utils.GetFileContents(filename)
	sections := strings.Split(content, "\n\n")
	ruleRows := strings.Split(sections[0], "\n")
	updates := strings.Split(sections[1], "\n")
	rulesMap := generateRulesMap(ruleRows)
	count := getValidUpdateCount(updates, rulesMap)

	println("Answer:", count)
}

func generateRulesMap(rows []string) map[string][]string {
	rulesMap := map[string][]string{}
	for _, row := range rows {
		parts := strings.Split(row, "|")
		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])
	}

	return rulesMap
}

func getValidUpdateCount(updates []string, rulesMap map[string][]string) int {
	count := 0
	for _, update := range updates {
		if !isValidUpdate(update, rulesMap) {
			sorted := sortUpdate(update, rulesMap)
			count += getMiddleValue(sorted)
		}
	}

	return count
}

func isValidUpdate(update string, rulesMap map[string][]string) bool {
	values := strings.Split(update, ",")
	for i := 0; i < len(values); i++ {
		subSection := values[:i]
		if utils.AnyInCommon(subSection, rulesMap[values[i]]) {
			return false
		}
	}

	return true
}

func getMiddleValue(update string) int {
	vals := strings.Split(update, ",")
	middleIndex := len(vals) / 2

	return utils.StringToInt(vals[middleIndex])
}

func sortUpdate(update string, rulesMap map[string][]string) string {
	values := strings.Split(update, ",")
	modified := values[:] // copy of slice

	for i := 0; i < len(values); i++ {
		subSection := values[:i]
		if utils.AnyInCommon(subSection, rulesMap[values[i]]) {
			// swap this value with the one before it
			// set i back to 0 to check from the start
			modified = utils.SwapValues(values, i)
			i = 0
		}
	}

	return strings.Join(modified, ",")
}
