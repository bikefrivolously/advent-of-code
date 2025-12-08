package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// returns the neighbour coords for pos r, c for a grid rowsxcols in size
// no wrap-around
func Neighbours8(r, c, rows, cols int) [][2]int {
	res := make([][2]int, 0, 8)
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				res = append(res, [2]int{nr, nc})
			}
		}
	}
	return res
}

type Grid [][]int

func (g Grid) Pos(p Position) int {
	return g[p.y][p.x]
}

func (g Grid) Set(p Position, v int) {
	g[p.y][p.x] = v
}

type Position struct {
	x, y int
}

func parseLines(lines []string) Grid {
	grid := make(Grid, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, c := range line {
			p := Position{x, y}
			if c == '@' {
				grid.Set(p, 1)
			} else {
				grid.Set(p, 0)
			}
		}
	}
	return grid
}

func positionAccessible(grid Grid, x, y int) bool {
	rows := len(grid)
	cols := len(grid[0])
	n := Neighbours8(y, x, rows, cols)
	num_filled := 0
	for _, p := range n {
		ny := p[0]
		nx := p[1]
		contents := grid[ny][nx]
		num_filled += contents
	}
	if num_filled < 4 {
		return true
	}
	return false
}

func solve1(lines []string) (string, error) {
	var answer int
	grid := parseLines(lines)
	for y, row := range grid {
		for x, val := range row {
			if val != 1 {
				continue
			}
			if positionAccessible(grid, x, y) {
				answer += 1
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	grid := parseLines(lines)
	for true {
		var accessible []Position
		for y, row := range grid {
			for x, val := range row {
				if val != 1 {
					continue
				}
				if positionAccessible(grid, x, y) {
					answer += 1
					p := Position{x, y}
					accessible = append(accessible, p)
				}
			}
		}
		if len(accessible) == 0 {
			break
		}
		for _, p := range accessible {
			grid.Set(p, 0)
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
