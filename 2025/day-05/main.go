package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Can't convert to int: %v\n", s))
	}
	return i
}

type Interval struct {
	start, end int // inclusive
}

func parseFresh(f []string) []Interval {
	var intervals []Interval
	for _, l := range f {
		before, after, found := strings.Cut(l, "-")
		if !found {
			panic(fmt.Sprintf("Parse error %v\n", l))
		}
		start, end := mustAtoi(before), mustAtoi(after)
		i := Interval{start, end}
		intervals = append(intervals, i)
	}

	// sort intervals by start, then end
	slices.SortFunc(intervals, func(a, b Interval) int {
		if n := cmp.Compare(a.start, b.start); n != 0 {
			return n
		}
		return cmp.Compare(a.end, b.end)
	})

	merged := []Interval{intervals[0]}
	for _, current := range intervals[1:] {
		last := &merged[len(merged)-1]
		if current.start <= last.end { // start of the interval is overlapping or directly adjacent
			if current.end > last.end {
				last.end = current.end // extend the already stored interval to include the current one
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func freshContains(intervals []Interval, x int) bool {
	// intervals must be sorted
	// use binary search to find the first index where start > x
	i := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].start > x
	})
	// i-1 is the candidate interval
	if i == 0 {
		return false
	}
	candidate := intervals[i-1]
	return x >= candidate.start && x <= candidate.end
}

func countFresh(intervals []Interval) int {
	fresh := 0
	for _, i := range intervals {
		fresh += i.end - i.start + 1
	}
	return fresh
}

func parseLines(lines []string) ([]string, []string) {
	var fresh []string
	var avail []string

	reading_fresh := true
	for _, line := range lines {
		if line == "" {
			reading_fresh = false
			continue
		}
		if reading_fresh {
			fresh = append(fresh, line)
		} else {
			avail = append(avail, line)
		}
	}
	return fresh, avail
}

func solve1(lines []string) (string, error) {
	var answer int
	fresh, avail := parseLines(lines)
	intervals := parseFresh(fresh)
	for _, a := range avail {
		if freshContains(intervals, mustAtoi(a)) {
			answer += 1
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	fresh, _ := parseLines(lines)
	intervals := parseFresh(fresh)
	answer = countFresh(intervals)
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
