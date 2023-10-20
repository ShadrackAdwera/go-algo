package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
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
			result := bubbleSort(testCase.unsortedSlice)
			require.Equal(t, result, testCase.sortedSlice)
		})
	}
}
