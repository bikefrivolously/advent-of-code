package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseReadings(lines []string) ([][]int, error) {
	var readings [][]int
	for _, line := range lines {
		var levels []int
		l := strings.Fields(line)
		for _, num := range l {
			i, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("error converting '%s' to int: '%v'", num, err)
			}
			levels = append(levels, i)
		}
		readings = append(readings, levels)
	}
	return readings, nil
}

func readingIsSafe(reading []int) bool {
	if len(reading) < 2 {
		return true
	}

	prevDiff := 0
	for i := 1; i < len(reading); i++ {
		diff := reading[i] - reading[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if prevDiff*diff < 0 { // changed direction
			return false
		}
		prevDiff = diff
	}
	return true
}

func removeLevel(reading []int, index int) []int {
	newSlice := make([]int, 0, len(reading)-1)
	newSlice = append(newSlice, reading[:index]...)
	newSlice = append(newSlice, reading[index+1:]...)
	return newSlice
}

func solve1(lines []string) (string, error) {
	var answer int
	readings, err := parseReadings(lines)
	if err != nil {
		return "", err
	}

	var numSafe int

	for _, reading := range readings {
		if readingIsSafe(reading) {
			numSafe++
		}
	}
	answer = numSafe

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	readings, err := parseReadings(lines)
	if err != nil {
		return "", err
	}

	var numSafe int

	for _, reading := range readings {
		if readingIsSafe(reading) {
			numSafe++
			continue
		}

		safeWithRemoval := false
		for i := 0; i < len(reading); i++ {
			modifiedReading := removeLevel(reading, i)
			if readingIsSafe(modifiedReading) {
				safeWithRemoval = true
				break
			}
		}
		if safeWithRemoval {
			numSafe++
		}
	}

	answer = numSafe

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
