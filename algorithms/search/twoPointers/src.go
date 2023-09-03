package twoPointers

// O(n) for time complexity
// O(1) for space Complexity
func twoPointers(input []int, target int) bool {
	//Input of this algorithm must be sorted

	rightPointer := 0
	leftPointer := len(input) - 1

	for rightPointer != leftPointer {
		summation := input[rightPointer] + input[leftPointer]
		if summation == target {
			return true
		} else if summation > target {
			leftPointer--
		} else if summation < target {
			rightPointer++
		}
	}

	return false
}
