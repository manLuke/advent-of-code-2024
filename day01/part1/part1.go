package part1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	distance := getDistance(getLists())
	fmt.Println("Distance:", distance)
}

func getDistance(list1 []int, list2 []int) int {
	distance := 0

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		distance += getDifference(list1[i], list2[i])
	}

	return distance
}

func getDifference(n1 int, n2 int) int {
	if n1 < n2 {
		return n2 - n1
	} else {
		return n1 - n2
	}
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
