package day12

import (
	"fmt"
	"riemer/utils"
)

type Region struct {
	PlantType uint8
	Plots     []utils.Vector
	Area      int
	Perimeter int
	Sides     int
}

func Process() {
	grid := utils.ReadFileAsGrid("day12/input.txt")
	regions := findRegions(grid)

	totalPrice1, totalPrice2 := 0, 0
	for _, region := range regions {
		totalPrice1 += region.Area * region.Perimeter
		totalPrice2 += region.Area * region.Sides
	}

	fmt.Println("Day 11 Results")
	fmt.Println("Part1", totalPrice1)
	fmt.Println("Part2", totalPrice2)
}

func findRegions(grid utils.Grid) []Region {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	var regions []Region
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if !visited[y][x] {
				region := exploreRegion(grid, visited, utils.Vector{x, y})
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func exploreRegion(grid utils.Grid, visited [][]bool, position utils.Vector) Region {
	plantType := grid[position[1]][position[0]]
	region := Region{
		PlantType: plantType,
		Plots:     []utils.Vector{},
	}

	floodFill(grid, visited, position, plantType, &region)
	region.Area = len(region.Plots)
	region.Perimeter = calculatePerimeter(grid, &region)
	region.Sides = calculateSides(grid, &region)

	return region
}

func floodFill(grid utils.Grid, visited [][]bool, position utils.Vector, goalPlantType uint8, region *Region) {
	plantType := grid[position[1]][position[0]]
	if plantType != goalPlantType {
		return
	}

	visited[position[1]][position[0]] = true
	region.Plots = append(region.Plots, position)

	for _, direction := range utils.FourDirections {
		newPosition := utils.VectorAdd(position, direction)
		if utils.GridValidPosition(grid, newPosition) && !visited[newPosition[1]][newPosition[0]] {
			floodFill(grid, visited, newPosition, plantType, region)
		}
	}
}

func calculatePerimeter(grid utils.Grid, region *Region) int {
	perimeter := 0

	for _, plot := range region.Plots {
		for _, direction := range utils.FourDirections {
			newPos := utils.VectorAdd(plot, direction)
			if !utils.GridValidPosition(grid, newPos) {
				perimeter++
			} else if grid[newPos[1]][newPos[0]] != region.PlantType {
				perimeter++
			}
		}
	}
	return perimeter
}

func calculateSides(grid utils.Grid, region *Region) int {
	return 0
}
