package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Robot struct {
	position Position
	velocity Velocity
}

func (r *Robot) step(width, height int) {
	r.position.x = (((r.position.x + r.velocity.dx) % width) + width) % width
	r.position.y = (((r.position.y + r.velocity.dy) % height) + height) % height
}

type Position struct {
	x, y int
}

type Velocity struct {
	dx, dy int
}

var directions [4]Position = [4]Position{
	{0, -1}, // up
	{0, 1},  //down
	{1, 0},  // right
	{-1, 0}, // left
}

func findFinalPosition(robot Robot, width, height, steps int) Position {
	tmpX := robot.position.x + robot.velocity.dx*steps
	tmpY := robot.position.y + robot.velocity.dy*steps
	finalX := ((tmpX % width) + width) % width
	finalY := ((tmpY % height) + height) % height
	return Position{finalX, finalY}
}

func getQuadrant(p Position, width, height int) int {
	i := 0
	if p.x == width/2 || p.y == height/2 {
		return -1
	}
	if p.x > width/2 {
		i += 1
	}
	if p.y > height/2 {
		i += 2
	}
	return i
}

func solve1(lines []string, width int, height int) (string, error) {
	var answer int
	const steps = 100

	robots := parseLines(lines)
	quadrants := make(map[int]int, 4)
	for _, robot := range robots {
		finalPos := findFinalPosition(robot, width, height, steps)
		quad := getQuadrant(finalPos, width, height)
		// fmt.Println(robot, finalPos, quad)
		if quad != -1 {
			quadrants[quad] += 1
		}

	}
	answer = 1
	for _, v := range quadrants {
		answer *= v
		fmt.Println(v, answer)
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string, width, height int) (string, error) {
	var answer int

	robots := parseLines(lines)
	maxNeighbours := -1
	scanner := bufio.NewScanner(os.Stdin)

	for t := 1; t <= 1_000_000; t++ {
		for i := range robots {
			robots[i].step(width, height)
		}

		n := countNeighbours(robots)

		if n > maxNeighbours {
			maxNeighbours = n
			clearScreen()
			moveCursorHome()

			fmt.Printf("Time: %d\n", t)
			fmt.Printf("n=%d, mn=%d\n", n, maxNeighbours)
			drawFrame(robots, width, height)
			scanner.Scan()
		}
	}

	return fmt.Sprintf("%d", answer), nil
}

func countNeighbours(robots []Robot) int {
	neighbours := 0
	locations := make(map[Position]bool)
	for _, r := range robots {
		locations[r.position] = true
	}
	for p := range locations {
		for _, d := range directions {
			nx := p.x + d.x
			ny := p.y + d.y
			if _, exists := locations[Position{nx, ny}]; exists {
				neighbours++
			}
		}
	}
	return neighbours
}

func drawFrame(robots []Robot, width, height int) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, r := range robots {
		grid[r.position.y][r.position.x] = 'X'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func moveCursorHome() {
	fmt.Print("\033[H")
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to int")
	}
	return i
}

func parseLines(lines []string) []Robot {
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			fmt.Println(line)
			panic("no match")
		}
		r := Robot{
			position: Position{mustAtoi(match[1]), mustAtoi(match[2])},
			velocity: Velocity{mustAtoi(match[3]), mustAtoi(match[4])},
		}
		robots[i] = r
	}
	return robots
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
	answer, _ := solve1(lines, 101, 103)
	duration := time.Since(start)
	fmt.Printf("Puzzle 1 Answer: %s (runtime: %v)\n", answer, duration)

	start = time.Now()
	answer2, _ := solve2(lines, 101, 103)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 Answer: %s (runtime: %v)\n", answer2, duration)
}
