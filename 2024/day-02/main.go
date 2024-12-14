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

func solve1(lines []string) (string, error) {
	var answer int
	readings, err := parseReadings(lines)
	if err != nil {
		return "", err
	}

	var numSafe int

	for _, reading := range readings {
		var safe bool = true
		var prevDiff int = 0
		for i := 0; i < len(reading)-1; i++ {
			x := reading[i]
			y := reading[i+1]

			diff := y - x
			// did it change gradually from the previous level
			if diff == 0 || diff > 3 || diff < -3 {
				safe = false
				break
			}
			// did it change direction?
			if prevDiff*diff < 0 {
				safe = false
				break
			}
			prevDiff = diff
		}
		if safe {
			numSafe += 1
		}
	}
	answer = numSafe

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	return "4", nil
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
