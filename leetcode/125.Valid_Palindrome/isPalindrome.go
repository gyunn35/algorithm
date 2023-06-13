package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	strs := []rune(strings.ToLower(s))

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		for i < j && !unicode.IsDigit(rune(strs[i])) && !unicode.IsLetter(rune(strs[i])) {
			i++
		}
		for i < j && !unicode.IsDigit(rune(strs[j])) && !unicode.IsLetter(rune(strs[j])) {
			j--
		}
		if strs[i] != strs[j] {
			return false
		}
	}

	return true
}
