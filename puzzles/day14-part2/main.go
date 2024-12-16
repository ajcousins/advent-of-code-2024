package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	X_BOUNDARY := 101
	Y_BOUNDARY := 103
	filename := "../../data/day14.txt"
	content := utils.GetFileContents(filename)
	robots := initialiseRobots(content, X_BOUNDARY, Y_BOUNDARY)
	iteratePositions(3965, 10000, robots, X_BOUNDARY, Y_BOUNDARY)
}

func iteratePositions(start int, end int, robots []Robot, xBound, yBound int) {
	for i := start; i < end; i += 103 {
		positions := getPositionsAt(robots, i)
		printPos(positions, i, xBound, yBound)
		time.Sleep(time.Millisecond * 100)
	}
}

func printPos(positions []grid.Vector, index, xBound, yBound int) {
	robMap := map[grid.Vector]bool{}
	for _, pos := range positions {
		robMap[pos] = true
	}

	fmt.Printf("_________________________________ %v _________________________________\n", index)

	printString := []string{}

	for y := 0; y < yBound; y++ {
		row := []string{}
		for x := 0; x < xBound; x++ {
			here := grid.Vector{Y: y, X: x}
			if robMap[here] {
				row = append(row, "#")
			} else {
				row = append(row, " ")
			}
		}
		row = append(row, "\n")
		joinedRow := strings.Join(row, "")
		printString = append(printString, joinedRow)
	}
	fmt.Println(printString)
}

type Robot struct {
	startingPos    grid.Vector
	velocity       grid.Vector
	gridBoundaries grid.Vector
}

func (thisRobot Robot) getPosition(seconds int) grid.Vector {
	newPosition := grid.AddVectors(thisRobot.startingPos, grid.MultiplyVector(thisRobot.velocity, seconds))
	newPosition.X = newPosition.X % thisRobot.gridBoundaries.X
	newPosition.Y = newPosition.Y % thisRobot.gridBoundaries.Y

	if newPosition.Y < 0 {
		newPosition.Y = thisRobot.gridBoundaries.Y + newPosition.Y
	}
	if newPosition.X < 0 {
		newPosition.X = thisRobot.gridBoundaries.X + newPosition.X
	}

	return newPosition
}

func initialiseRobots(content string, xBoundary, yBoundary int) []Robot {
	robots := []Robot{}
	rows := strings.Split(content, "\n")
	for _, row := range rows {
		r := strings.Split(row, " ")
		startingPos := vectorFromString(r[0])
		velocity := vectorFromString(r[1])
		robots = append(robots, Robot{
			startingPos:    startingPos,
			velocity:       velocity,
			gridBoundaries: grid.Vector{X: xBoundary, Y: yBoundary},
		})
	}

	return robots
}

func vectorFromString(r string) grid.Vector {
	vectorStr := strings.Split(r, "=")
	vals := strings.Split(vectorStr[1], ",")
	return grid.Vector{X: utils.StringToInt(vals[0]), Y: utils.StringToInt(vals[1])}
}

func getPositionsAt(robots []Robot, seconds int) []grid.Vector {
	positions := []grid.Vector{}
	for _, robot := range robots {
		positions = append(positions, robot.getPosition(seconds))
	}
	return positions
}
