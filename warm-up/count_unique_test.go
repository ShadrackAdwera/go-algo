package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountUniqueValues(t *testing.T) {
	testCases := []struct{
		name string
		values []int
		expectedResult int
	}{
		{
			name: "Test All unique values",
			values: []int{1, 2, 3, 4, 5},
			expectedResult: 5,
		},
		{
			name: "Test Empty Slice",
			values: []int{},
			expectedResult: 0,
		},
		{
			name: "Test Single element",
			values: []int{42},
			expectedResult: 1,
		},
		{
			name: "Test One Unique",
			values: []int{1,1,1,1,1,1,1,2},
			expectedResult: 2,
		},
		{
			name: "Test All same values",
			values: []int{7, 7, 7, 7, 7},
			expectedResult: 1,
		},
		{
			name: "Test Mixed values",
			values: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
			expectedResult: 5,
		},
		{
			name: "Test Negative Values",
			values: []int{-3, -2, -2, -1, -1, 0, 1, 1, 2, 3},
			expectedResult: 7,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := countUniqueValues(testCase.values)
			require.Equal(t, testCase.expectedResult, result)
		})
	}
}