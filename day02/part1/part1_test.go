package part1

import "testing"

func TestIsValidReport(t *testing.T) {
	tests := []struct {
		input  []int
		result bool
	}{
		// From the assignment
		{[]int{7, 6, 4, 2, 1}, true},  // Safe because the levels are all decreasing by 1 or 2.
		{[]int{1, 2, 7, 8, 9}, false}, // Unsafe because 2 7 is an increase of 5.
		{[]int{9, 7, 6, 2, 1}, false}, // Unsafe because 6 2 is a decrease of 4.
		{[]int{1, 3, 2, 4, 5}, false}, // Unsafe because 1 3 is increasing but 3 2 is decreasing.
		{[]int{8, 6, 4, 4, 1}, false}, // Unsafe because 4 4 is neither an increase or a decrease.
		{[]int{1, 3, 6, 7, 9}, true},  // Safe because the levels are all increasing by 1, 2, or 3.

		// Valid cases
		{[]int{1, 2, 3, 4, 5, 6, 7}, true},                    // Strictly increasing by 1
		{[]int{20, 18, 15, 12, 10, 7, 4}, true},               // Strictly decreasing by 2-3
		{[]int{5, 8, 11, 14, 17, 20, 23}, true},               // Increasing by exactly 3
		{[]int{30, 27, 24, 21, 18, 15, 12}, true},             // Decreasing by 3 consistently
		{[]int{1, 3, 5, 7, 9, 11, 13, 15}, true},              // Increasing by 2
		{[]int{100, 97, 94, 91, 88, 85, 82, 79}, true},        // Decreasing by 3
		{[]int{10, 11, 12, 13, 14, 15, 16, 17, 18}, true},     // Strictly increasing by 1
		{[]int{50, 48, 46, 44, 42, 40, 38, 36, 34}, true},     // Decreasing by 2
		{[]int{1, 3, 6, 9, 12, 15, 18, 21}, true},             // All increasing but one jump is > 3
		{[]int{7, 6, 5, 4, 3, 2, 1, 0, -1}, true},             // Strictly decreasing by 1
		{[]int{20, 22, 24, 26, 28, 30, 32, 34, 36, 38}, true}, // Strictly increasing by 2

		// Invalid cases
		{[]int{1, 3, 6, 10, 14, 18, 22}, false},             // Jump exceeds 3
		{[]int{10, 11, 12, 14, 14, 15, 16}, false},          // Contains duplicates
		{[]int{2, 4, 7, 10, 8, 6, 4}, false},                // Mixed direction
		{[]int{1, 2, 3, 4, 5, 9, 12}, false},                // A jump exceeds 3
		{[]int{9, 9, 8, 7, 6, 5, 4}, false},                 // Contains same adjacent levels
		{[]int{1, 2, 4, 8, 16, 32, 64}, false},              // Exponential jumps
		{[]int{10, 11, 13, 16, 20, 25, 31}, false},          // Jumps exceed 3
		{[]int{5, 5, 6, 7, 8, 9, 10}, false},                // Duplicate values
		{[]int{10, 12, 14, 15, 13, 11, 9, 7}, false},        // Direction changes
		{[]int{100, 98, 95, 91, 88, 85, 83, 80, 78}, false}, // Mixed valid and invalid jumps
		{[]int{3, 5, 7, 10, 13, 17, 20, 22, 25}, false},     // Mixed step sizes
		{[]int{3, 5, 7, 10, 13, 17, 20, 22, 25}, false},     // Mixed step sizes
		{[]int{49, 50, 48, 46, 44, 42, 40, 38, 36, 34}, false},
	}

	for _, test := range tests {
		result := isValidReport(test.input)
		if result != test.result {
			t.Errorf("isValidReport(%d) = %t; want %t", test.input, result, test.result)
		}
	}
}
