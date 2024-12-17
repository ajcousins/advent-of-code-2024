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
		if val == "O" {
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
	if destVal == "O" {
		warehouse = moveElement(move, destAdd, warehouse)
		warehouse.Swap(elementLocation, destAdd)
		return warehouse
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
	if destVal == "O" {
		return isMoveValid(move, destAdd, warehouse)
	}
	return true
}

func parse(content string) (grid.Grid, string) {
	blocks := strings.Split(content, "\n\n")
	warehouse := grid.New(blocks[0])
	moves := strings.Join(strings.Split(blocks[1], "\n"), "")
	return warehouse, moves
}
