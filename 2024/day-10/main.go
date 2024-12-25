package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

type QueueElement struct {
	pos             Position
	targetElevation int
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to int")
	}
	return i
}

func parseLines(lines []string) Grid {
	grid := make(Grid, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, c := range line {
			p := Position{x, y}
			grid.Set(p, mustAtoi(string(c)))
		}
	}
	return grid
}

func findTrailheads(grid Grid) []Position {
	var th []Position

	for y, row := range grid {
		for x, pos := range row {
			if pos == 0 {
				th = append(th, Position{x, y})
			}
		}
	}
	return th
}

func nextPos(g Grid, p Position) []Position {
	next := make([]Position, 0, 4)
	h := len(g)
	w := len(g[0])

	if p.x-1 >= 0 {
		next = append(next, Position{p.x - 1, p.y})
	}
	if p.x+1 < w {
		next = append(next, Position{p.x + 1, p.y})
	}
	if p.y-1 >= 0 {
		next = append(next, Position{p.x, p.y - 1})
	}
	if p.y+1 < h {
		next = append(next, Position{p.x, p.y + 1})
	}
	return next
}

func bfs(g Grid, start Position, target int, countDistinct bool) int {
	targetsFound := make(map[Position]int)
	queue := list.New()
	queue.PushBack(QueueElement{start, 0})
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(QueueElement)
		if g.Pos(node.pos) != node.targetElevation {
			continue
		}
		if g.Pos(node.pos) == target {
			targetsFound[node.pos] += 1
		}
		neighbours := nextPos(g, node.pos)
		for _, neighbour := range neighbours {
			queue.PushBack(QueueElement{neighbour, node.targetElevation + 1})
		}
	}
	if countDistinct {
		distinct := 0
		for _, v := range targetsFound {
			distinct += v
		}
		return distinct

	} else {
		return len(targetsFound)
	}
}

func solve1(lines []string) (string, error) {
	var answer int
	grid := parseLines(lines)
	trailheads := findTrailheads(grid)
	for _, th := range trailheads {
		answer += bfs(grid, th, 9, false)
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int
	grid := parseLines(lines)
	trailheads := findTrailheads(grid)
	for _, th := range trailheads {
		answer += bfs(grid, th, 9, true)
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
