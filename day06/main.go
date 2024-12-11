package day06

import (
	"fmt"
	"riemer/utils"
)

type Guard struct {
	position  utils.Vector
	direction utils.Vector
}

func Process() {
	total1 := processPart1()
	total2 := processPart2()
	fmt.Println("Day 6 Results")
	fmt.Println("Part1", total1)
	fmt.Println("Part2", total2)
}

func processPart1() int {
	grid := utils.ReadFileAsGrid("day06/input.txt")
	_, total := simulateGuard(grid, false)
	return total
}

func processPart2() int {
	grid := utils.ReadFileAsGrid("day06/input.txt")
	guard := findGuard(grid)

	validPositions := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '.' && (x != guard.position[0] || y != guard.position[1]) {
				grid[y][x] = '#'
				loop, _ := simulateGuard(grid, true)
				if loop {
					validPositions++
				}
				grid[y][x] = '.'
			}
		}
	}
	return validPositions
}

// simulateGuard - Simulates the movements of a guard inside a grid
func simulateGuard(grid utils.Grid, detectLoops bool) (bool, int) {
	guard := findGuard(grid)
	score := 1
	visited := map[int]bool{}

	for {
		// Set new position and check if still in bounds of the grid
		newPosition := utils.VectorAdd(guard.position, guard.direction)
		if !utils.GridValidPosition(grid, newPosition) {
			break
		}

		// Loop detection
		if detectLoops {
			guardKey := (newPosition[0] << 20) | (newPosition[1] << 12) | ((guard.direction[0] + 2) << 4) | (guard.direction[1] + 2)
			if visited[guardKey] {
				return true, score
			}
			visited[guardKey] = true
		}

		// Progress Guard movements
		lookahead := grid[newPosition[1]][newPosition[0]]
		if lookahead != '#' {
			guard.position = newPosition
			if !detectLoops && lookahead != 'X' {
				score++
				grid[newPosition[1]][newPosition[0]] = 'X'
			}
		} else {
			guard.direction = utils.VectorTurnRight(guard.direction)
		}
	}

	return false, score
}

// findPlayer - Returns a Guard object at the position found in the grid
func findGuard(grid utils.Grid) Guard {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return Guard{
					position:  utils.Vector{x, y},
					direction: utils.DirUp,
				}
			}
		}
	}
	panic("Guard not found")
}
