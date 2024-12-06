package part1

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 6 - Part 1")
	count := countDistinctPositions(getInput())
	fmt.Println("Count of distinct positions: ", count)
}

const GUARD = '^'
const OBSTACLE = '#'

var directions = [4][2]int{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

func countDistinctPositions(grid [][]byte) int {
	visited := make(map[string][4]bool)
	coordinates := getGuardCoordinates(grid)
	d := 0
	outOfBoundary := false

	for !outOfBoundary {
		markVisited(coordinates[0], coordinates[1], d, visited)
		nextMove := [2]int{
			coordinates[0] + directions[d][0],
			coordinates[1] + directions[d][1],
		}

		if !isWithinBoundary(grid, nextMove[0], nextMove[1]) {
			outOfBoundary = true
			continue
		}
		if handleObstacle(grid, nextMove[0], nextMove[1], &d) {
			continue
		}

		coordinates[0], coordinates[1] = nextMove[0], nextMove[1]
	}

	return len(visited)
}

func markVisited(x int, y int, d int, visited map[string][4]bool) {
	key := fmt.Sprintf("%d-%d", x, y)
	cell, existed := visited[key]
	if !existed {
		visited[key] = [4]bool{}
	}
	if cell[d] {
		panic("Guard is a loop")
	}
	cell[d] = true
	visited[key] = cell
}

func isWithinBoundary(grid [][]byte, x int, y int) bool {
	withinXBoundary := x >= 0 && x < len(grid[0])
	withinYBoundary := y >= 0 && y < len(grid)
	return withinXBoundary && withinYBoundary
}

func handleObstacle(grid [][]byte, x int, y int, d *int) bool {
	if grid[y][x] == OBSTACLE {
		*d = (*d + 1) % 4
		return true
	}
	return false
}

func getGuardCoordinates(grid [][]byte) [2]int {
	for y, row := range grid {
		for x, cell := range row {
			if cell == GUARD {
				return [2]int{x, y}
			}
		}
	}
	panic("No guard found")
}

func getInput() [][]byte {
	file, err := os.Open("./day06/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	return grid
}
