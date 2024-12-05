package part1

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 4 - Part 1")
	count := countXMasMatches(getMatrix())
	fmt.Println("Count:", count)
}

const searchedWord = "XMAS"

func countXMasMatches(matrix [][]byte) int {
	count := 0
	for y := range len(matrix) {
		for x := range len(matrix[0]) {
			for _, dir := range directions {
				if isWithinBoundary(matrix, x, y, dir, searchedWord) && containWord(matrix, x, y, dir, searchedWord) {
					count++
				}
			}
		}
	}
	return count
}

func isWithinBoundary(matrix [][]byte, startX int, startY int, direction Direction, word string) bool {
	endX := startX + (len(word)-1)*direction.dx
	endY := startY + (len(word)-1)*direction.dy

	withinXBoundary := endX >= 0 && endX < len(matrix[0])
	withinYBoundary := endY >= 0 && endY < len(matrix)

	return withinXBoundary && withinYBoundary
}

func containWord(matrix [][]byte, startX int, startY int, direction Direction, word string) bool {
	for i := range word {
		char := matrix[startY+direction.dy*i][startX+direction.dx*i]
		if char != word[i] {
			return false
		}
	}
	return true
}

type Direction struct {
	dx int
	dy int
}

var directions = map[string]Direction{
	"up":         {dx: 0, dy: -1},
	"up-right":   {dx: 1, dy: -1},
	"right":      {dx: 1, dy: 0},
	"down-right": {dx: 1, dy: 1},
	"down":       {dx: 0, dy: 1},
	"down-left":  {dx: -1, dy: 1},
	"left":       {dx: -1, dy: 0},
	"up-left":    {dx: -1, dy: -1},
}

func getMatrix() [][]byte {
	file, err := os.Open("./day04/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	return matrix
}
