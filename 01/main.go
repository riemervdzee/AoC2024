package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	var list1 = make([]int, 0, 1000)
	var list2 = make([]int, 0, 1000)

	// Read file input and fill list1 and list2
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ph [2]int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &ph[0], &ph[1])
		check(err)
		list1 = append(list1, ph[0])
		list2 = append(list2, ph[1])
	}
	sort.Ints(list1)
	sort.Ints(list2)

	// Loop through list and get sum of the distance and similarities
	var distanceSum uint64 = 0
	var similaritySum uint64 = 0
	for i := 0; i < len(list1); i++ {
		distanceSum += uint64(abs(list1[i] - list2[i]))
		similaritySum += uint64(countOccurrences(list2, list1[i]) * list1[i])
	}

	// Answers for Day 1
	fmt.Println("Distance", distanceSum)
	fmt.Println("Similarity", similaritySum)
}

func countOccurrences(slice []int, target int) int {
	first := findFirst(slice, target)
	if first == -1 {
		return 0
	}
	last := findLast(slice, target, first)
	return last - first + 1
}

func findFirst(slice []int, target int) int {
	// binary search
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

func findLast(slice []int, target int, start int) int {
	for i := start; i < len(slice); i++ {
		if slice[i] != target {
			return i - 1
		}
	}
	return len(slice) - 1
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
