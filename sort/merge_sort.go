package sort

func merge(leftSlice []int, rightSlice []int) []int {
	var result []int

	i, j := 0, 0

	for i < len(leftSlice) && j < len(rightSlice) {
		if leftSlice[i] < rightSlice[j] {
			result = append(result, leftSlice[i])
			i++
		} else {
			result = append(result, rightSlice[j])
			j++
		}
	}

	for i < len(leftSlice) {
		result = append(result, leftSlice[i])
		i++
	}

	for j < len(rightSlice) {
		result = append(result, rightSlice[j])
		j++
	}

	return result
}

func mergeSort(values []int) []int {
	// base case
	if len(values) <= 1 {
		return values
	}

	middleIdx := len(values) / 2

	return merge(mergeSort(values[0:middleIdx]), mergeSort(values[middleIdx:]))
}
