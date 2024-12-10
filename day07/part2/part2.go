package part2

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 7 - Part 2")
	sum := calculateSumOfValidEquations(getEquations())
	fmt.Println("Sum of valid equations:", sum)
}

func calculateSumOfValidEquations(equations []Equation) int {
	sum := 0

	for _, eq := range equations {
		if evaluateEquation(eq, 0, 0) {
			sum += eq.Solution
		}
	}
	return sum
}

func evaluateEquation(eq Equation, subtotal int, i int) bool {
	if i >= len(eq.Nums) {
		return eq.Solution == subtotal
	}

	operations := []int{
		subtotal * eq.Nums[i],
		subtotal + eq.Nums[i],
		common.ConcatenateNumbers(subtotal, eq.Nums[i]),
	}

	for _, newSubtotal := range operations {
		if isSubtotalValid(newSubtotal, eq.Solution) {
			if evaluateEquation(eq, newSubtotal, i+1) {
				return true
			}
		}
	}

	return false
}

func isSubtotalValid(subtotal int, solution int) bool {
	return subtotal <= solution
}

type Equation struct {
	Solution int
	Nums     []int
}

func getEquations() []Equation {
	file, err := os.Open("./day07/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var equations []Equation
	for scanner.Scan() {
		line := scanner.Text()
		equation := strings.Split(line, ": ")
		equations = append(equations, Equation{
			Solution: common.StringToInt(equation[0]),
			Nums:     common.ConvertStringsToInts(strings.Split(equation[1], " ")),
		})

	}

	return equations
}
