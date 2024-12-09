package main

import (
	"fmt"

	"github.com/ajcousins/advent-of-code-2024/utils"
)

func main() {
	filename := "../../data/day09.txt"
	content := utils.GetFileContents(filename)
	dotFormat := convertToDotFormat(content)
	compacted := compact(dotFormat)
	checkSum := getChecksum(compacted)

	println("Answer:", checkSum)
}

func convertToDotFormat(diskmap string) []string {
	charSlice := []string{}
	for i, char := range diskmap {
		timesToRepeat := utils.StringToInt(string(char))
		var fileId string
		if i%2 == 0 {
			fileId = fmt.Sprintf("%v", i/2)
		} else {
			fileId = "."
		}
		block := utils.RepeatChar(fileId, timesToRepeat)
		charSlice = append(charSlice, block...)
	}

	return charSlice
}

func compact(dotFormat []string) []string {
	reversed := utils.FilterElement(utils.ReverseSlice(dotFormat), ".")
	/*
		What is the length of the final compacted slice?
		We need to know this so that we know when to stop iterating.
	*/
	expectedContiguousLength := getExpectedContiguousLength(dotFormat)
	reversedIndexCursor := 0
	formatted := dotFormat[:]

	for i := 0; i < len(dotFormat); i++ {
		if i >= expectedContiguousLength {
			formatted[i] = "."
			continue
		}
		if formatted[i] == "." {
			formatted[i] = reversed[reversedIndexCursor]
			reversedIndexCursor++
		}
	}

	return formatted
}

func getExpectedContiguousLength(dotFormat []string) int {
	count := 0
	for _, char := range dotFormat {
		if char == "." {
			continue
		}
		count++
	}
	return count
}

func getChecksum(compacted []string) int {
	total := 0
	for i, block := range compacted {
		if block == "." {
			break
		}
		total += i * utils.StringToInt(string(block))
	}
	return total
}
