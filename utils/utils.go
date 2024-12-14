package utils

import (
	"bufio"
	"os"
	"strconv"
)

// ReadFile - Reads the file as lines
func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// StringArrayToIntArray - Converts a []string to []int
func StringArrayToIntArray(stringArray []string) []int {
	integers := make([]int, 0, len(stringArray))
	for _, part := range stringArray {
		num, err := strconv.Atoi(part)
		Check(err)
		integers = append(integers, num)
	}
	return integers
}

// Abs - simple Abs function for int types
func Abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

// Check - Simple error-checker function which panics when something went wrong
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
