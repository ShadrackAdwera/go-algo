package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSame(t *testing.T) {
	testCases := []struct{
		name string
		num1 []int
		num2 []int
		expectedResult bool
	}{
		{
			name: "test uneven slice length",
			num1: []int{1,2,3},
			num2: []int{1,9},
			expectedResult: false,
		},
		{
			name: "test OK",
			num1: []int{1,2,3},
			num2: []int{1,4,9},
			expectedResult: true,
		},
		{
			name: "test different frequencies",
			num1: []int{1,2,1},
			num2: []int{4,4,1},
			expectedResult: false,
		},
		{
			name: "test empty slices",
			num1: []int{},
			num2: []int{},
			expectedResult: true,
		},
		{
			name: "test repeated numbers",
			num1: []int{1, 1, 2, 2, 3, 3},
			num2: []int{1, 1, 4, 4, 9, 9},
			expectedResult: true,
		},
		{
			name: "test negative numbers",
			num1: []int{-1, -2, -3},
			num2: []int{1, 4, 9},
			expectedResult: true,
		},
		{
			name: "test non-squared values",
			num1: []int{1, 2, 3},
			num2: []int{1, 3, 5},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Same(testCase.num1, testCase.num2)
			require.Equal(t, result, testCase.expectedResult)
		})
	}
}