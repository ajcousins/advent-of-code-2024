package main

import (
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

func main() {
	filename := "../../data/day14.txt"
	content := utils.GetFileContents(filename)
	robots := initialiseRobots(content, 101, 103)
	positions := getPositionsAt(robots, 100)
	answer := quadrants(positions, robots)

	println("Answer:", answer)
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

func quadrants(positions []grid.Vector, robots []Robot) int {
	xMiddleIndex := robots[0].gridBoundaries.X / 2
	yMiddleIndex := robots[0].gridBoundaries.Y / 2
	counts := [4]int{}

	for _, pos := range positions {
		if pos.X < xMiddleIndex {
			if pos.Y < yMiddleIndex {
				counts[0]++
			} else if pos.Y > yMiddleIndex {
				counts[2]++
			}
		} else if pos.X > xMiddleIndex {
			if pos.Y < yMiddleIndex {
				counts[1]++
			} else if pos.Y > yMiddleIndex {
				counts[3]++
			}
		}
	}

	return counts[0] * counts[1] * counts[2] * counts[3]
}
