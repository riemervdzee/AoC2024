package day14

import (
	"fmt"
	"math"
	"regexp"
	"riemer/utils"
)

var re = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

const gridWidth = 101
const gridHeight = 103
const secondsPart1 = 100
const gridWM = gridWidth / 2
const gridHM = gridHeight / 2

type Robot struct {
	pos utils.Vector
	vel utils.Vector
}

func Process() {
	lines := utils.ReadFile("day14/input.txt")
	robots := make([]Robot, 0, len(lines))

	// Read lines into robots
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		robots = append(robots, Robot{
			utils.Vector{utils.StringToInt(matches[1]), utils.StringToInt(matches[2])},
			utils.Vector{utils.StringToInt(matches[3]), utils.StringToInt(matches[4])}},
		)
	}

	// Just loop an X amount of times
	part1 := 0
	part2 := 0
	safetyFactorP2 := math.MaxInt64
	for i := 1; i < gridWidth*gridHeight; i++ {
		// Move all robots and get quadrant scores
		var quadrants [4]int
		for i := range robots {
			robots[i].pos = utils.VectorAdd(robots[i].pos, robots[i].vel)
			robots[i].pos = utils.Vector{modPositive(robots[i].pos[0], gridWidth), modPositive(robots[i].pos[1], gridHeight)}
			if index, ok := quadIndex(robots[i].pos); ok {
				quadrants[index] = quadrants[index] + 1
			}
		}

		// Part 1 and 2 handling
		sf := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
		if i == secondsPart1 {
			part1 = sf
		}
		if sf < safetyFactorP2 {
			safetyFactorP2 = sf
			part2 = i
		}
	}

	fmt.Println("Day 14 Results")
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}

// quadIndex returns in which quadrant a position is inside
func quadIndex(vector utils.Vector) (quad int, found bool) {
	if vector[0] == gridWM || vector[1] == gridHM {
		return quad, false
	}
	if vector[0] < gridWM {
		quad++
	}
	if vector[1] < gridHM {
		quad += 2
	}
	return quad, true
}

// mod is a modulo with a positive remainder
func modPositive(a, b int) int {
	return (a%b + b) % b
}
