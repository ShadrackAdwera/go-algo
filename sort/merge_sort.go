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
