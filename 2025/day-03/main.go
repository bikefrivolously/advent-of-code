package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func findMax(bank string, start int) (int, int) {
	max := 0
	pos := start
	for i := start; i < len(bank); i++ {
		v := int(bank[i] - '0')
		if v > max {
			max = v
			pos = i
		}
	}
	return max, pos
}

func findMaxJolt(bank string) int {
	// find the largest digit that isn't at the very end of bank
	left, pos := findMax(bank[:len(bank)-1], 0)
	// find the next largest digit that comes after left
	right, _ := findMax(bank, pos+1)
	return left*10 + right
}

func findMaxBigJolt(bank string) int {
	// now we have to find 12 digits
	var val int
	var d int
	padding := 12 - 1
	pos := -1
	for padding >= 0 {
		d, pos = findMax(bank[:len(bank)-padding], pos+1)
		val += d * int(math.Pow10(padding))
		padding--
	}
	return val
}

func solve1(lines []string) (string, error) {
	var answer int
	for _, line := range lines {
		answer += findMaxJolt(line)
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	for _, line := range lines {
		answer += findMaxBigJolt(line)
	}

	return fmt.Sprintf("%d", answer), nil
}

type ReadFileOptions struct {
	BufferSize int
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
