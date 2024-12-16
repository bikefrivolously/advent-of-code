package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before string
	after  string
}

type ruleList []Rule

// holds rules, and can be used to look up all rules involving a number
type ruleBook struct {
	rules map[string]ruleList
}

func newRuleBook() *ruleBook {
	return &ruleBook{
		rules: make(map[string]ruleList),
	}
}

func (r *ruleBook) loadRules(rules []string) {
	for _, rule := range rules {
		operands := strings.SplitN(rule, "|", 2)
		realRule := Rule{before: operands[0], after: operands[1]}
		r.rules[operands[0]] = append(r.rules[operands[0]], realRule)
		r.rules[operands[1]] = append(r.rules[operands[1]], realRule)
	}
}

func (r *ruleBook) IsOrderValid(before string, after string) bool {
	for _, rule := range r.rules[before] {
		fmt.Printf("Checking if %s belongs before %s. Rule: before=%s, after=%s\n", before, after, rule.before, rule.after)
		if rule.before == before && rule.after == after {
			fmt.Println("True")
			return true
		}
		if rule.after == before && rule.before == after {
			fmt.Println("False")
			return false
		}
	}
	return true
}

// check all pairs of pages against each rule
func checkUpdateValid(book *ruleBook, pages []string) bool {
	for i := 0; i < len(pages)-1; i++ {
		p1 := pages[i]
		for j := i + 1; j < len(pages); j++ {
			p2 := pages[j]
			if !book.IsOrderValid(p1, p2) {
				return false
			}
		}
	}
	return true
}

func parseData(lines []string) ([]string, [][]string) {
	var rules []string
	var updates [][]string

	readingRules := true
	for _, line := range lines {
		if readingRules {
			if line == "" {
				readingRules = false
			} else {
				rules = append(rules, line)
			}
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	return rules, updates
}

func solve1(lines []string) (string, error) {
	var answer int

	rules, updates := parseData(lines)
	book := newRuleBook()
	book.loadRules(rules)
	for _, update := range updates {
		if checkUpdateValid(book, update) {
			mid, _ := strconv.Atoi(update[len(update)/2])
			answer += mid
		}
	}

	return fmt.Sprintf("%d", answer), nil
}

func solve2(lines []string) (string, error) {
	var answer int

	answer = 143
	return fmt.Sprintf("%d", answer), nil
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
