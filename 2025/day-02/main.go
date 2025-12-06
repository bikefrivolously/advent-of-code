package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	start int
	end   int
}

func SplitInt(n int) (int, int, bool) {
	digits := int(math.Log10(float64(n))) + 1
	if digits%2 != 0 {
		return n, 0, false
	}
	half := digits / 2
	div := int(math.Pow10(half))
	left := n / div
	right := n % div
	return left, right, true
}

func checkForRepeats(n int) bool {
	s := strconv.Itoa(n)
	max_size := len(s) / 2
	for size := 1; size <= max_size; size++ {

		if len(s)%size != 0 {

			continue
		}
		first_part := s[0:size]

		var repeat_found bool
		var part string
		for j := size; j+size <= len(s); j += size {
			repeat_found = true
			part = s[j : j+size]

			if part != first_part {
				repeat_found = false
				break
			}
		}
		if repeat_found {

			return true
		}
	}

	return false
}

func parseInput(line string) []Range {
	var ranges []Range
	for part := range strings.SplitSeq(line, ",") {
		before, after, found := strings.Cut(part, "-")
		if !found {
			panic(fmt.Sprintf("Error parsing input at part %s\n", part))
		}
		start, err := strconv.Atoi(before)
		if err != nil {
			panic("Could not parse start")
		}
		end, err := strconv.Atoi(after)
		if err != nil {
			panic("Could not parse end")
		}
		r := Range{start, end}
		ranges = append(ranges, r)
	}
	return ranges
}

func solve1(lines []string) (string, error) {
	var answer int
	ranges := parseInput(lines[0])
	for _, r := range ranges {
		for i, j := r.start, r.end; i <= j; i++ {
			left, right, split := SplitInt(i)
			if !split {
				continue
			}
			if left == right {
				answer += i
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	ranges := parseInput(lines[0])
	for _, r := range ranges {
		for i, j := r.start, r.end; i <= j; i++ {
			if checkForRepeats(i) {
				answer += i
			}
		}
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
