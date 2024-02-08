package arrays_and_hashing

// link: https://leetcode.com/problems/valid-anagram/description/
// isAnagram checks if two strings are anagrams of each other.
//
// Parameter(s):
// s string - the first string to be checked for anagrams.
// t string - the second string to be checked for anagrams.
//
// Return type(s):
// bool - true if the strings are anagrams, false otherwise.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Create an array of size 26 to store the count of each character in s and t.
	store := make([]int, 26)

	for i := 0; i < len(s); i++ {
		store[s[i]-'a']++
		store[t[i]-'a']--
	}

	// Check if all the values in the array are 0.
	for i := 0; i < len(store); i++ {
		if store[i] != 0 {
			return false
		}
	}

	return true
}
