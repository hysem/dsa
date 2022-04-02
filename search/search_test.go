package search_test

import (
	"fmt"
	"testing"

	"github.com/hysem/dsa/search"
	"github.com/stretchr/testify/assert"
)

var (
	totalItems = 10000
	haystack   = make([]int, 0, totalItems)
)

func init() {
	for i := 0; i < totalItems; i++ {
		haystack = append(haystack, i+1)
	}
}

func testSearch(t *testing.T, searchFn func(haystack []int, needle int) int) {
	testCases := map[string]struct {
		needle         int
		expectedResult int
	}{
		"no such element": {
			needle:         -1,
			expectedResult: -1,
		},
		"found element at index:0": {
			needle:         1,
			expectedResult: 0,
		},
		"found element at index:10000": {
			needle:         10000,
			expectedResult: 9999,
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actualResult := searchFn(haystack, tc.needle)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

func benchmarkSearch(b *testing.B, searchFn func(haystack []int, needle int) int) {
	for needle := 0; needle <= totalItems; needle += 1000 {
		b.Run(fmt.Sprintf("items in haystack: %d, needle: %d", totalItems, needle), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				searchFn(haystack, needle)
			}
		})
	}
}

func TestLinearSearch(t *testing.T) {
	testSearch(t, search.LinearSearch)
}

func BenchmarkLinearSearch(b *testing.B) {
	benchmarkSearch(b, search.LinearSearch)
}

func TestBinarySearch(t *testing.T) {
	testSearch(t, search.BinarySearch)
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmarkSearch(b, search.BinarySearch)
}

func TestRecursiveBinarySearch(t *testing.T) {
	testSearch(t, search.RecursiveBinarySearch)
}

func BenchmarkRecursiveBinarySearch(b *testing.B) {
	benchmarkSearch(b, search.RecursiveBinarySearch)
}
