package search

import (
	"math"
	"sort"
)

func binarySearch(values []string, valueToSearch string) int {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	startIndex := 0
	endIndex := len(values) - 1
	// [1,2,3,4,5,6,7,8,9]
	for startIndex <= endIndex {
		middleIndex := startIndex + int(math.Floor((float64(endIndex)-float64(startIndex))/2))

		if values[middleIndex] == valueToSearch {
			return middleIndex
		}
		if values[middleIndex] < valueToSearch {
			startIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}

	return -1
}

// binary search with dynamic parameters

func binarySearchWithDynamicParameters(values []interface{}, valueToSearch interface{}) int {
	switch valueToSearch := valueToSearch.(type) {
	case int:
		vals := make([]int, len(values))

		for i, v := range values {
			vals[i] = v.(int)
		}

		idx := sort.SearchInts(vals, valueToSearch)

		if idx < len(vals) && vals[idx] == valueToSearch {
			return idx
		}
	case string:
		vals := make([]string, len(values))
		for i, v := range values {
			vals[i] = v.(string)
		}

		idx := sort.SearchStrings(vals, valueToSearch)

		if idx < len(vals) && vals[idx] == valueToSearch {
			return idx
		}
	}

	return -1
}
