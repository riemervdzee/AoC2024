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
const seconds = 100
const gridWM = gridWidth / 2
const gridHM = gridHeight / 2

type Robot struct {
	pos utils.Vector
	vel utils.Vector
}

func Process() {
	lines := utils.ReadFile("day14/input.txt")
	robots := make([]Robot, 0, len(lines))
	var quadrantsPart1 [4]int

	// Read lines into robots and solve part 1
	for _, line := range lines {
		// Get position, velocity and create a robot
		matches := re.FindStringSubmatch(line)
		p := utils.Vector{utils.StringToInt(matches[1]), utils.StringToInt(matches[2])}
		v := utils.Vector{utils.StringToInt(matches[3]), utils.StringToInt(matches[4])}
		robots = append(robots, Robot{p, v})

		// Solve part 1
		v100 := utils.Vector{v[0] * seconds, v[1] * seconds}
		p1 := utils.VectorAdd(p, v100)
		p1 = utils.Vector{modPositive(p1[0], gridWidth), modPositive(p1[1], gridHeight)}
		if index, found := quadIndex(p1); found {
			quadrantsPart1[index] = quadrantsPart1[index] + 1
		}
	}

	fmt.Println("Day 14 Results")
	fmt.Println("Part1", quadrantsPart1[0]*quadrantsPart1[1]*quadrantsPart1[2]*quadrantsPart1[3])

	// For part 2 we try to find the lowest quadrant-score, this suggests there is a nice distribution of all points in each quadrant
	safetyFactor := math.MaxInt64
	foundSeconds := 0
	for i := 1; i < gridWidth*gridHeight; i++ {
		var quadrants [4]int
		for i := range robots {
			robots[i].pos = utils.VectorAdd(robots[i].pos, robots[i].vel)
			robots[i].pos = utils.Vector{modPositive(robots[i].pos[0], gridWidth), modPositive(robots[i].pos[1], gridHeight)}
			if index, ok := quadIndex(robots[i].pos); ok {
				quadrants[index] = quadrants[index] + 1
			}
		}

		sf := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
		if sf < safetyFactor {
			safetyFactor = sf
			foundSeconds = i
		}
	}

	fmt.Println("Part2", foundSeconds)
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
