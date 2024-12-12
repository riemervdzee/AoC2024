package day08

import (
	"fmt"
	"riemer/utils"
)

type NodeMap map[uint8][]utils.Vector

func Process() {
	grid := utils.ReadFileAsGrid("day08/input.txt")
	part1Total := 0

	nodeMap := getNodes(grid)
	antiNodes := getAntiNodes(grid, nodeMap)
	part1Total = len(antiNodes)

	fmt.Println("Day 8 Results")
	fmt.Println("Part1", part1Total)
}

func getAntiNodes(grid utils.Grid, nodeMap NodeMap) map[utils.Vector]bool {
	antiNodes := make(map[utils.Vector]bool)
	for _, nodes := range nodeMap {
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {
				v1 := nodes[i]
				v2 := nodes[j]
				diff := utils.VectorSub(v2, v1)
				x1 := utils.VectorSub(v1, diff)
				x2 := utils.VectorAdd(v2, diff)

				if utils.GridValidPosition(grid, x1) {
					antiNodes[x1] = true
				}
				if utils.GridValidPosition(grid, x2) {
					antiNodes[x2] = true
				}
			}
		}
	}

	return antiNodes
}

// getNodes - generates a NodeMap from the given grid, using the signal as key
func getNodes(grid utils.Grid) NodeMap {
	nodes := make(NodeMap)
	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				_, ok := nodes[cell]
				vector := utils.Vector{x, y}
				if ok {
					nodes[cell] = append(nodes[cell], vector)
				} else {
					nodes[cell] = []utils.Vector{vector}
				}
			}
		}
	}
	return nodes
}
