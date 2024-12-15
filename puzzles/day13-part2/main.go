package main

import (
	"math"
	"strings"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

/*
	aX*aPresses + bX*bPresses = targetX
	aY*aPresses + bY*bpresses = targetY

	94a + 22b = 8400
	34a + 67b = 5400

	Express a in terms of b:
	94a = 8400 - 22b
	a = (8400 - 22b) / 94

	Substitute a:
	34*((8400 - 22b) / 94) + 67b = 5400

	((285600 - 748b) / 94) + 67b = 5400

	285600 - 748b + 6298b = 507600
	5550b = 222000
	b = 40 <===

	Solve for a:
	34a + 67b = 5400
	34a + 67*40 = 5400
	34a + 2680 = 5400
	a = 80 <===

	34*((8400 - 22b) / 94) + 67b = 5400
	34*8400 - 34*22b + 94*67b = 5400*94
	-34*22b + 94*67b = 5400*94 - 34*8400
	b(-34*22 + 94*67) = 5400*94 - 34*8400
	b = (5400*94 - 34*8400) / (-34*22 + 94*67)
*/

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
	targetX := float64(machine.Target.X)
	targetY := float64(machine.Target.Y)
	buttonAy := float64(machine.ButtonA.Y)
	buttonAx := float64(machine.ButtonA.X)
	buttonBy := float64(machine.ButtonB.Y)
	buttonBx := float64(machine.ButtonB.X)
	bPresses := (targetY*buttonAx - buttonAy*targetX) /
		(-buttonAy*buttonBx + buttonAx*buttonBy)

	aPresses := (targetY - buttonBy*bPresses) / buttonAy

	if math.Mod(aPresses, 1) == 0 && math.Mod(bPresses, 1) == 0 {
		return int(3*aPresses + bPresses)
	}

	return 0
}

func getMachines(content string) []Machine {
	bigNumber := 10000000000000
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
			Target:  grid.Vector{X: prizeX + bigNumber, Y: prizeY + bigNumber},
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
