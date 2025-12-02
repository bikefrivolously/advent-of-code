package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

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
	new_pos := (((pos + clicks) % 100) + 100) % 100
	return new_pos
}

func rotateDialCountZeros(pos int, clicks int) (int, int) {
	var zeros int
	zeros += Abs(clicks / 100)
	remainder := clicks % 100
	fmt.Printf("Position: %d, Clicks: %d, Zeros: %d, Remainder: %d\n", pos, clicks, zeros, remainder)
	if remainder == 0 {
		return pos, zeros
	}
	new_pos := rotateDial(pos, remainder)
	if pos == 0 {
		return new_pos, zeros
	}
	if (remainder > 0 && new_pos < pos) || (remainder < 0 && new_pos > pos) || new_pos == 0 {
		zeros += 1
	}
	fmt.Printf("New Pos: %d, Zeros: %d\n", new_pos, zeros)
	return new_pos, zeros
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

	moves := parseLines(lines)
	dial_position := 50
	zeros := 0
	for _, move := range moves {
		dial_position, zeros = rotateDialCountZeros(dial_position, move)
		answer += zeros
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
