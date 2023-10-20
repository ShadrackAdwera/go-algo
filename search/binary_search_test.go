package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name           string
		values         []string
		valueToSearch  string
		indexOfElement int
	}{
		{
			name:           "Test Found String in Slice",
			values:         []string{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "evil",
			indexOfElement: 4,
		},
		{
			name:           "Test String Not Found in Slice",
			values:         []string{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "fruit",
			indexOfElement: -1,
		},
		{
			name:           "Test Empty Slice",
			values:         []string{},
			valueToSearch:  "apple",
			indexOfElement: -1,
		},
		{
			name:           "Test Found String in Long Slice",
			values:         []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
			valueToSearch:  "k",
			indexOfElement: 10,
		},
		{
			name:           "Test Found String at First Index",
			values:         []string{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "apple",
			indexOfElement: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := binarySearch(testCase.values, testCase.valueToSearch)
			require.Equal(t, testCase.indexOfElement, result)
		})
	}
}

func TestBinarySearchWithDynamicParams(t *testing.T) {
	testCases := []struct {
		name           string
		values         []interface{}
		valueToSearch  interface{}
		indexOfElement int
	}{
		{
			name:           "Test Found String in Interface",
			values:         []interface{}{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "evil",
			indexOfElement: 4,
		},
		{
			name:           "Test String Not Found in Interface",
			values:         []interface{}{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "fruit",
			indexOfElement: -1,
		},
		{
			name:           "Test Empty Interface",
			values:         []interface{}{},
			valueToSearch:  "apple",
			indexOfElement: -1,
		},
		{
			name:           "Test Found String in Long Interface",
			values:         []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
			valueToSearch:  "k",
			indexOfElement: 10,
		},
		{
			name:           "Test Found String at First Index",
			values:         []interface{}{"apple", "banana", "cat", "dog", "evil"},
			valueToSearch:  "apple",
			indexOfElement: 0,
		},
		{
			name:           "Test Found Int in Interface",
			values:         []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			valueToSearch:  5,
			indexOfElement: 4,
		},
		{
			name:           "Test Int Not Found in Interface",
			values:         []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			valueToSearch:  10,
			indexOfElement: -1,
		},
		{
			name:           "Test Empty Interface",
			values:         []interface{}{},
			valueToSearch:  3,
			indexOfElement: -1,
		},
		{
			name:           "Test Found Int in Long Interface",
			values:         []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
			valueToSearch:  11,
			indexOfElement: 10,
		},
		{
			name:           "Test Found Int at First Index",
			values:         []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			valueToSearch:  1,
			indexOfElement: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := binarySearchWithDynamicParameters(testCase.values, testCase.valueToSearch)
			require.Equal(t, testCase.indexOfElement, result)
		})
	}
}
