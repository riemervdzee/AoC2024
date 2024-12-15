package day11

import (
	"fmt"
	"math"
	"riemer/utils"
	"strings"
)

// StoneMap - Instead of a big list, we use Maps with as key the stone number, and the value as the counts
type StoneMap map[int]int

func Process() {
	lines := utils.ReadFile("day11/input.txt")
	blinkTimesP1 := 25
	blinkTimesP2 := 75
	stones := utils.StringArrayToIntArray(strings.Fields(lines[0]))
	stoneMap := generateStoneMap(stones)

	stoneMap = blinkTimes(stoneMap, blinkTimesP1)
	fmt.Println("Day 11 Results")
	fmt.Println("Part1", countStoneMap(stoneMap))

	stoneMap = blinkTimes(stoneMap, blinkTimesP2-blinkTimesP1)
	fmt.Println("Part2", countStoneMap(stoneMap))
}

// blinkTimes - just blink X times
func blinkTimes(stoneMap StoneMap, amount int) StoneMap {
	for i := 0; i < amount; i++ {
		stoneMap = blink(stoneMap)
	}
	return stoneMap
}

// blink - processes each stone group and returns a new StoneMap
func blink(stones StoneMap) StoneMap {
	ret := make(StoneMap)
	for stone, count := range stones {
		newStones := processStone(stone)
		ret[newStones[0]] += count
		if len(newStones) > 1 {
			ret[newStones[1]] += count
		}
	}
	return ret
}

// processStone -  processes a single stone
func processStone(num int) []int {
	if num == 0 {
		return []int{1}
	}
	digits := countDigits(num)
	if digits%2 == 0 {
		left, right := splitNumber(num, digits)
		return []int{left, right}
	}
	return []int{num * 2024}
}

// splitNumber - Splits the number in two halves according to the digits
func splitNumber(num int, digits int) (int, int) {
	power := int(math.Pow10(digits / 2))
	left := num / power
	right := num % power
	return left, right
}

// countDigits - Count digits fast
func countDigits(i int) int {
	if i < 10 {
		return 1
	}
	if i < 100 {
		return 2
	}
	if i < 1000 {
		return 3
	}
	if i < 10000 {
		return 4
	}

	count := 4
	i = i / 10000
	for i > 0 {
		i = i / 10
		count++
	}

	return count
}

// generateStoneMap - Generates a StoneMap according to the input
func generateStoneMap(stones []int) StoneMap {
	ret := make(StoneMap)
	for _, stone := range stones {
		ret[stone] = 1
	}
	return ret
}

// countStoneMap - Sums the stone values
func countStoneMap(stoneMap StoneMap) (ret int) {
	for _, stone := range stoneMap {
		ret += stone
	}
	return ret
}
