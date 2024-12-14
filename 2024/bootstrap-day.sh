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

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning input:", err)
    }
}
EOL

# Create input files
touch input.txt test_input.txt

echo "Bootstrapped $DAY_FOLDER with main.go, go.mod, input.txt, and test_input.txt."
