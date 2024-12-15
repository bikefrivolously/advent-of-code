package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve1(lines []string) (string, error) {
	var answer int
	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			answer += x * y
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	enabled := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				answer += x * y
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
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
