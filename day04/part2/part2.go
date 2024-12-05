package part2

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 4 - Part 2")
	count := countXMasMatches(getMatrix())
	fmt.Println("Count:", count)
}

func countXMasMatches(matrix [][]byte) int {
	count := 0
	for y := range len(matrix) {
		for x := range len(matrix[0]) {
			if matrix[y][x] == 'A' && isValidDiagonalPositions(matrix, x, y) && checkDiagonalCharDifference(matrix, x, y) {
				count++
			}
		}
	}
	return count
}

func isValidDiagonalPositions(matrix [][]byte, startX int, startY int) bool {
	for _, dir := range directions {
		if !isWithinBoundary(matrix, startX+dir.dx, startY+dir.dy) {
			return false
		}
	}
	return true
}

func isWithinBoundary(matrix [][]byte, x int, y int) bool {
	withinXBoundary := x >= 0 && x < len(matrix[0])
	withinYBoundary := y >= 0 && y < len(matrix)
	return withinXBoundary && withinYBoundary
}

func checkDiagonalCharDifference(matrix [][]byte, startX int, startY int) bool {
	upLeftChar := matrix[startY-1][startX-1]
	upRightChar := matrix[startY-1][startX+1]
	downLeftChart := matrix[startY+1][startX-1]
	downRightChart := matrix[startY+1][startX+1]

	return common.Abs(int(upLeftChar)-int(downRightChart)) == 6 && common.Abs(int(upRightChar)-int(downLeftChart)) == 6
}

type Direction struct {
	dx int
	dy int
}

var directions = map[string]Direction{
	"up-left":    {dx: -1, dy: -1},
	"up-right":   {dx: 1, dy: -1},
	"down-left":  {dx: -1, dy: 1},
	"down-right": {dx: 1, dy: 1},
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
