package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 7 - Part 1")
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

	multiplicationSubtotal := subtotal * eq.Nums[i]
	additionSubtotal := subtotal + eq.Nums[i]

	validMultiplicationSubtotal := isSubtotalValid(multiplicationSubtotal, eq.Solution)
	validAdditionSubtotal := isSubtotalValid(additionSubtotal, eq.Solution)

	if validMultiplicationSubtotal && validAdditionSubtotal {
		return evaluateEquation(eq, multiplicationSubtotal, i+1) || evaluateEquation(eq, additionSubtotal, i+1)
	} else if validMultiplicationSubtotal {
		return evaluateEquation(eq, multiplicationSubtotal, i+1)
	} else if validAdditionSubtotal {
		return evaluateEquation(eq, additionSubtotal, i+1)
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
