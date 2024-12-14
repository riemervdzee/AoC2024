package day02

import (
	"fmt"
	"riemer/utils"
	"strconv"
	"strings"
)

func Process() {
	safeStrictCount := 0
	safeDampedCount := 0

	fileLines := utils.ReadFile("day02/input.txt")
	for _, line := range fileLines {
		parts := strings.Fields(line)
		integers := make([]int, 0, len(parts))
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			utils.Check(err)
			integers = append(integers, num)
		}

		safe := checkSlopeSlice(integers)
		if safe {
			safeStrictCount++
			safeDampedCount++
		} else {
			for i := range integers {
				newIntegers := append(integers[:i:i], integers[i+1:]...)
				safe = checkSlopeSlice(newIntegers)
				if safe {
					safeDampedCount++
					break
				}
			}
		}
	}

	fmt.Println("Day 2 Results")
	fmt.Println("Strict", safeStrictCount)
	fmt.Println("Part2", safeDampedCount)
}

// checkSlopeSlice - returns whether the input slice meets the slope criteria
func checkSlopeSlice(input []int) bool {
	ascending := input[0] < input[1]
	for i := 1; i < len(input); i++ {
		if !checkSlope(input[i-1], input[i], ascending) {
			return false
		}
	}

	return true
}

// checkSlope - checks the slope criteria for 2 elements
func checkSlope(prev int, curr int, asc bool) bool {
	return prev != curr && utils.Abs(prev-curr) <= 3 && (prev < curr) == asc
}
