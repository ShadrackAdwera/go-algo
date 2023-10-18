package sort

func bubbleSort(values []int) []int {
	for i := len(values); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if values[j] > values[j+1] {
				temp := values[j+1]
				values[j+1] = values[j]
				values[j] = temp
			}
		}
	}
	return values
}
