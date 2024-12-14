package day04

import (
	"fmt"
	"riemer/utils"
)

func Process() {
	grid := utils.ReadFileAsGrid("day04/input.txt")

	part1 := processPart1(grid)
	part2 := processPart2(grid)

	fmt.Println("Day 4 Results")
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}

func processPart1(grid utils.Grid) int {
	total := 0
	for y, row := range grid {
		for x, cell := range row {
			// Check if current cell is X
			if cell == 'X' {
				position := utils.Vector{x, y}
				// Check all directions and search for MAS
				for _, direction := range utils.AllDirections {
					total += scanForString(grid, "MAS", position, direction)
				}
			}
		}
	}
	return total
}

func processPart2(grid utils.Grid) int {
	total := 0
	for y, row := range grid {
		for x, cell := range row {
			// We center around A now
			if cell == 'A' {
				position := utils.Vector{x, y}
				// Check up-left to right-down diagonal
				diag1 := scanForString(grid, "M", position, utils.DirUpLeft)+scanForString(grid, "S", position, utils.DirDownRight) == 2 ||
					scanForString(grid, "S", position, utils.DirUpLeft)+scanForString(grid, "M", position, utils.DirDownRight) == 2
				// Check up-right to down-left diagonal
				diag2 := scanForString(grid, "M", position, utils.DirUpRight)+scanForString(grid, "S", position, utils.DirDownLeft) == 2 ||
					scanForString(grid, "S", position, utils.DirUpRight)+scanForString(grid, "M", position, utils.DirDownLeft) == 2

				// If both contain MAS one way or another, increase total
				if diag1 && diag2 {
					total += 1
				}
			}
		}
	}
	return total
}

// scanForString - Checks if the search needle continues at `position` in the `direction` in the grid
func scanForString(grid utils.Grid, needle string, position utils.Vector, direction utils.Vector) int {
	newPosition := utils.VectorAdd(position, direction)
	if utils.GridValidPosition(grid, newPosition) && grid[newPosition[1]][newPosition[0]] == needle[0] {
		// Check if we have a search-needle remaining, use recursion. Otherwise return 1 as we completed the search
		if len(needle) > 1 {
			newNeedle := needle[1:]
			return scanForString(grid, newNeedle, newPosition, direction)
		}
		return 1
	}

	return 0
}
