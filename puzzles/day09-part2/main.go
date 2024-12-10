package main

import (
	"errors"
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
	newDotFormat := dotFormat[:]
	fileLocationsMap := getFileLocationsMap(dotFormat)
	highestFileId := len(fileLocationsMap) - 1

	for i := highestFileId; i >= 0; i-- {
		/*
			To optimise, could probably keep a record of the lowest index of each space
			length to save having to start iterating from index 0 each time.
		*/
		lowestValidIndex, err := getLowestValidSpaceIndex(newDotFormat, fileLocationsMap[i])
		if err != nil {
			continue
		}
		if lowestValidIndex < fileLocationsMap[i].startingIndex {
			newDotFormat = utils.SwapChunk(
				lowestValidIndex,
				fileLocationsMap[i].startingIndex,
				fileLocationsMap[i].length,
				newDotFormat,
			)
		}
	}

	return newDotFormat
}

func getLowestValidSpaceIndex(dotFormat []string, block Block) (int, error) {
	spaceToFind := block.length
	cursorIndex := 0
	for cursorIndex < block.startingIndex {
		if dotFormat[cursorIndex] != "." {
			cursorIndex++
			continue
		}
		dotLength := 1
		secondCursor := cursorIndex + 1
		for {
			if dotFormat[secondCursor] == "." {
				dotLength++
				secondCursor++
				continue
			}
			break
		}

		if dotLength >= spaceToFind {
			return cursorIndex, nil
		}
		cursorIndex = secondCursor
	}

	return cursorIndex, errors.New("no space")
}

type Block struct {
	startingIndex int
	length        int
}

func getFileLocationsMap(dotFormat []string) map[int]Block {
	fileLocationsMap := map[int]Block{}
	highestFileId := utils.StringToInt(dotFormat[len(dotFormat)-1])
	// Iterate through fileIdNumbers to initilise
	for i := range highestFileId + 1 {
		fileLocationsMap[i] = Block{
			startingIndex: len(dotFormat),
			length:        0,
		}
	}

	// Interate through dotFormat to correct startingIndexes and lengths
	for i := len(dotFormat) - 1; i >= 0; i-- {
		if dotFormat[i] == "." {
			continue
		}
		fileId := utils.StringToInt(dotFormat[i])
		existingEntry := fileLocationsMap[fileId]
		fileLocationsMap[fileId] = Block{
			startingIndex: min(existingEntry.startingIndex, i),
			length:        existingEntry.length + 1,
		}
	}

	return fileLocationsMap
}

func getChecksum(compacted []string) int {
	total := 0
	for i, block := range compacted {
		if block == "." {
			continue
		}
		total += i * utils.StringToInt(string(block))
	}

	return total
}
