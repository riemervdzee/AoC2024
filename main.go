package main

import (
	"fmt"
	"riemer/day01"
	"riemer/day02"
	"riemer/day03"
	"riemer/day04"
	"riemer/day05"
	"time"
)

func main() {
	days := []func(){day01.Process, day02.Process, day03.Process, day04.Process, day05.Process}

	for _, function := range days {
		fmt.Println("===================")
		start := time.Now()
		function()
		fmt.Printf("Solved in: %v\n\n", time.Since(start))
	}
}
