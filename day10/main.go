package day10

import (
	"fmt"
	"riemer/utils"
)

type QueueItem struct {
	dis byte
	pos utils.Vector
}

type Queue []QueueItem

func Process() {
	grid := utils.ReadFileAsGrid("day10/input.txt")

	part1Total, part2Total := 0, 0

	for _, point := range findStartPositions(grid) {
		r1, r2 := findTrailHeadCount(grid, point)
		part1Total += r1
		part2Total += r2
	}

	fmt.Println("Day 10 Results")
	fmt.Println("Part1", part1Total)
	fmt.Println("Part2", part2Total)
}

func findTrailHeadCount(grid utils.Grid, start utils.Vector) (int, int) {
	finished1 := make(map[utils.Vector]bool)
	finished2 := 0
	queue := make(Queue, 0)
	queue = append(queue, QueueItem{0, start})

	for len(queue) > 0 {
		// pop
		x := queue[0]
		queue = queue[1:]

		for _, direction := range utils.FourDirections {
			distance := x.dis + 1
			newPosition := utils.VectorAdd(x.pos, direction)

			if utils.GridValidPosition(grid, newPosition) {
				gridPos := grid[newPosition[1]][newPosition[0]] - '0'
				if gridPos == distance {
					if gridPos == 9 {
						finished1[newPosition] = true
						finished2++
					} else {
						queue = append(queue, QueueItem{distance, newPosition})
					}
				}
			}
		}
	}

	return len(finished1), finished2
}

// findStart - Returns a Vector at the start position found in the grid
func findStartPositions(grid utils.Grid) []utils.Vector {
	starts := make([]utils.Vector, 0)
	for y, row := range grid {
		for x, cell := range row {
			if cell == '0' {
				starts = append(starts, utils.Vector{x, y})
			}
		}
	}
	if len(starts) == 0 {
		panic("Start not found")
	}
	return starts
}
