package main

import (
	"math"

	"github.com/ajcousins/advent-of-code-2024/utils"
	"github.com/ajcousins/advent-of-code-2024/utils/grid"
)

type NodeState struct {
	address grid.Vector
	facing  grid.Vector
}

type Step struct {
	NodeState
	cost int
}

type WeightedGraph struct {
	adjacencies map[NodeState][]Step
}

func main() {
	filename := "../../data/day16.txt"
	content := utils.GetFileContents(filename)
	maze := grid.New(content)
	answer := getLowestScore(maze)

	println("Answer:", answer)
}

func getLowestScore(maze grid.Grid) int {
	startPos, _ := maze.GetAddressFromValue("S")
	endPos, _ := maze.GetAddressFromValue("E")
	graph := WeightedGraph{}
	graph.InitialiseAdjacencies(maze)

	return dijkstras(graph, startPos, endPos)
}

func dijkstras(graph WeightedGraph, startPos, endPos grid.Vector) int {
	LARGE_NUMBER := 1_000_000
	initialNodeState := NodeState{
		address: startPos,
		facing:  grid.Vector{Y: 0, X: 1}, // Start facing East
	}

	// Populate distances map
	distances := map[NodeState]int{}
	for key := range graph.adjacencies {
		if key == initialNodeState {
			distances[key] = 0
			continue
		}
		distances[key] = LARGE_NUMBER
	}

	// Populate queue
	queue := utils.PriorityQueue[NodeState]{}
	for key := range graph.adjacencies {
		if key == initialNodeState {
			queue.Enqueue(initialNodeState, 0)
			continue
		}
		queue.Enqueue(key, LARGE_NUMBER)
	}

	previous := map[NodeState]NodeState{}

	for queue.Length() > 0 {
		smallest := queue.Dequeue()
		if smallest.address == endPos {
			return distances[smallest]
		}
		neighbours := graph.adjacencies[smallest]
		for _, neighbour := range neighbours {
			candidate := distances[smallest] + neighbour.cost
			if candidate < distances[neighbour.NodeState] {
				distances[neighbour.NodeState] = candidate
				previous[neighbour.NodeState] = smallest
				queue.Enqueue(neighbour.NodeState, candidate)
			}
		}
	}
	return 0
}

func (wg *WeightedGraph) InitialiseAdjacencies(maze grid.Grid) {
	/*
		Create a map of "NodeStates" for each valid location
		and facing direction. Adjacencies should include the cost
		of moving to each neighbouring node.
		...
		{8 11}_West  : adjs: [{{7 11} 1001} {{9 11} 1001}],
		{8 11}_South : adjs: [{{7 11} 2001} {{9 11} 1}],
		{1 9}_North  : adjs: [{{1 10} 1001} {{2 9} 2001}],
		...
	*/
	wg.adjacencies = make(map[NodeState][]Step)

	dirs := grid.GetDirs()
	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			address := grid.Vector{Y: y, X: x}

			cellContent := maze.Get(address)
			if cellContent == "#" {
				continue
			}
			for _, dir := range dirs {
				nodeState := NodeState{
					address: address,
					facing:  dir,
				}
				adjacencies := getAdjacencies(nodeState, maze)
				wg.adjacencies[nodeState] = adjacencies
			}
		}
	}
}

func getAdjacencies(nodeState NodeState, maze grid.Grid) []Step {
	steps := []Step{}
	dirs := grid.GetDirs()
	for _, dir := range dirs {
		adjAddress := grid.AddVectors(nodeState.address, dir)
		adjContent := maze.Get(adjAddress)
		if adjContent == "#" {
			continue
		}
		newStep := Step{
			NodeState: NodeState{
				address: adjAddress,
				facing:  dir,
			},
			cost: getCost(nodeState.facing, dir),
		}
		steps = append(steps, newStep)
	}
	return steps
}

func getCost(dirA, dirB grid.Vector) int {
	combined := grid.AddVectors(dirA, dirB)
	xMag := math.Abs(float64(combined.X))
	yMag := math.Abs(float64(combined.Y))
	if xMag == 2 || yMag == 2 {
		// facing the same directon
		return 1
	}
	if xMag == 0 && yMag == 0 {
		// facing the opposite direction
		return 2001
	}
	// facing 90-degs
	return 1001
}
