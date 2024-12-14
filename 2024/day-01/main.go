package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readListsFromFile(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list_one []int
	var list_two []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		if len(words) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: '%s'", line)
		}

		int1, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting '%s' to int: '%v'", words[0], err)
		}

		int2, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting '%s' to int: '%v'", words[1], err)
		}

		list_one = append(list_one, int1)
		list_two = append(list_two, int2)
	}

	slices.Sort(list_one)
	slices.Sort(list_two)
	return list_one, list_two, nil
}

func main() {
	list_one, list_two, err := readListsFromFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file: '%v'\n", err)
		os.Exit(1)
	}

	var total_distance int
	for i := 0; i < len(list_one); i++ {
		n1 := list_one[i]
		n2 := list_two[i]
		d := n1 - n2
		fmt.Println(n1, n2, d)
		if d < 0 {
			d = -d
		}
		fmt.Println(d)
		total_distance += d
		fmt.Println(total_distance)
	}
	fmt.Printf("The total distance is: %d\n", total_distance)
}
