package search

// RecursiveBinarySearch finds an element (needle) in a list of values (haystack)
// If the element is not found returns -1
func RecursiveBinarySearch(haystack []int, needle int) int {
	return _recursiveBinarySearch(haystack, needle, 0, len(haystack)-1)
}

func _recursiveBinarySearch(haystack []int, needle int, leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex {
		return -1
	}
	middleIndex := (leftIndex + rightIndex) / 2

	if haystack[middleIndex] == needle {
		return middleIndex
	}

	if haystack[middleIndex] < needle {
		return _recursiveBinarySearch(haystack, needle, middleIndex+1, rightIndex)
	}
	return _recursiveBinarySearch(haystack, needle, leftIndex, middleIndex-1)
}
