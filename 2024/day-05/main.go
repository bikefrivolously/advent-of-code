package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	rules    map[string]ruleList
	allRules ruleList
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
		r.allRules = append(r.allRules, realRule)
	}
}

func (r *ruleBook) IsOrderValid(before string, after string) bool {
	for _, rule := range r.rules[before] {
		if rule.before == before && rule.after == after {
			return true
		}
		if rule.after == before && rule.before == after {
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

func createGraph(rules ruleList, updates [][]string) (map[string][]string, map[string]int) {
	graph := make(map[string][]string)
	inDegree := make(map[string]int)

	for _, rule := range rules {
		graph[rule.before] = append(graph[rule.before], rule.after)
		inDegree[rule.after]++
		if _, exists := inDegree[rule.before]; !exists {
			inDegree[rule.before] = 0
		}
	}

	// make sure all nodes exist in inDegree
	for _, update := range updates {
		for _, node := range update {
			if _, exists := inDegree[node]; !exists {
				inDegree[node] = 0
			}
		}
	}
	return graph, inDegree
}

func createSubGraph(graph map[string][]string, nodes []string) (map[string][]string, map[string]int) {
	subGraph := make(map[string][]string)
	inDegree := make(map[string]int)

	for _, node := range nodes {
		for _, edge := range graph[node] {
			if idx := slices.Index(nodes, edge); idx >= 0 {
				subGraph[node] = append(subGraph[node], edge)
				inDegree[edge]++
			}
			if _, exists := inDegree[node]; !exists {
				inDegree[node] = 0
			}
		}
	}

	return subGraph, inDegree
}

func reorderList(sortedOrder []string, origList []string) []string {
	orderMap := make(map[string]int)
	for i, v := range sortedOrder {
		orderMap[v] = i
	}

	sortedList := make([]string, len(origList))
	copy(sortedList, origList)
	for i := 0; i < len(sortedList); i++ {
		for j := 0; j < len(sortedList); j++ {
			if orderMap[sortedList[i]] > orderMap[sortedList[j]] {
				sortedList[i], sortedList[j] = sortedList[j], sortedList[i]
			}
		}
	}
	return sortedList
}

func topologicalSort(graph map[string][]string, inDegree map[string]int) ([]string, error) {
	// Kahn's algorithm
	var sorted []string
	var queue []string
	// create a queue with all the nodes that have no incoming edges
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = slices.Delete(queue, 0, 1)
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// check if all nodes are included
	if len(sorted) != len(inDegree) {
		return nil, fmt.Errorf("graph contains a cycle")
	}

	return sorted, nil

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

	rules, updates := parseData(lines)
	book := newRuleBook()
	book.loadRules(rules)

	graph, _ := createGraph(book.allRules, updates)
	for _, update := range updates {
		if checkUpdateValid(book, update) {
			continue
		}
		subGraph, subInDegree := createSubGraph(graph, update)

		sortedNodes, err := topologicalSort(subGraph, subInDegree)
		if err != nil {
			return "", err
		}
		sortedList := reorderList(sortedNodes, update)
		fmt.Println(update)
		fmt.Println(sortedNodes)
		mid, _ := strconv.Atoi(sortedList[len(sortedList)/2])
		answer += mid
	}
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
