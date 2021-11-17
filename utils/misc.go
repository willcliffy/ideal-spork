package utils

import (
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func StringMatch(str1, str2 string, r float64) bool {
	ratio := levenshtein.RatioForStrings(
		[]rune(strings.ToLower(str1)),
		[]rune(strings.ToLower(str2)),
		levenshtein.DefaultOptions)
	return ratio >= r
}

func StringArrayIndexOf(arr []string, target string) int {
	for i, x := range arr {
		if x == target {
			return i
		}
	}
	return -1
}

func BubbleSortWithIndices(arr []int) (sorted, indices []int) {
	n := len(arr)
	x, y := 0, 0

	sorted = make([]int, n)
	indices = make([]int, n)


	for i := 0; i < n; i++ {
		sorted[i] = arr[i]
		indices[i] = i
	}

	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < n - 1; i++ {
			for j := 0; j < n - i - 1; j++ {
				if sorted[i] < sorted[i+1] {
					swapped = true
					x, y  = sorted[i], indices[i]
					sorted[i], indices[i] = sorted[i+1], indices[i+1]
					sorted[i+1], indices[i+1] = x, y
				}
			}
		}
	}

	return
}
