package search

// LinearSearch finds an element (needle) in a list of values (haystack)
// If the element is not found returns -1
func LinearSearch(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}
