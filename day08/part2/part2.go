package part2

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 8 - Part 2")
	uniqueAntinodeCount := countUniqueAntinodePositions(getMap())
	fmt.Println("Number of unique positions of antinode: ", uniqueAntinodeCount)
}

func countUniqueAntinodePositions(grid [][]byte) int {
	antennaLocations := getAntennaLocations(grid)
	antinodeLocations := make(map[string]bool)
	for _, positions := range antennaLocations {
		for i := 0; i < len(positions)-1; i++ {
			for j := 0 + i + 1; j < len(positions); j++ {
				addAntinodeLocations(grid, antinodeLocations, positions[i], positions[j])
			}
		}
	}
	return len(antinodeLocations)
}

func addAntinodeLocations(grid [][]byte, antinodes map[string]bool, firstAntenna Coordinate, secondAntenna Coordinate) {
	vectorCoordinates := Coordinate{
		X: secondAntenna.X - firstAntenna.X,
		Y: secondAntenna.Y - firstAntenna.Y,
	}
	potentialAntinodes := [2]AnnotatedCoordinate{
		{
			X:      firstAntenna.X,
			Y:      firstAntenna.Y,
			Adding: false,
		},
		{
			X:      secondAntenna.X,
			Y:      secondAntenna.Y,
			Adding: true,
		},
	}
	for i := 0; i < len(potentialAntinodes); i++ {
		processAntinode(grid, antinodes, &potentialAntinodes[i], vectorCoordinates)
	}
}

func processAntinode(grid [][]byte, antinodes map[string]bool, antinode *AnnotatedCoordinate, vector Coordinate) {
	for isWithinBoundary(grid, antinode.X, antinode.Y) {
		key := fmt.Sprintf("%d-%d", antinode.X, antinode.Y)
		antinodes[key] = true
		moveAntinode(antinode, vector)
	}
}

func moveAntinode(antinode *AnnotatedCoordinate, vector Coordinate) {
	if antinode.Adding {
		antinode.X += vector.X
		antinode.Y += vector.Y
	} else {
		antinode.X -= vector.X
		antinode.Y -= vector.Y
	}
}

func isWithinBoundary(matrix [][]byte, x int, y int) bool {
	withinXBoundary := x >= 0 && x < len(matrix[0])
	withinYBoundary := y >= 0 && y < len(matrix)
	return withinXBoundary && withinYBoundary
}

func getAntennaLocations(grid [][]byte) map[byte][]Coordinate {
	coordinates := make(map[byte][]Coordinate)
	for y, row := range grid {
		for x, antenna := range row {
			if antenna != '.' {
				_, exists := coordinates[antenna]
				if !exists {
					coordinates[antenna] = make([]Coordinate, 0)
				}
				coordinates[antenna] = append(coordinates[antenna], Coordinate{X: x, Y: y})
			}
		}
	}

	return coordinates
}

type Coordinate struct {
	X, Y int
}

type AnnotatedCoordinate struct {
	X, Y   int
	Adding bool
}

func getMap() [][]byte {
	file, err := os.Open("./day08/input.txt")
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
