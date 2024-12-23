package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
)

func calculateChecksum(data []int) int {
	var checksum int
	for i, v := range data {
		if v != -1 {
			checksum += i * v
		}
	}
	return checksum
}

func solve1(lines []string) (string, error) {
	var answer int

	data := parseLines(lines)

	f_next, f_stop := iter.Pull2(slices.All(data))
	b_next, b_stop := iter.Pull2(slices.Backward(data))

	defer f_stop()
	defer b_stop()

	fi, fv, _ := f_next()
	bi, bv, _ := b_next()
	for {
		for fv != -1 {
			fi, fv, _ = f_next()
		}

		for bv == -1 {
			bi, bv, _ = b_next()
		}

		if fi >= bi {
			break
		}
		data[fi] = bv
		data[bi] = fv

		fi, fv, _ = f_next()
		bi, bv, _ = b_next()
	}
	answer = calculateChecksum(data)
	return fmt.Sprintf("%d", answer), nil
}

func contigSize(data []int, start int) int {
	var size int
	val := data[start]
	for i := start; data[i] == val; i++ {
		size++
	}
	return size
}

func contigSizeBackwards(data []int, end int) int {
	var size int
	val := data[end]
	for i := end; data[i] == val; i-- {
		size++
	}
	return size
}

func moveFile(data []int, src, dest, size int) {
	for i := range size {
		data[dest+i], data[src+i] = data[src+i], data[dest+i]
	}
}

func solve2(lines []string) (string, error) {
	var answer int

	data := parseLines(lines)

	bi := len(data)
	bv := -1
	first := true
	var smallestBv int

	for {
		for bv == -1 {
			bi--
			bv = data[bi]
		}
		if first {
			smallestBv = bv
		}
		if bv == 0 {
			break
		}
		smallestBv = min(smallestBv, bv)
		fileSize := contigSizeBackwards(data, bi)
		fileStart := bi - (fileSize - 1)

		bi = fileStart
		bv = -1

		var holeSize int

		for fi := 0; fi < bi; fi++ {
			if data[fi] == -1 {
				holeSize = contigSize(data, fi)
				if holeSize < fileSize {
					fi += holeSize
					continue
				}
				moveFile(data, fileStart, fi, fileSize)
				break
			}
		}
	}
	answer = calculateChecksum(data)
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
