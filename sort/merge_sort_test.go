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
