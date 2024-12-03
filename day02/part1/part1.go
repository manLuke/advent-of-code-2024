package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 2 - Part 1")
	safeReports := getNumOfSafeReports(getReports())
	fmt.Println("Safe reports:", safeReports)
}

func getNumOfSafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isValidReport(report) {
			safeReports++
		}
	}

	return safeReports
}

func isValidReport(report []int) bool {
	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		l1 := report[i]
		l2 := report[i+1]
		if (increasing && l1 >= l2) || (!increasing && l1 <= l2) || !isValidDifference(l1, l2) {
			return false
		}
	}
	return true
}

func isValidDifference(n1 int, n2 int) bool {
	var difference int
	if n1 < n2 {
		difference = n2 - n1
	} else {
		difference = n1 - n2
	}
	return difference > 0 && difference <= 3
}

func getReports() [][]int {
	file, err := os.Open("./day02/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, common.ConvertStringsToInts(strings.Split(line, " ")))
	}

	return reports
}
