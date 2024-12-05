package part1

import "testing"

func TestIsWithinBoundary(t *testing.T) {
	matrix := [][]byte{
		{'A', 'B', 'C', 'D'},
		{'E', 'F', 'G', 'H'},
		{'I', 'J', 'K', 'L'},
		{'M', 'N', 'O', 'P'},
	}

	tests := []struct {
		startX, startY int
		direction      Direction
		word           string
		expected       bool
	}{
		{1, 1, Direction{dx: 1, dy: 0}, "HELLO", false},
		{0, 0, Direction{dx: 1, dy: 0}, "ABCD", true},
		{0, 0, Direction{dx: 0, dy: 1}, "AEIM", true},
		{3, 3, Direction{dx: 0, dy: 1}, "XYZ", false},
		{1, 1, Direction{dx: 1, dy: 1}, "FGK", true},
		{2, 2, Direction{dx: 1, dy: 1}, "KLP", false},
		{-1, 0, Direction{dx: 0, dy: 0}, "AA", false},
		{0, -1, Direction{dx: 0, dy: 0}, "AA", false},
		{3, 3, Direction{dx: 1, dy: 0}, "PP", false},
		{3, 3, Direction{dx: 0, dy: 1}, "PP", false},
		{3, 3, Direction{dx: 0, dy: -1}, "PP", true},
		{3, 3, Direction{dx: -1, dy: 0}, "PP", true},
	}

	for _, test := range tests {
		result := isWithinBoundary(matrix, test.startX, test.startY, test.direction, test.word)
		if result != test.expected {
			t.Errorf("isWithinBoundary(%d, %d, %v, %s) = %v; want %v",
				test.startX, test.startY, test.direction, test.word, result, test.expected)
		}
	}
}
