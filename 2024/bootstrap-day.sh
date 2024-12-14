#!/bin/bash

YEAR=2024
DAY=$(printf "%02d" $1)
DAY_FOLDER="day-$DAY"
MODULE_PATH="github.com/bikefrivolously/advent-of-code/$YEAR/$DAY_FOLDER"

if [ -d "$DAY_FOLDER" ]; then
    echo "Folder $DAY_FOLDER already exists."
    exit 1
fi

# create a new branch for the day's puzzles
git fetch
git switch -c $YEAR-$DAY_FOLDER origin/main

mkdir -p $DAY_FOLDER

# Initialize go.mod
cd $DAY_FOLDER
go mod init $MODULE_PATH

# Create main.go
cat > main.go <<EOL
package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve1(lines []string) (string, error) {
	return "TODO", nil
}

func solve2(lines []string) (string, error) {
	return "TODO", nil
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

EOL

# Create input files
touch input.txt test_input.txt

echo "Bootstrapped $DAY_FOLDER with main.go, go.mod, input.txt, and test_input.txt."
