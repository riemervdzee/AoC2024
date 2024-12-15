package day11

import (
	"fmt"
	"riemer/utils"
	"strconv"
	"strings"
)

func Process() {
	lines := utils.ReadFile("day11/input.txt")
	blinkTimesP1 := 25
	//blinkTimesP2 := 75
	stones := utils.StringArrayToIntArray(strings.Fields(lines[0]))

	stones = blinkTimes(stones, blinkTimesP1)
	fmt.Println("Day 11 Results")
	fmt.Println("Part1", len(stones))

	//blinkTimes(stones, blinkTimesP2-blinkTimesP1)
	//fmt.Println("Part2", len(stones))
}

func blinkTimes(stones []int, amount int) []int {
	for i := 0; i < amount; i++ {
		stones = blink(stones)
	}
	return stones
}

func blink(stones []int) []int {
	for i, stone := range stones {
		newStones := processStone(stone)
		stones[i] = newStones[0]
		if len(newStones) > 1 {
			stones = append(stones, newStones[1])
		}
	}
	return stones
}

func processStone(num int) []int {
	if num == 0 {
		return []int{1}
	}
	if countDigits(num)%2 == 0 {
		str := strconv.Itoa(num)
		left := str[0 : len(str)/2]
		right := str[len(str)/2:]
		return []int{utils.StringToInt(left), utils.StringToInt(right)}
	}
	return []int{num * 2024}
}

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
	if i < 100000 {
		return 5
	}

	count := 6
	i = i / 1000000

	for i > 0 {
		i = i / 10
		count++
	}

	return count
}
