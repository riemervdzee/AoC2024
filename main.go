package main

import (
	"fmt"
	"riemer/day14"
	"time"
)

func main() {
	//days := []func(){
	//	day01.Process, day02.Process, day03.Process, day04.Process, day05.Process,
	//	day06.Process, day07.Process, day08.Process, day09.Process, day10.Process,
	//	day11.Process, day12.Process, day13.Process, day14.Process,
	//}
	days := []func(){day14.Process}
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
