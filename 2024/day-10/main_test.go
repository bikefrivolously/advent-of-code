package main

import (
	"fmt"
	"testing"
)

func TestSolve1(t *testing.T) {
	lines, err := readFile("test_input.txt", nil)
	if err != nil {
		t.Fatalf("Error reading input file '%s': %v\n", "test_input.txt", err)
	}

	expected := "-1"
	result, _ := solve1(lines)

	if result != expected {
		t.Errorf("solve1() = %s; want %s", result, expected)
	}
}

func TestSolve2(t *testing.T) {
	lines, err := readFile("test_input.txt", nil)
	if err != nil {
		t.Fatalf("Error reading input file '%s': %v\n", "test_input.txt", err)
	}

	expected := "36"
	result, _ := solve2(lines)

	if result != expected {
		t.Errorf("solve2() = %s; want %s", result, expected)
	}
}

func BenchmarkSolve1(b *testing.B) {
	lines, err := readFile("input.txt", nil)
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input.txt", err)
		panic("crash")
	}
	b.ResetTimer()
	for range b.N {
		solve1(lines)
	}
}

func BenchmarkSolve2(b *testing.B) {
	lines, err := readFile("input.txt", nil)
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", "input.txt", err)
		panic("crash")
	}
	b.ResetTimer()
	for range b.N {
		solve2(lines)
	}
}
