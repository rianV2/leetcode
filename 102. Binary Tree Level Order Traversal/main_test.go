package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "Balanced tree",
			input:    []int{3, 9, 20, -1, -1, 15, 7},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Complex tree",
			input:    []int{1, 2, 3, 4, -1, -1, 5},
			expected: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name:     "Complex tree 2",
			input:    []int{1, 2, 3, 4, -1, -1, 5, 1, 1, 1, -1},
			expected: [][]int{{1}, {2, 3}, {4, 5}, {1, 1, 1}},
		},
		{
			name:     "Left skewed tree",
			input:    []int{1, 2, -1, 3, -1, 4, -1, 5},
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromArray(tt.input)
			result := levelOrder(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}
