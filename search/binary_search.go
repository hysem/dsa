package search

// BinarySearch finds an element (needle) in a list of values (haystack)
// If the element is not found returns -1
func BinarySearch(haystack []int, needle int) int {
	leftIndex := 0
	rightIndex := len(haystack)

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2

		if haystack[middleIndex] == needle {
			return middleIndex
		}

		if haystack[middleIndex] < needle {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex - 1
		}
	}
	return -1
}
