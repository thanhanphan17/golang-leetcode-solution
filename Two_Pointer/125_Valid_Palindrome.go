package twopointer

import "unicode"

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// link: https://leetcode.com/problems/valid-palindrome/description/
// isPalindrome checks if a string is a palindrome.
//
// Parameter(s):
// s string - the string to be checked.
// Return type(s):
// bool - true if the string is a palindrome, false otherwise.
func isPalindrome(s string) bool {
	first := 0
	last := len(s) - 1

	for first < last {
		if !isAlphanumeric(rune(s[first])) {
			first++
			continue
		}

		if !isAlphanumeric(rune(s[last])) {
			last--
			continue
		}

		if unicode.ToLower(rune(s[first])) != unicode.ToLower(rune(s[last])) {
			return false
		}

		first++
		last--
	}

	return true
}
