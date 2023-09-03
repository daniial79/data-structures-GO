package linearSearch

// O(n) for time complexity
// O(1) for space complexity
func linearSearch(input []int, target int) int {
	for i := 0; i < len(input); i++ {
		if input[i] == target {
			return i
		}
	}

	return -1
}
