package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("input.txt")
	ruleLines, updateLines := splitLines(lines)
	ruleMap := generateRuleMap(ruleLines)

	totalPart1 := 0
	totalPart2 := 0
	for _, updateLine := range updateLines {
		update := stringArrayToIntArray(strings.Split(updateLine, ","))
		correct := sortUpdate(ruleMap, update)
		middle := len(update) / 2

		// Part 1
		if correct {
			totalPart1 += update[middle]
		} else {
			totalPart2 += update[middle]
		}
	}

	fmt.Println("Day 5 Part 1")
	fmt.Println(totalPart1)

	fmt.Println("Day 5 Part 2")
	fmt.Println(totalPart2)
}

// sortUpdate - Sorts an update according to the rulebook and returns whether it was correct in the first place or not
func sortUpdate(ruleMap map[int][]int, update []int) bool {
	correct := true
	sort.Slice(update, func(i, j int) bool {
		if !checkValue(ruleMap, update[i], update[j]) {
			correct = false
			return true
		}
		return false
	})
	return correct
}

// checkValue - Checks if X is valid before Y
func checkValue(ruleMap map[int][]int, x int, y int) bool {
	rules, ok := ruleMap[x]
	return !ok || !slices.Contains(rules, y)
}

// generateRuleMap - We split the string `Val1|Val2` into a `map[Val1] = []Val2`
func generateRuleMap(rules []string) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		key, _ := strconv.Atoi(split[0])
		value, _ := strconv.Atoi(split[1])
		_, ok := ruleMap[key]
		if ok {
			ruleMap[key] = append(ruleMap[key], value)
		} else {
			ruleMap[key] = []int{value}
		}
	}
	return ruleMap
}

// splitLines - Splits the lines input into the Page Ordering Rules and Updates lines
func splitLines(lines []string) ([]string, []string) {
	// We search for an empty string,and slice the array accordingly
	for i, line := range lines {
		if line == "" {
			// Make sure to skip the empty string
			return lines[:i], lines[i+1:]
		}
	}
	panic("No Empty String delimiter found!")
}

// readFile - Reads the file as lines
func readFile(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Helper func, converts a []string to []int
func stringArrayToIntArray(stringArray []string) []int {
	integers := make([]int, 0, len(stringArray))
	for _, part := range stringArray {
		num, err := strconv.Atoi(part)
		check(err)
		integers = append(integers, num)
	}
	return integers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
