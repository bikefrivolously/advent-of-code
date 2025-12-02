package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func parseLines(lines []string) []int {
	var moves []int
	for _, line := range lines {
		dir := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("Unable to parse line")
		}
		if dir == 'L' {
			count *= -1
		}
		moves = append(moves, count)
	}
	return moves
}

func rotateDial(pos int, clicks int) int {
	new_pos := (pos + clicks) % 100
	return new_pos
}

func solve1(lines []string) (string, error) {
	var answer int
	moves := parseLines(lines)
	dial_position := 50
	for _, move := range moves {
		dial_position = rotateDial(dial_position, move)
		if dial_position == 0 {
			answer += 1
		}
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
