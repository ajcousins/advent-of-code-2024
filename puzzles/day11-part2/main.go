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

type Chunk struct {
	value      int
	multiplier int
}

func main() {
	filename := "../../data/day11.txt"
	content := utils.GetFileContents(filename)
	chunks := blinkAll(content, 75)
	result := evaluateLength(chunks)

	println("Answer:", result)
}

func evaluateLength(chunks []Chunk) int {
	total := 0
	for _, chunk := range chunks {
		total += chunk.multiplier
	}
	return total
}

func blinkAll(content string, repeat int) []Chunk {
	values := utils.MapSlice(strings.Split(content, " "), utils.StringToInt)
	chunks := []Chunk{}

	for _, val := range values {
		newChunk := Chunk{value: val, multiplier: 1}
		chunks = append(chunks, newChunk)
	}

	for i := 0; i < repeat; i++ {
		chunks = blink(chunks)
	}

	return chunks
}

func blink(chunks []Chunk) []Chunk {
	newChunks := []Chunk{}
	for _, chunk := range chunks {
		if chunk.value == 0 {
			newChunks = append(newChunks, Chunk{value: 1, multiplier: chunk.multiplier})
			continue
		}
		if len(strconv.Itoa(chunk.value))%2 == 0 {
			newChunks = append(newChunks, bisectChunkValue(strconv.Itoa(chunk.value), chunk.multiplier)...)
			continue
		}
		newChunks = append(newChunks, Chunk{value: chunk.value * 2024, multiplier: chunk.multiplier})
	}
	newChunks = compressChunks(newChunks)

	return newChunks
}

func compressChunks(chunks []Chunk) []Chunk {
	chunkMap := map[int]int{}

	for _, chunk := range chunks {
		chunkMap[chunk.value] += chunk.multiplier
	}

	newChunks := []Chunk{}
	for key := range chunkMap {
		chunk := Chunk{value: key, multiplier: chunkMap[key]}
		newChunks = append(newChunks, chunk)
	}

	return newChunks
}

func bisectChunkValue(textNumber string, multiplier int) []Chunk {
	midIndex := len(textNumber) / 2
	first := utils.StringToInt(textNumber[:midIndex])
	second := utils.StringToInt(textNumber[midIndex:])

	firstChunk := Chunk{value: first, multiplier: multiplier}
	secondChunk := Chunk{value: second, multiplier: multiplier}

	return []Chunk{firstChunk, secondChunk}
}
