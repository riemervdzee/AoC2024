package day06

import (
	"fmt"
	"riemer/utils"
)

type Player struct {
	position  utils.Vector
	direction utils.Vector
}

func Process() {
	total1 := processPart1()
	fmt.Println("Day 6 Results")
	fmt.Println("Part1", total1)
}

func processPart1() int {
	grid := utils.ReadFileAsGrid("day06/input.txt")
	player := FindPlayer(grid)

	for {
		newPosition := utils.VectorAdd(player.position, player.direction)
		if !utils.GridValidPosition(grid, newPosition) {
			break
		}

		lookahead := grid[newPosition[1]][newPosition[0]]
		fmt.Println("Stepping at", newPosition, lookahead)

		if lookahead != '#' {
			player.position = newPosition
			grid[newPosition[1]][newPosition[0]] = 'X'
		} else {
			fmt.Println("Blocked at", newPosition)
			player.direction = utils.VectorTurnRight(player.direction)
		}
	}

	// Count Distinct places
	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 'X' || cell == '^' {
				total++
			}
		}
	}
	return total
}

func FindPlayer(grid utils.Grid) Player {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return Player{
					position:  utils.Vector{x, y},
					direction: utils.DirUp,
				}
			}
		}
	}
	panic("Player not found")
}
