package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 5 - Part 1")
	fmt.Println("Sum:", calculateMiddleSum(getInput()))
}

func calculateMiddleSum(rules [][]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		relevantRules := common.Filter(rules, func(rule []int) bool {
			return common.Contains(update, rule[0]) && common.Contains(update, rule[1])
		})
		graph := buildGraph(relevantRules)

		if isValidOrder(graph, update) {
			middleIndex := len(update) / 2
			sum += update[middleIndex]
		}
	}

	return sum
}

func isValidOrder(graph map[int][]int, update []int) bool {
	inDegree := calculateInDegree(graph)
	topoOrder := getTopoOrder(graph, inDegree)
	return checkOrder(topoOrder, update)
}

func getTopoOrder(graph map[int][]int, inDegree map[int]int) []int {
	var queue []int

	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var topoOrder []int
	for len(queue) > 0 {
		var node int
		node, queue = queue[0], queue[1:]
		topoOrder = append(topoOrder, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return topoOrder
}

func checkOrder(topoOrder []int, updateOrder []int) bool {
	indexMap := make(map[int]int)
	for idx, num := range topoOrder {
		indexMap[num] = idx
	}

	for i := 0; i < len(updateOrder)-1; i++ {
		if indexMap[updateOrder[i]] > indexMap[updateOrder[i+1]] {
			return false
		}
	}
	return true
}

func calculateInDegree(graph map[int][]int) map[int]int {
	inDegree := make(map[int]int)

	for node := range graph {
		inDegree[node] = 0
	}

	for _, adj := range graph {
		for _, neighbor := range adj {
			inDegree[neighbor]++
		}
	}

	return inDegree
}

func buildGraph(rules [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, rule := range rules {
		from, to := rule[0], rule[1]
		graph[from] = append(graph[from], to)
	}

	return graph
}

func getInput() ([][]int, [][]int) {
	file, err := os.Open("./day05/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules [][]int
	var updates [][]int
	addingUpdates := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			addingUpdates = true
		} else if !addingUpdates {
			pageRules := common.ConvertStringsToInts(strings.Split(line, "|"))
			rules = append(rules, pageRules)
		} else if addingUpdates {
			pages := common.ConvertStringsToInts(strings.Split(line, ","))
			updates = append(updates, pages)
		}
	}

	return rules, updates
}
