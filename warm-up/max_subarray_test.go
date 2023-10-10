package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSubSliceNaive(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		subSliceLength int
		expectedResult int
	}{
		{
			name:           "Normal slice with positive numbers",
			nums:           []int{1, 2, 5, 2, 8, 1, 5},
			subSliceLength: 2,
			expectedResult: 10,
		},
		{
			name:           "All negative values",
			nums:           []int{-2, -3, -4, -1, -5},
			subSliceLength: 3,
			expectedResult: -8,
		},
		{
			name:           "Empty slice",
			nums:           []int{},
			subSliceLength: 2,
			expectedResult: 0,
		},
		{
			name:           "Slice length equal to the length of nums",
			nums:           []int{1, 2, 3, 4, 5},
			subSliceLength: 5,
			expectedResult: 15,
		},
		{
			name:           "Normal slice with mixed numbers",
			nums:           []int{1, -2, 3, -4, 5, -6, 7},
			subSliceLength: 4,
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := maxSubSliceSumNaive(testCase.nums, testCase.subSliceLength)
			require.Equal(t, result, testCase.expectedResult)
		})
	}
}

func TestMaxSubSliceOptimized(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		subSliceLength int
		expectedResult int
	}{
		{
			name:           "Normal slice with positive numbers",
			nums:           []int{1, 2, 5, 2, 8, 1, 5},
			subSliceLength: 2,
			expectedResult: 10,
		},
		{
			name:           "All negative values",
			nums:           []int{-2, -3, -4, -1, -5},
			subSliceLength: 3,
			expectedResult: -8,
		},
		{
			name:           "Empty slice",
			nums:           []int{},
			subSliceLength: 2,
			expectedResult: 0,
		},
		{
			name:           "Slice length equal to the length of nums",
			nums:           []int{1, 2, 3, 4, 5},
			subSliceLength: 5,
			expectedResult: 15,
		},
		{
			name:           "Normal slice with mixed numbers",
			nums:           []int{1, -2, 3, -4, 5, -6, 7},
			subSliceLength: 4,
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := maxSubSliceSumOptimized(testCase.nums, testCase.subSliceLength)
			require.Equal(t, result, testCase.expectedResult)
		})
	}
}
