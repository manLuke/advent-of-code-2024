package part2

import (
	"reflect"
	"testing"
)

func TestGetNumFrequency(t *testing.T) {
	tests := []struct {
		input    []int
		expected map[int]int
	}{
		{
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expected: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			input:    []int{5, 5, 5, 5, 5},
			expected: map[int]int{5: 5},
		},
		{
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: map[int]int{1: 2, 2: 2, 3: 2, 4: 2},
		},
		{
			input:    []int{},
			expected: map[int]int{},
		},
	}

	for _, test := range tests {
		result := getNumFrequency(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("getNumFrequency(%v) = %v; want  %v", test.input, result, test.expected)
		}
	}
}
