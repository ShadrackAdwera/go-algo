package search

func linearSearch(values []string, valueToSearch string) int {

	for i := 0; i < len(values); i++ {
		if values[i] == valueToSearch {
			return i
		}
	}

	return -1
}
