package main

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	lines, err := readFile("test_input.txt")
	if err != nil {
		t.Fatalf("Error reading input file '%s': %v\n", "test_input.txt", err)
	}

	expected := "18"
	result, _ := solve1(lines)

	if result != expected {
		t.Errorf("solve1() = %s; want %s", result, expected)
	}
}

func TestSolve2(t *testing.T) {
	lines, err := readFile("test_input.txt")
	if err != nil {
		t.Fatalf("Error reading input file '%s': %v\n", "test_input.txt", err)
	}

	expected := "9"
	result, _ := solve2(lines)

	if result != expected {
		t.Errorf("solve2() = %s; want %s", result, expected)
	}
}
