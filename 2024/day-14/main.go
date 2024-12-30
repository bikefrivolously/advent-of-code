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

type Position struct {
	x, y int
}

type Velocity struct {
	dx, dy int
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
		fmt.Println(robot, finalPos, quad)
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

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
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
	answer2, _ := solve2(lines)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 Answer: %s (runtime: %v)\n", answer2, duration)
}
