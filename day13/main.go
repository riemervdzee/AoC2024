package day13

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"regexp"
	"riemer/utils"
	"sync"
)

type ClawMachine struct {
	ButtonA utils.Vector
	ButtonB utils.Vector
	Prize   utils.Vector
}

type Result struct {
	Tokens   int
	Possible bool
}

var regexNumbers = regexp.MustCompile(`X[+=](\d+), Y[+=](\d+)`)

func Process() {
	lines := utils.ReadFile("day13/input.txt")
	var machines []ClawMachine
	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, ClawMachine{
			ButtonA: parseLine(lines[i]),
			ButtonB: parseLine(lines[i+1]),
			Prize:   parseLine(lines[i+2]),
		})
	}

	// Start a goroutine for every machine
	results := make(chan [2]Result, len(machines))
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)
		go solveMachineConcurrent(machine, results, &wg)
	}

	// Start a goroutine to wait for all results and close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Stream results and process them till it gets closed
	totalTokens1 := 0
	totalTokens2 := 0
	for result := range results {
		if result[0].Possible {
			totalTokens1 += result[0].Tokens
		}
		if result[1].Possible {
			totalTokens2 += result[1].Tokens
		}
	}

	fmt.Println("Day 13 Results")
	fmt.Println("Part1", totalTokens1)
	fmt.Println("Part2", totalTokens2)
}

// solveMachineConcurrent solves both parts parallel for one machine
func solveMachineConcurrent(machine ClawMachine, results chan<- [2]Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Start goroutines for both parts
	part1Chan := make(chan Result, 1)
	go func() {
		tokens, possible := solveMachine(machine, false)
		part1Chan <- Result{tokens, possible}
	}()

	part2Chan := make(chan Result, 1)
	go func() {
		tokens, possible := solveMachine(machine, true)
		part2Chan <- Result{tokens, possible}
	}()

	// Return results
	results <- [2]Result{<-part1Chan, <-part2Chan}
}

// solveMachine try to find a solution for a machine
func solveMachine(machine ClawMachine, part2 bool) (int, bool) {
	a, b, err := solveSystem(machine, part2)
	if err != nil || !isValidSolution(a, b, part2) {
		return 0, false
	}

	tokens := int(math.Round(a))*3 + int(math.Round(b))*1
	return tokens, true
}

// parseLine parses a single line to a utils.Vector
func parseLine(line string) utils.Vector {
	matches := regexNumbers.FindStringSubmatch(line)
	return utils.Vector{utils.StringToInt(matches[1]), utils.StringToInt(matches[2])}
}

// solveSystem tries to find a solution and solve the ClawMachine as a linear equation matrix
func solveSystem(machine ClawMachine, part2 bool) (a, b float64, err error) {
	if part2 {
		machine.Prize[0] += 10000000000000
		machine.Prize[1] += 10000000000000
	}

	A := mat.NewDense(2, 2, []float64{
		float64(machine.ButtonA[0]), float64(machine.ButtonB[0]),
		float64(machine.ButtonA[1]), float64(machine.ButtonB[1]),
	})
	B := mat.NewVecDense(2, []float64{
		float64(machine.Prize[0]),
		float64(machine.Prize[1]),
	})

	// Solve
	x := mat.NewVecDense(2, nil)
	err = x.SolveVec(A, B)
	if err != nil {
		return 0, 0, fmt.Errorf("no solution found: %v", err)
	}

	return x.AtVec(0), x.AtVec(1), nil
}

// isValidSolution checks if the solution is according to the rules
func isValidSolution(a, b float64, part2 bool) bool {
	// Check limits for part 1
	if !part2 && (a > 100 || b > 100) {
		return false
	}

	// Check if the numbers are integers
	tolerance := 1e-4
	if math.Abs(a-math.Round(a)) > tolerance || math.Abs(b-math.Round(b)) > tolerance {
		return false
	}

	return true
}
