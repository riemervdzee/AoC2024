package day14

import (
	"fmt"
	"math"
	"regexp"
	"riemer/utils"
)

var re = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

const (
	gridWidth    = 101
	gridHeight   = 103
	secondsPart1 = 100
	gridWM       = gridWidth / 2
	gridHM       = gridHeight / 2
)

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
	var part1, part2 int
	safetyFactorP2 := math.MaxInt64
	quadrants := make([]int, 4)

	for i := 1; i < gridWidth*gridHeight; i++ {
		// Move all robots and get quadrant scores
		for i := range quadrants {
			quadrants[i] = 0
		}
		for i := range robots {
			robots[i].pos[0] = (robots[i].pos[0] + robots[i].vel[0] + gridWidth) % gridWidth
			robots[i].pos[1] = (robots[i].pos[1] + robots[i].vel[1] + gridHeight) % gridHeight
			x, y := robots[i].pos[0], robots[i].pos[1]
			if x != gridWM && y != gridHM {
				quadrant := 0
				if x < gridWM {
					quadrant++
				}
				if y < gridHM {
					quadrant += 2
				}
				quadrants[quadrant]++
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
