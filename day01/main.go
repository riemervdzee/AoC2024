package day01

import (
	"fmt"
	"riemer/utils"
	"sort"
	"strings"
)

func Process() {
	var list1 = make([]int, 0, 1000)
	var list2 = make([]int, 0, 1000)

	// Read file input and fill list1 and list2
	fileLines := utils.ReadFile("day01/input.txt")
	for _, line := range fileLines {
		parts := strings.Fields(line)
		i1 := utils.StringToInt(parts[0])
		i2 := utils.StringToInt(parts[1])
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	// Loop through list and get sum of the distance and similarities
	distanceSum := 0
	similaritySum := 0
	for i := 0; i < len(list1); i++ {
		distanceSum += utils.Abs(list1[i] - list2[i])
		similaritySum += countOccurrences(list2, list1[i]) * list1[i]
	}

	// Answers for Day 1
	fmt.Println("Day 1 Results")
	fmt.Println("Distance", distanceSum)
	fmt.Println("Similarity", similaritySum)
}

// Returns the amount of occurrences of `target` in `slice`
func countOccurrences(slice []int, target int) int {
	first := findFirst(slice, target)
	if first == -1 {
		return 0
	}
	last := findLast(slice, target, first)
	return last - first + 1
}

// findFirst - Uses binary search to get the first occurrence of `target` in `slice`
func findFirst(slice []int, target int) int {
	low, high := 0, len(slice)-1
	result := -1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			result = mid
			high = mid - 1
		} else if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

// findLast - Find the last occurrence of `target` in `slice`. Uses simple iteration
func findLast(slice []int, target int, start int) int {
	for i := start; i < len(slice); i++ {
		if slice[i] != target {
			return i - 1
		}
	}
	return len(slice) - 1
}
