package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Day 3 Part 1")
	process(false)

	fmt.Println("Day 3 Part 2")
	process(true)
}

func process(strip bool) {
	b, err := os.ReadFile("input.txt")
	check(err)
	input := string(b)

	if strip {
		r := regexp.MustCompile(`don't\(\).*?do\(\)`)
		input = r.ReplaceAllString(input, "")
	}

	multiplicationSum := 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		num1, _ := strconv.Atoi(v[1])
		num2, _ := strconv.Atoi(v[2])
		multiplicationSum += num1 * num2
	}

	fmt.Println(multiplicationSum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
