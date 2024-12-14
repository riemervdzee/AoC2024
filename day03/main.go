package day03

import (
	"fmt"
	"os"
	"regexp"
	"riemer/utils"
	"strconv"
)

var regexDont = regexp.MustCompile(`don't\(\).*?do\(\)`)
var regexMulp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func Process() {
	// Read the file as one string
	b, err := os.ReadFile("day03/input.txt")
	utils.Check(err)
	input := string(b)

	part1 := processSub(input, false)
	part2 := processSub(input, true)

	fmt.Println("Day 3 Results")
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}

func processSub(input string, strip bool) int {
	// If requested strip everything between `don't()` and `do()`
	if strip {
		input = regexDont.ReplaceAllString(input, "")
	}

	// Calculate multiplications
	multiplicationSum := 0
	matches := regexMulp.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		num1, _ := strconv.Atoi(v[1])
		num2, _ := strconv.Atoi(v[2])
		multiplicationSum += num1 * num2
	}

	return multiplicationSum
}
