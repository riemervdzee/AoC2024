package utils

import (
	"bufio"
	"os"
)

type Grid [][]byte

// ReadFileAsGrid - Reads the file as an 2-dimensional grid array
func ReadFileAsGrid(filename string) Grid {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	var grid Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid
}

// GridValidPosition - Check if a position is valid inside the grid
func GridValidPosition(grid Grid, position Vector) bool {
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	x := position[0]
	y := position[1]
	return x >= 0 && x < gridWidth && y >= 0 && y < gridHeight
}
