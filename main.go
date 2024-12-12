package main

import (
	"fmt"
	"riemer/day08"
	"time"
)

func main() {
	//days := []func(){day01.Process, day02.Process, day03.Process, day04.Process, day05.Process, day06.Process, day07.Process}
	days := []func(){day08.Process}
	totalStart := time.Now()

	for _, function := range days {
		fmt.Println("===================")
		start := time.Now()
		function()
		fmt.Printf("Solved in: %v\n\n", time.Since(start))
	}

	fmt.Println()
	fmt.Printf("Total time: %v\n\n", time.Since(totalStart))
}
