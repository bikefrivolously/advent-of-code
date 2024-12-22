package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
)

func solve1(lines []string) (string, error) {
	var answer int

	data := parseLines(lines)

	f_next, f_stop := iter.Pull2(slices.All(data))
	b_next, b_stop := iter.Pull2(slices.Backward(data))

	defer f_stop()
	defer b_stop()

	f_i, f_v, _ := f_next()
	b_i, b_v, _ := b_next()
	for {
		for f_v != -1 {
			f_i, f_v, _ = f_next()
		}

		for b_v == -1 {
			b_i, b_v, _ = b_next()
		}

		if f_i >= b_i {
			break
		}
		// fmt.Printf("before swap fi=%d, fv=%d, bi=%d, bv=%d\n", f_i, f_v, b_i, b_v)
		data[f_i] = b_v
		data[b_i] = f_v
		// fmt.Printf("after swap fi=%d, fv=%d, bi=%d, bv=%d\n", f_i, f_v, b_i, b_v)

		f_i, f_v, _ = f_next()
		b_i, b_v, _ = b_next()
	}

	// fmt.Println(data)

	for i, v := range data {
		if v != -1 {
			answer += i * v
		}
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = -1
	return fmt.Sprintf("%d", answer), nil
}

// an iterator that returns v n times
func emitN(v, n int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for n > 0 {
			if !yield(v) {
				return
			}
			n--
		}
	}
}

func parseLines(lines []string) []int {
	var data []int
	for i, c := range lines[0] {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			// data
			data = slices.AppendSeq(data, emitN(i/2, n))
		} else {
			// free space represented by -1
			data = slices.AppendSeq(data, emitN(-1, n))
		}
	}
	return data
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
