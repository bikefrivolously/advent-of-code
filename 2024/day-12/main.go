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

type Edge struct {
	p1, p2 Position
}

func NewEdge(x1, y1, x2, y2 int) Edge {
	return Edge{
		Position{x1, y1},
		Position{x2, y2},
	}
}

var directions = []struct {
	dx, dy int
}{
	{0, 1},  // Right
	{0, -1}, // Left
	{1, 0},  // Down
	{-1, 0}, // Up
}

type Direction struct {
	dx, dy int
}

func findDir(p1, p2 Position) Direction {
	return Direction{
		p2.x - p1.x,
		p2.y - p1.y,
	}
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

func findBoundaryEdges(region map[Position]bool) map[Edge]bool {
	// each position is just an x, y
	// it may have up to four boundary edges

	// top
	// x, y   -> x+1, y
	// bottom
	// x, y+1 -> x+1, y+1
	// left
	// x, y   -> x, y+1
	// right
	// x+1, y -> x+1, y+1

	edges := map[Edge]bool{}
	for p := range region {
		// top
		if !region[Position{p.x, p.y - 1}] {
			e := NewEdge(p.x, p.y, p.x+1, p.y)
			edges[e] = true
		}
		// bottom
		if !region[Position{p.x, p.y + 1}] {
			e := NewEdge(p.x, p.y+1, p.x+1, p.y+1)
			edges[e] = true
		}
		// left
		if !region[Position{p.x - 1, p.y}] {
			e := NewEdge(p.x, p.y, p.x, p.y+1)
			edges[e] = true
		}
		// right
		if !region[Position{p.x + 1, p.y}] {
			e := NewEdge(p.x+1, p.y, p.x+1, p.y+1)
			edges[e] = true
		}
	}
	return edges
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
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	grid := parseLines(lines)
	for y, row := range grid {
		for x, v := range row {
			if v == '.' {
				continue
			}
			region := bfs(grid, Position{x, y})
			area := len(region)
			boundaryEdges := findBoundaryEdges(region)

			adj := map[Position][]Edge{}
			for e := range boundaryEdges {
				adj[e.p1] = append(adj[e.p1], e)
				adj[e.p2] = append(adj[e.p2], e)
			}

			visited := map[Edge]bool{}
			totalSides := 0
			for e0 := range boundaryEdges {
				if visited[e0] {
					continue
				}

				// start walking from e0.p1
				var sideCount int
				var stepDirections []Direction

				// consider e0 in a directed manner, from p1 to p2
				currPoint := e0.p2
				currDir := findDir(e0.p1, e0.p2)
				stepDirections = append(stepDirections, currDir)
				visited[e0] = true

				for {
					var nextEdge Edge
					var nextPoint Position
					var nextDir Direction

					// look for an unvisited edge that starts from our current point
					for _, edge := range adj[currPoint] {
						if visited[edge] {
							continue
						}

						// edges may be stored in either direction
						// set nextPoint to the point that isn't the current location
						if edge.p1 == currPoint {
							nextPoint = edge.p2
						} else {
							nextPoint = edge.p1
						}
						// check the direction of the edge we've selected
						// if there are only 2 edges extending from the current point
						// then use the edge
						// if there are more than 2 edges, we don't want to go straight
						// only select the edge if it means we turn a corner
						nextDir = findDir(currPoint, nextPoint)
						if len(adj[currPoint]) == 2 || nextDir != currDir {
							nextEdge = edge
							break
						}
					}

					stepDirections = append(stepDirections, nextDir)
					visited[nextEdge] = true
					currPoint = nextPoint
					currDir = nextDir

					if currPoint == e0.p1 {
						// loop completed
						break
					}
				}
				// calculate how many times we change direction
				// this is the number of sides
				for i := 1; i < len(stepDirections); i++ {
					if stepDirections[i-1] != stepDirections[i] {
						sideCount++
					}
				}
				if stepDirections[0] != stepDirections[len(stepDirections)-1] {
					sideCount++
				}
				totalSides += sideCount
			}
			answer += area * totalSides
		}
	}
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
