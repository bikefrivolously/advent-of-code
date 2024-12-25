package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func numDigits(i int) int {
	if i == 0 {
		return 1
	}
	digits := 0
	for i > 0 {
		i /= 10
		digits++
	}
	return digits
}

func intPower(base, exp int) int {
	if exp == 0 {
		return 1
	}

	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func countAfterBlinks(i int, blinks int) int {
	// fmt.Printf("countAfterBlinks: i=%d, blinks=%d\n", i, blinks)
	if blinks == 0 {
		return 1
	}
	if i == 0 {
		return countAfterBlinks(1, blinks-1)
	} else if digits := numDigits(i); digits%2 == 0 {
		// split in two
		l := i / intPower(10, digits/2)
		r := i - (l * intPower(10, digits/2))
		// fmt.Printf("i: %d, l: %d, r: %d\n", i, l, r)

		return countAfterBlinks(l, blinks-1) + countAfterBlinks(r, blinks-1)
	} else {
		return countAfterBlinks(i*2024, blinks-1)
	}
}

func solve1(lines []string) (string, error) {
	var answer int

	queue := parseLines(lines)
	for e := queue.Front(); e != nil; e = e.Next() {
		answer += countAfterBlinks(e.Value.(int), 25)
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	queue := parseLines(lines)
	for e := queue.Front(); e != nil; e = e.Next() {
		answer += countAfterBlinks(e.Value.(int), 75)
	}

	return fmt.Sprintf("%d", answer), nil
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to int")
	}
	return i
}

func parseLines(lines []string) *list.List {
	l := list.New()
	for _, line := range lines {
		nums := strings.Fields(line)
		for _, num := range nums {
			i := mustAtoi(num)
			l.PushBack(i)
		}
	}
	return l
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
