package part2

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 1 - Part 2")
	score := getScore(getLists())
	fmt.Println("Score:", score)
}

func getScore(list1 []int, list2 []int) int {
	distance := 0

	numFrequency := getNumFrequency(list2)

	for _, num := range list1 {
		distance += num * numFrequency[num]
	}

	return distance
}

func getNumFrequency(list []int) map[int]int {
	numFrequency := make(map[int]int)
	for _, num := range list {
		numFrequency[num]++
	}
	return numFrequency
}

func getLists() ([]int, []int) {
	file, err := os.Open("./day01/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list1 []int
	var list2 []int
	for scanner.Scan() {
		line := scanner.Text()
		nums := common.ConvertStringsToInts(strings.Split(line, "   "))
		list1 = append(list1, nums[0])
		list2 = append(list2, nums[1])
	}

	return list1, list2
}
