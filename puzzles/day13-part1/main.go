package main

import (
	"errors"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

type Machine struct {
	ButtonA grid.Vector
	ButtonB grid.Vector
	Target  grid.Vector
}

func main() {
	filename := "../../data/day13.txt"
	content := utils.GetFileContents(filename)
	machines := getMachines(content)
	fewestTokens := getTotalFewestTokens(machines)

	println("Answer:", fewestTokens)
}

func getTotalFewestTokens(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		total += getFewestTokens(machine)
	}
	return total
}

func getFewestTokens(machine Machine) int {
	fewestTokens := 1_000_000
	aButtonTries := 0

	for aButtonTries <= 100 {
		/*
			attempt 1: B B B B B B B B B B
			attempt 2: A B B B B B B B B B B
			attempt 3: A A B B B B B B B B B B
		*/
		tokensCost, err := attempt(machine, aButtonTries)
		if err != nil {
			aButtonTries++
			continue
		}
		fewestTokens = min(fewestTokens, tokensCost)
		aButtonTries++
	}

	if fewestTokens == 1_000_000 {
		return 0
	}

	return fewestTokens
}

func attempt(machine Machine, aButtonTries int) (int, error) {
	aButtonToTry := aButtonTries
	aTries := 0
	bButtonToTry := 100
	bTries := 0
	costA := 3
	costB := 1
	target := machine.Target
	current := grid.Vector{X: 0, Y: 0}

	for current.X < target.X && current.Y < target.Y {

		if aTries < aButtonToTry {
			current = grid.AddVectors(current, machine.ButtonA)
			aTries++
		} else if bTries < bButtonToTry {
			current = grid.AddVectors(current, machine.ButtonB)
			bTries++
		} else {
			// No more tries
			return 0, errors.New("unsuccessful")
		}

		if current.X == target.X && current.Y == target.Y {
			// Successful attempt
			return costA*aTries + costB*bTries, nil
		}
	}

	return 0, errors.New("unsuccessful")
}

func getMachines(content string) []Machine {
	blocks := strings.Split(content, "\n\n")
	machines := []Machine{}
	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		buttonA := getButtonVals(rows[0])
		buttonB := getButtonVals(rows[1])
		prize := strings.Split(rows[2], "=")
		prizeX := utils.StringToInt(strings.Split(prize[1], ",")[0])
		prizeY := utils.StringToInt(prize[len(prize)-1])
		machines = append(machines, Machine{
			ButtonA: buttonA,
			ButtonB: buttonB,
			Target:  grid.Vector{X: prizeX, Y: prizeY},
		})
	}
	return machines
}

func getButtonVals(row string) grid.Vector {
	chunks := strings.Split(row, "+")
	x := utils.StringToInt(strings.Split(chunks[1], ",")[0])
	y := utils.StringToInt(chunks[len(chunks)-1])
	return grid.Vector{X: x, Y: y}
}
