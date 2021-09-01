package utils

import "github.com/texttheater/golang-levenshtein/levenshtein"

func StringMatch(str1, str2 string, r float64) bool {
	ratio := levenshtein.RatioForStrings([]rune(str1), []rune(str2), levenshtein.DefaultOptions)
	return ratio >= r
}
