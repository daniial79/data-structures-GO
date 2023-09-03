package binarySearch

func binarySearch(input []int, target int) int {
	low := 0
	high := len(input) - 1

	for low < high {
		mid := (low + high) / 2

		if input[mid] == target {
			return mid
		} else if input[mid] < target {
			high = mid
		} else if input[mid] > target {
			low = mid
		}
	}

	return -1
}
