package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	safeStrictCount := 0
	safeDampedCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		integers := make([]int, 0, len(parts))
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			check(err)
			integers = append(integers, num)
		}

		safe := checkSlopeSlice(integers)
		fmt.Println("=======")
		fmt.Println(integers)
		fmt.Println("Strict", safe)

		if safe {
			safeStrictCount++
			safeDampedCount++
		} else {
			for i := range integers {
				newIntegers := append(integers[:i:i], integers[i+1:]...)
				safe = checkSlopeSlice(newIntegers)
				if safe {
					safeDampedCount++
					fmt.Println("Damped!")
					break
				}
			}
		}
	}
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Safe reports: ")
	fmt.Println("Part1", safeStrictCount)
	fmt.Println("Part2", safeDampedCount)
}

func checkSlopeSlice(integers []int) bool {
	ascending := integers[0] < integers[1]
	for i := 1; i < len(integers); i++ {
		if !checkSlope(integers[i-1], integers[i], ascending) {
			return false
		}
	}

	return true
}

func checkSlope(prev int, curr int, asc bool) bool {
	return prev != curr && abs(prev-curr) < 4 && (prev < curr) == asc
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
