package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeFn(t *testing.T) {
	testCases := []struct {
		name       string
		leftSlice  []int
		rightSlice []int
		result     []int
	}{
		{
			name:       "Equal length slices",
			leftSlice:  []int{1, 2, 3, 4},
			rightSlice: []int{5, 6, 7, 8},
			result:     []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:       "Smaller Right Slice",
			leftSlice:  []int{5, 6, 7, 8},
			rightSlice: []int{1, 2, 3, 4},
			result:     []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:       "Short left slice",
			leftSlice:  []int{1, 2},
			rightSlice: []int{5, 6, 7, 8},
			result:     []int{1, 2, 5, 6, 7, 8},
		},
		{
			name:       "Short right slice",
			leftSlice:  []int{1, 2, 3, 4, 5, 6},
			rightSlice: []int{7, 8},
			result:     []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:       "Empty left slice",
			leftSlice:  []int{},
			rightSlice: []int{6, 7, 8},
			result:     []int{6, 7, 8},
		},
		{
			name:       "Empty right slice",
			leftSlice:  []int{1, 2, 3, 4, 5, 6},
			rightSlice: []int{},
			result:     []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:       "Empty slices",
			leftSlice:  []int{},
			rightSlice: []int{},
			result:     nil,
		},
		{
			name:       "Duplicates",
			leftSlice:  []int{1, 2, 3, 3, 5, 5},
			rightSlice: []int{6, 7, 8},
			result:     []int{1, 2, 3, 3, 5, 5, 6, 7, 8},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := merge(testCase.leftSlice, testCase.rightSlice)
			require.Equal(t, result, testCase.result)
		})
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name          string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			name:          "Normal unsorted slice",
			unsortedSlice: []int{3, 6, 2, 9, 1},
			sortedSlice:   []int{1, 2, 3, 6, 9},
		},
		{
			name:          "Empty slice",
			unsortedSlice: []int{},
			sortedSlice:   []int{},
		},
		{
			name:          "Already sorted slice",
			unsortedSlice: []int{1, 2, 3, 4, 5},
			sortedSlice:   []int{1, 2, 3, 4, 5},
		},
		{
			name:          "All equal elements",
			unsortedSlice: []int{5, 5, 5, 5, 5},
			sortedSlice:   []int{5, 5, 5, 5, 5},
		},
		{
			name:          "Duplicate elements",
			unsortedSlice: []int{3, 1, 2, 1, 4, 2, 3},
			sortedSlice:   []int{1, 1, 2, 2, 3, 3, 4},
		},
		{
			name:          "Large unsorted slice",
			unsortedSlice: []int{100, 50, 75, 200, 150, 25, 10, 300},
			sortedSlice:   []int{10, 25, 50, 75, 100, 150, 200, 300},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mergeSort(testCase.unsortedSlice)
			require.Equal(t, result, testCase.sortedSlice)
		})
	}
}
