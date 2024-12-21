package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Test struct {
	value    int
	operands []int
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func runTest(test Test, operators []func(int, int) int) bool {
	fmt.Println("runTest: ", test, operators)
	// base case
	if len(test.operands) == 1 {
		return test.operands[0] == test.value
	}

	a, b := test.operands[0], test.operands[1]

	for _, op := range operators {
		result := op(a, b)
		newOperands := make([]int, 0, len(test.operands)-1) // set initial length 0
		newOperands = append(newOperands, result)
		newOperands = append(newOperands, test.operands[2:]...)
		t := Test{value: test.value, operands: newOperands}
		if runTest(t, operators) {
			return true
		}
	}
	return false
}

func solve1(lines []string) (string, error) {
	var answer int

	tests, err := parseLines(lines)
	if err != nil {
		return "", err
	}

	var operators []func(int, int) int
	operators = append(operators, add)
	operators = append(operators, mul)

	for _, test := range tests {
		if runTest(test, operators) {
			answer += test.value
		}
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
}

func parseLines(lines []string) ([]Test, error) {
	var tests []Test

	for _, line := range lines {
		s := strings.SplitN(line, ": ", 2)
		val, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, err
		}
		var operands []int
		for _, o := range strings.Fields(s[1]) {
			o, err := strconv.Atoi(o)
			if err != nil {
				return nil, err
			}
			operands = append(operands, o)
		}
		tests = append(tests, Test{value: val, operands: operands})
	}
	return tests, nil
}

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
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
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input.txt", err)
		os.Exit(1)
	}
	answer, _ := solve1(lines)
	fmt.Printf("Puzzle 1 Answer: %s\n", answer)

	answer2, _ := solve2(lines)
	fmt.Printf("Puzzle 2 Answer: %s\n", answer2)
}
