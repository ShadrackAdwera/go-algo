package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearSearch(t *testing.T) {
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
			result := linearSearch(testCase.values, testCase.valueToSearch)
			require.Equal(t, testCase.indexOfElement, result)
		})
	}
}
