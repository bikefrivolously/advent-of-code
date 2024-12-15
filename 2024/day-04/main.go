package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFromGrid(lines []string, x int, y int) string {
	return string(lines[y][x])
}

func stringsFromGrid(lines []string, x int, y int) []string {
	const stringLen = 4
	var s []string
	// N
	if y-stringLen >= -1 {
		var tmp string
		for j := y; j > y-stringLen; j-- {
			tmp += string(lines[j][x])
		}
		s = append(s, tmp)
	}

	// E
	if x+stringLen <= len(lines[y]) {
		var tmp string
		for i := x; i < x+stringLen; i++ {
			tmp += string(lines[y][i])
		}
		s = append(s, tmp)
	}

	// S
	if y+stringLen <= len(lines) {
		var tmp string
		for j := y; j < y+stringLen; j++ {
			tmp += string(lines[j][x])
		}
		s = append(s, tmp)
	}

	// W
	if x-stringLen >= -1 {
		var tmp string
		for i := x; i > x-stringLen; i-- {
			tmp += string(lines[y][i])
		}
		s = append(s, tmp)
	}

	// NE
	if y-stringLen >= -1 && x+stringLen <= len(lines[y]) {
		var tmp string
		for i := 0; i < stringLen; i++ {
			tmp += string(lines[y-i][x+i])
		}
		s = append(s, tmp)
	}

	// SE
	if y+stringLen <= len(lines) && x+stringLen <= len(lines[y]) {
		var tmp string
		for i := 0; i < stringLen; i++ {
			tmp += string(lines[y+i][x+i])
		}
		s = append(s, tmp)
	}

	// SW
	if y+stringLen <= len(lines) && x-stringLen >= -1 {
		var tmp string
		for i := 0; i < stringLen; i++ {
			tmp += string(lines[y+i][x-i])
		}
		s = append(s, tmp)
	}

	// NW
	if y-stringLen >= -1 && x-stringLen >= -1 {
		var tmp string
		for i := 0; i < stringLen; i++ {
			tmp += string(lines[y-i][x-i])
		}
		s = append(s, tmp)
	}
	return s
}

func xFromCenterGrid(lines []string, x int, y int) []string {
	var s []string
	var tmp string
	// starting from center A, we need room to go 1 in each direction without exceeding the bounds
	if y-1 < 0 || y+1 >= len(lines) || x-1 < 0 || x+1 >= len(lines[y]) {
		return nil
	}
	//NE
	tmp = ""
	for i, j := x-1, y+1; i <= x+1 && j >= y-1; i, j = i+1, j-1 {
		tmp += string(lines[j][i])
	}
	s = append(s, tmp)

	//SE
	tmp = ""
	for i, j := x-1, y-1; i <= x+1 && j <= y+1; i, j = i+1, j+1 {
		tmp += string(lines[j][i])
	}
	s = append(s, tmp)
	return s
}

func solve1(lines []string) (string, error) {
	var answer int
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			c := getFromGrid(lines, x, y)
			if c != "X" {
				continue
			}
			options := stringsFromGrid(lines, x, y)
			for _, word := range options {
				if word == "XMAS" {
					answer += 1
				}
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			c := getFromGrid(lines, x, y)
			if c != "A" {
				continue
			}
			word := xFromCenterGrid(lines, x, y)
			if word == nil {
				continue
			}
			if (word[0] == "MAS" || word[0] == "SAM") && (word[1] == "MAS" || word[1] == "SAM") {
				answer++
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
