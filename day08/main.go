package day08

import (
	"fmt"
	"riemer/utils"
)

type NodeMap map[uint8][]utils.Vector

func Process() {
	grid := utils.ReadFileAsGrid("day08/input.txt")

	nodeMap := getNodes(grid)
	part1Total := len(getAntiNodes(grid, nodeMap, false))
	part2Nodes := getAntiNodes(grid, nodeMap, true)

	// Also mark antenna's in part2Nodes
	for _, nodes := range nodeMap {
		for _, node := range nodes {
			part2Nodes[node] = true
		}
	}
	part2Total := len(part2Nodes)

	fmt.Println("Day 8 Results")
	fmt.Println("Part1", part1Total)
	fmt.Println("Part2", part2Total)
}

func getAntiNodes(grid utils.Grid, nodeMap NodeMap, repeat bool) map[utils.Vector]bool {
	antiNodes := make(map[utils.Vector]bool)
	for _, nodes := range nodeMap {
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {
				v1 := nodes[i]
				v2 := nodes[j]
				diff := utils.VectorSub(v2, v1)
				x1 := utils.VectorSub(v1, diff)
				x2 := utils.VectorAdd(v2, diff)

				for utils.GridValidPosition(grid, x1) {
					antiNodes[x1] = true
					x1 = utils.VectorSub(x1, diff)
					if !repeat {
						break
					}
				}
				for utils.GridValidPosition(grid, x2) {
					antiNodes[x2] = true
					x2 = utils.VectorAdd(x2, diff)
					if !repeat {
						break
					}
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
				vector := utils.Vector{x, y}
				nodes[cell] = append(nodes[cell], vector)
			}
		}
	}
	return nodes
}
