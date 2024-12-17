package main

import (
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day15.txt"
	content := utils.GetFileContents(filename)
	warehouse, moves := parse(content)
	finalState := executeMoves(warehouse, moves)
	answer := getGPS(finalState)

	println("Answer:", answer)
}

func getGPS(warehouse grid.Grid) int {
	total := 0
	for address, val := range warehouse.Cells {
		if val == "[" {
			total += address.Y*100 + address.X
		}
	}
	return total
}

func executeMoves(warehouse grid.Grid, moves string) grid.Grid {
	for _, move := range moves {
		robotLocation, _ := warehouse.GetAddressFromValue("@")
		if isMoveValid(move, robotLocation, warehouse) {
			warehouse = moveElement(move, robotLocation, warehouse)
		}
	}
	return warehouse
}

func moveElement(move rune, elementLocation grid.Vector, warehouse grid.Grid) grid.Grid {
	moveMap := grid.MoveMap()
	destAdd := grid.AddVectors(moveMap[string(move)], elementLocation)
	destVal := warehouse.Get(destAdd)
	if move == '>' || move == '<' {
		if destVal == "[" || destVal == "]" {
			warehouse = moveElement(move, destAdd, warehouse)
			warehouse.Swap(elementLocation, destAdd)
			return warehouse
		}
	}
	if move == '^' || move == 'v' {
		if destVal == "[" {
			otherHalfAddress := grid.AddVectors(destAdd, grid.Vector{Y: 0, X: 1})
			warehouse = moveElement(move, destAdd, warehouse)
			warehouse = moveElement(move, otherHalfAddress, warehouse)
			warehouse.Swap(elementLocation, destAdd)
			warehouse.Swap(elementLocation, otherHalfAddress)
			return warehouse
		}
		if destVal == "]" {
			otherHalfAddress := grid.AddVectors(destAdd, grid.Vector{Y: 0, X: -1})
			warehouse = moveElement(move, destAdd, warehouse)
			warehouse = moveElement(move, otherHalfAddress, warehouse)
			warehouse.Swap(elementLocation, destAdd)
			warehouse.Swap(elementLocation, otherHalfAddress)
			return warehouse
		}
	}
	if destVal == "." {
		warehouse.Swap(elementLocation, destAdd)
	}
	return warehouse
}

func isMoveValid(move rune, elementLocation grid.Vector, warehouse grid.Grid) bool {
	moveMap := grid.MoveMap()
	destAdd := grid.AddVectors(moveMap[string(move)], elementLocation)
	destVal := warehouse.Get(destAdd)
	if destVal == "#" {
		return false
	}
	if move == '<' || move == '>' {
		if destVal == "[" || destVal == "]" {
			return isMoveValid(move, destAdd, warehouse)
		}
	}
	if move == '^' || move == 'v' {
		if destVal == "]" {
			otherHalfAddress := grid.AddVectors(destAdd, grid.Vector{Y: 0, X: -1})
			return isMoveValid(move, destAdd, warehouse) && isMoveValid(move, otherHalfAddress, warehouse)
		}
		if destVal == "[" {
			otherHalfAddress := grid.AddVectors(destAdd, grid.Vector{Y: 0, X: 1})
			return isMoveValid(move, destAdd, warehouse) && isMoveValid(move, otherHalfAddress, warehouse)
		}
	}

	return true
}

func parse(content string) (grid.Grid, string) {
	blocks := strings.Split(content, "\n\n")
	warehouse := grid.New(scaleUp(blocks[0]))
	moves := strings.Join(strings.Split(blocks[1], "\n"), "")
	return warehouse, moves
}

func scaleUp(content string) string {
	rows := strings.Split(content, "\n")
	newRows := []string{}
	for _, row := range rows {
		chars := strings.Split(row, "")
		newChars := []string{}
		for _, char := range chars {
			switch char {
			case "#":
				newChars = append(newChars, "##")
			case "O":
				newChars = append(newChars, "[]")
			case ".":
				newChars = append(newChars, "..")
			case "@":
				newChars = append(newChars, "@.")
			}
		}
		newRows = append(newRows, strings.Join(newChars, ""))
	}
	return strings.Join(newRows, "\n")
}
