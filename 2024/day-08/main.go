package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type FreqPoints map[string][]Point

func calcRiseRun(p1, p2 Point) (int, int) {
	rise := p2.y - p1.y
	run := p2.x - p1.x
	return rise, run
}

func isInBounds(bounds, p Point) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	if p.x > bounds.x || p.y > bounds.y {
		return false
	}
	return true
}

func solve1(lines []string) (string, error) {
	var answer int

	antinodes := map[Point]bool{}

	freqs, bounds := parseLines(lines)
	for freq, points := range freqs {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				p1 := points[i]
				p2 := points[j]
				rise, run := calcRiseRun(p1, p2)
				fmt.Printf("%s: point1=%v, point2=%v, rise=%d, run=%d\n", freq, p1, p2, rise, run)
				// from p1 sub rise and run
				// from p2 add rise and run
				a1 := Point{x: p1.x - run, y: p1.y - rise}
				a2 := Point{x: p2.x + run, y: p2.y + rise}
				fmt.Printf("a1=%v, a2=%v\n", a1, a2)
				if isInBounds(bounds, a1) {
					antinodes[a1] = true
				}
				if isInBounds(bounds, a2) {
					antinodes[a2] = true
				}
			}
		}
	}
	answer = len(antinodes)

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
}

func parseLines(lines []string) (FreqPoints, Point) {
	freqs := make(FreqPoints)
	max_y := len(lines) - 1
	max_x := len(lines[0]) - 1
	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}
			k := string(c)
			freqs[k] = append(freqs[k], Point{x: x, y: y})
		}
	}
	return freqs, Point{x: max_x, y: max_y}
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
