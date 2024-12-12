package day07

import (
	"fmt"
	"riemer/utils"
	"slices"
	"strings"
)

func Process() {
	lines := utils.ReadFile("day07/input.txt")
	part1Total := 0
	part2Total := 0
	for _, line := range lines {
		stripped := strings.Replace(line, ":", "", 1)
		args := utils.StringArrayToIntArray(strings.Fields(stripped))

		if solve(args[0], args[2:], args[1:2], false) {
			part1Total += args[0]
		}
		if solve(args[0], args[2:], args[1:2], true) {
			part2Total += args[0]
		}
	}

	fmt.Println("Day 7 Results")
	fmt.Println("Part1", part1Total)
	fmt.Println("Part2", part2Total)
}

func solve(goal int, args []int, results []int, concat bool) bool {
	// Exit condition. If we have no more operands to process, check for result
	if len(args) == 0 {
		return slices.Contains(results, goal)
	}

	// Create new resultArray for new Solve() call
	lenMul := 2
	if concat {
		lenMul = 3
	}
	resultsNew := make([]int, 0, len(results)*lenMul)

	// Fill resultsNew where we do all requested operands (add, multiplication, concat if requested)
	input := args[0]
	for i := 0; i < len(results); i++ {
		resultsNew = append(resultsNew, results[i]+input)
		resultsNew = append(resultsNew, results[i]*input)
		if concat {
			resultsNew = append(resultsNew, concatint(results[i], input))
		}
	}

	// New iteration of Solve()
	return solve(goal, args[1:], resultsNew, concat)
}

// Fast contacting of 2 integers
func concatint(a int, b int) int {
	if b < 10 {
		return a*10 + b
	} else if b < 100 {
		return a*100 + b
	} else if b < 1000 {
		return a*1000 + b
	}
	panic("B is too large!")
}