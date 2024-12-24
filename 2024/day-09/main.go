package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/bikefrivolously/advent-of-code/2024/day-09/freeheap"
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

	data, _ := parseLines(lines)

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

	data, freeSlots := parseLines(lines)

	bi := len(data)
	bv := -1
	prevMoved := -1

	for {
		bv = -1
		for bv == -1 {
			bi--
			bv = data[bi]
		}
		if bv == 0 {
			break
		}
		fileSize := contigSizeBackwards(data, bi)
		fileStart := bi - (fileSize - 1)
		bi = fileStart

		// make sure we aren't trying to move the same fileID more than once
		if prevMoved != -1 && prevMoved <= bv {
			continue
		}

		var earliestHoleStart int = fileStart
		var holeStart int
		var holeSize int
		foundHole := false
		// start looking for the hole with smallest index that will fit the file
		for s := fileSize; s < 10; s++ {
			if freeSlots[s].Len() > 0 {
				holeStart = freeSlots[s][0]
				if holeStart < earliestHoleStart {
					holeSize = s
					foundHole = true
					earliestHoleStart = holeStart
				}
			}
		}
		if foundHole {
			holeStart = heap.Pop(&freeSlots[holeSize]).(int)
			holeNewSize := holeSize - fileSize
			if holeNewSize > 0 {
				heap.Push(&freeSlots[holeNewSize], holeStart+fileSize)
			}
			moveFile(data, fileStart, holeStart, fileSize)
		}

		prevMoved = bv
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

func parseLines(lines []string) ([]int, [10]freeheap.FreeHeap) {
	var data []int
	var freeSlots [10]freeheap.FreeHeap

	for i, c := range lines[0] {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			// data
			data = slices.AppendSeq(data, emitN(i/2, n))
		} else {
			// free space represented by -1
			index := len(data)
			data = slices.AppendSeq(data, emitN(-1, n))
			freeSlots[n] = append(freeSlots[n], index) // keep track of the starting index for each free chunk of size n
		}
	}

	for i := range freeSlots {
		heap.Init(&freeSlots[i])
	}

	return data, freeSlots
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

	linesHard, err := readFile("input_hard1.txt", nil)
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input_hard1.txt", err)
		os.Exit(1)
	}

	linesEvil, err := readFile("input_evil1.txt", &ReadFileOptions{BufferSize: 200 * 1024})
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input_evil1.txt", err)
		os.Exit(1)
	}

	start := time.Now()
	answer, _ := solve1(lines)
	duration := time.Since(start)
	fmt.Printf("Puzzle 1 Answer: %s (runtime: %v)\n", answer, duration)
	if answer != "6446899523367" {
		fmt.Println("Puzzle 1 - expected 6446899523367")
	}

	start = time.Now()
	answer2, _ := solve2(lines)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 Answer: %s (runtime: %v)\n", answer2, duration)
	if answer2 != "6478232739671" {
		fmt.Println("Puzzle 2 - expected 6478232739671")
	}

	start = time.Now()
	answerHard, _ := solve2(linesHard)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 (Hard) Answer: %s (expected: 97898222299196) (runtime: %v)\n", answerHard, duration)

	start = time.Now()
	answerEvil, _ := solve2(linesEvil)
	duration = time.Since(start)
	fmt.Printf("Puzzle 2 (Evil) Answer: %s (expected: 5799706413896802) (runtime: %v)\n", answerEvil, duration)
}
