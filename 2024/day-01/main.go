package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve1(lines []string) (string, error) {
	var list_one []int
	var list_two []int

	for _, line := range lines {
		words := strings.Fields(line)

		if len(words) != 2 {
			return "", fmt.Errorf("invalid line format: '%s'", line)
		}

		int1, err := strconv.Atoi(words[0])
		if err != nil {
			return "", fmt.Errorf("error converting '%s' to int: '%v'", words[0], err)
		}

		int2, err := strconv.Atoi(words[1])
		if err != nil {
			return "", fmt.Errorf("error converting '%s' to int: '%v'", words[1], err)
		}

		list_one = append(list_one, int1)
		list_two = append(list_two, int2)
	}

	slices.Sort(list_one)
	slices.Sort(list_two)

	var total_distance int
	for i := 0; i < len(list_one); i++ {
		n1 := list_one[i]
		n2 := list_two[i]
		d := n1 - n2
		if d < 0 {
			d = -d
		}
		total_distance += d
	}

	return fmt.Sprintf("%d", total_distance), nil
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
}
