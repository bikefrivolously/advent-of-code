package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]Cell

func (g *Grid) pos(p Position) *Cell {
	return &(*g)[p.y][p.x]
}

type Position struct {
	x int
	y int
}

type Direction int

type GuardInfo struct {
	direction Direction
	position  Position
}

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

type Cell struct {
	visited    bool
	obstructed bool
}

func loadGrid(lines []string) (Grid, GuardInfo) {
	height := len(lines)
	width := len(lines[0])
	grid := make(Grid, height)
	guard := GuardInfo{}
	for i := range grid {
		grid[i] = make([]Cell, width)
	}

	for y, line := range lines {
		for x, char := range line {
			fmt.Println(char)
			c := Cell{
				visited:    false,
				obstructed: char == '#',
			}
			grid[y][x] = c

			if char == '.' {
				continue
			}

			switch char {
			case '^':
				guard.direction = Up
			case 'v':
				guard.direction = Down
			case '>':
				guard.direction = Right
			case '<':
				guard.direction = Left
			}
			if !c.obstructed {
				guard.position.x = x
				guard.position.y = y
			}
		}
	}

	grid.pos(guard.position).visited = true
	return grid, guard
}

func solve1(lines []string) (string, error) {
	answer := 1

	grid, guard := loadGrid(lines)
	fmt.Println(grid, guard)
	h := len(grid)
	w := len(grid[0])

	running := true
	var next Position
	var nextDirection Direction
	for running {
		switch guard.direction {
		case Up:
			next = Position{x: guard.position.x, y: guard.position.y - 1}
			nextDirection = Right
		case Down:
			next = Position{x: guard.position.x, y: guard.position.y + 1}
			nextDirection = Left
		case Right:
			next = Position{x: guard.position.x + 1, y: guard.position.y}
			nextDirection = Down
		case Left:
			next = Position{x: guard.position.x - 1, y: guard.position.y}
			nextDirection = Up
		}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			running = false
			continue
		}

		c := grid.pos(next)
		if c.obstructed {
			guard.direction = nextDirection
		} else {
			if !c.visited {
				answer++
			}
			c.visited = true
			guard.position = next
		}
		fmt.Printf("Direction: %s, Position: %v\n", guard.direction, guard.position)
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
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
