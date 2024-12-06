package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector [2]int

var (
	DIR_UP         = Vector{0, -1}
	DIR_DOWN       = Vector{0, 1}
	DIR_LEFT       = Vector{-1, 0}
	DIR_RIGHT      = Vector{1, 0}
	DIR_UP_LEFT    = Vector{-1, -1}
	DIR_UP_RIGHT   = Vector{1, -1}
	DIR_DOWN_LEFT  = Vector{-1, 1}
	DIR_DOWN_RIGHT = Vector{1, 1}
)

var directionsPart1 = []Vector{
	DIR_UP,
	DIR_DOWN,
	DIR_LEFT,
	DIR_RIGHT,
	DIR_UP_LEFT,
	DIR_UP_RIGHT,
	DIR_DOWN_LEFT,
	DIR_DOWN_RIGHT,
}

func main() {
	fmt.Println("Day 4 Part 1")
	processPart1()

	fmt.Println("Day 4 Part 2")
	processPart2()
}

func processPart1() {
	grid := readFileAsGrid("input.txt")

	total := 0
	for y, row := range grid {
		for x, cell := range row {
			// Check if current cell is X
			if cell == 'X' {
				position := Vector{x, y}
				// Check all directions and search for MAS
				for _, direction := range directionsPart1 {
					total += scanForString(grid, "MAS", position, direction)
				}
			}
		}
	}
	fmt.Println(total)
}

func processPart2() {
	grid := readFileAsGrid("input.txt")

	total := 0
	for y, row := range grid {
		for x, cell := range row {
			// We center around A now
			if cell == 'A' {
				position := Vector{x, y}
				// Check up-left to right-down diagonal
				diag1 := scanForString(grid, "M", position, DIR_UP_LEFT)+scanForString(grid, "S", position, DIR_DOWN_RIGHT) == 2 ||
					scanForString(grid, "S", position, DIR_UP_LEFT)+scanForString(grid, "M", position, DIR_DOWN_RIGHT) == 2
				// Check up-right to down-left diagonal
				diag2 := scanForString(grid, "M", position, DIR_UP_RIGHT)+scanForString(grid, "S", position, DIR_DOWN_LEFT) == 2 ||
					scanForString(grid, "S", position, DIR_UP_RIGHT)+scanForString(grid, "M", position, DIR_DOWN_LEFT) == 2

				// If both contain MAS one way or another, increase total
				if diag1 && diag2 {
					total += 1
				}
			}
		}
	}
	fmt.Println(total)
}

// scanForString - Checks if the search needle continues at `position` in the `direction` in the grid
func scanForString(grid [][]byte, needle string, position Vector, direction Vector) int {
	newPosition := VectorAdd(position, direction)
	if validPosition(grid, newPosition) && grid[newPosition[1]][newPosition[0]] == needle[0] {
		// Check if we have a search-needle remaining, use recursion. Otherwise return 1 as we completed the search
		newNeedle := needle[1:]
		if len(newNeedle) > 0 {
			return scanForString(grid, newNeedle, newPosition, direction)
		}
		return 1
	}

	return 0
}

// VectorAdd - Simply adds to 2 vectors and return it
func VectorAdd(a, b Vector) Vector {
	return Vector{a[0] + b[0], a[1] + b[1]}
}

// validPosition - Check if a position is valid inside the grid
func validPosition(grid [][]byte, position Vector) bool {
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	x := position[0]
	y := position[1]
	return x >= 0 && x < gridWidth && y >= 0 && y < gridHeight
}

// readFileAsGrid - Reads the file as an 2-dimensional grid array
func readFileAsGrid(filename string) [][]byte {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	return grid
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
