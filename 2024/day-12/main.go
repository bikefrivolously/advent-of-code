package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"time"
)

type Grid [][]rune

func (g Grid) AtPosition(p Position) rune {
	return g[p.y][p.x]
}

type Position struct {
	x, y int
}

var directions = []struct {
	dx, dy int
}{
	{0, 1},  // Right
	{0, -1}, // Left
	{1, 0},  // Down
	{-1, 0}, // Up
}

func findNeighbours(g Grid, p Position) []Position {
	h := len(g)
	w := len(g[0])
	neighbours := []Position{}

	for _, d := range directions {
		nx, ny := p.x+d.dx, p.y+d.dy
		if nx >= 0 && nx < w && ny >= 0 && ny < h {
			neighbours = append(neighbours, Position{nx, ny})
		}
	}
	return neighbours
}

func calculatePerimeter(region map[Position]bool) int {
	regionPerim := 0
	for p := range region {
		plotPerim := 4
		for _, d := range directions {
			nx, ny := p.x+d.dx, p.y+d.dy
			if region[Position{nx, ny}] {
				plotPerim--
			}
		}
		regionPerim += plotPerim
	}
	return regionPerim
}

func bfs(g Grid, start Position) map[Position]bool {
	region := map[Position]bool{}
	plotType := g.AtPosition(start)

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		plot := queue.Remove(queue.Front()).(Position)
		if g.AtPosition(plot) != plotType {
			continue
		}

		// add the current plot to the list and mark as visited '.'
		region[plot] = true
		g[plot.y][plot.x] = '.'

		// add all in bounds neighbours to the queue
		for _, neighbour := range findNeighbours(g, plot) {
			queue.PushBack(neighbour)
		}
	}
	return region
}

func solve1(lines []string) (string, error) {
	var answer int

	grid := parseLines(lines)
	for y, row := range grid {
		for x, v := range row {
			if v == '.' {
				continue
			}
			region := bfs(grid, Position{x, y})
			area := len(region)
			perim := calculatePerimeter(region)
			answer += area * perim
			// fmt.Printf("Found region of size %d with perimiter %d\n", area, perim)
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
}

func parseLines(lines []string) Grid {
	grid := make(Grid, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, c := range line {
			grid[y][x] = c
		}
	}
	return grid
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
