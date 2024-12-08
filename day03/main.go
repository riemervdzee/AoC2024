package day03

import (
	"fmt"
	"os"
	"regexp"
	"riemer/utils"
	"strconv"
)

func Process() {
	part1 := processSub(false)
	part2 := processSub(true)

	fmt.Println("Day 3 Results")
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}

func processSub(strip bool) int {
	// Read the file as one string
	b, err := os.ReadFile("day03/input.txt")
	utils.Check(err)
	input := string(b)

	// If requested strip everything between `don't()` and `do()`
	if strip {
		r := regexp.MustCompile(`don't\(\).*?do\(\)`)
		input = r.ReplaceAllString(input, "")
	}

	//
	multiplicationSum := 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		num1, _ := strconv.Atoi(v[1])
		num2, _ := strconv.Atoi(v[2])
		multiplicationSum += num1 * num2
	}

	return multiplicationSum
}
