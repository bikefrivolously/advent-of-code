package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type EquationPair struct {
	one LinearEquation
	two LinearEquation
}

type LinearEquation struct {
	aCoefficent int
	bCoefficent int
	constant    int
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to int")
	}
	return i
}

func solveLinearSystem(eq1 LinearEquation, eq2 LinearEquation, limit int) (int, int, error) {
	for b := limit; b >= 0; b-- {
		for a := 0; a <= limit; a++ {
			// ans1 := a*eq1.aCoefficent + b*eq1.bCoefficent
			// ans2 := b*eq2.aCoefficent + b*eq2.bCoefficent

			if a*eq1.aCoefficent+b*eq1.bCoefficent == eq1.constant && a*eq2.aCoefficent+b*eq2.bCoefficent == eq2.constant {
				fmt.Printf("Solved: %d, %d\n", a, b)
				return a, b, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no solution for %v %v", eq1, eq2)
}

func solve1(lines []string) (string, error) {
	var answer int

	equations := parseLines(lines)
	for _, system := range equations {
		a, b, err := solveLinearSystem(system.one, system.two, 100)
		if err != nil {
			continue
		}
		answer += 3*a + 1*b
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
}

type ReadFileOptions struct {
	BufferSize int
}

func parseLines(lines []string) []EquationPair {
	reA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	reB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var equations []EquationPair
	var currentPair *EquationPair

	for _, line := range lines {
		if match := reA.FindStringSubmatch(line); match != nil {
			currentPair = &EquationPair{}
			(*currentPair).one.aCoefficent = mustAtoi(match[1])
			(*currentPair).two.aCoefficent = mustAtoi(match[2])
		} else if match := reB.FindStringSubmatch(line); match != nil {
			(*currentPair).one.bCoefficent = mustAtoi(match[1])
			(*currentPair).two.bCoefficent = mustAtoi(match[2])
		} else if match := rePrize.FindStringSubmatch(line); match != nil {
			(*currentPair).one.constant = mustAtoi(match[1])
			(*currentPair).two.constant = mustAtoi(match[2])
			equations = append(equations, *currentPair)
		}
	}
	return equations
}

func readFile(filePath string, options *ReadFileOptions) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	if options != nil && options.BufferSize > 0 {
		buf := make([]byte, options.BufferSize)
		scanner.Buffer(buf, options.BufferSize)
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func main() {
	lines, err := readFile("input.txt", nil)
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input.txt", err)
		os.Exit(1)
	}
	start := time.Now()
	answer, _ := solve1(lines)
	duration := time.Since(start)
	fmt.Printf("Puzzle 1 Answer: %s (runtime: %v)\n", answer, duration)

	start = time.Now()
	answer2, _ := solve2(lines)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 Answer: %s (runtime: %v)\n", answer2, duration)
}
