package arrays_and_hashing

// link: https://leetcode.com/problems/group-anagrams/description/
// groupAnagrams groups an array of strings into anagrams.
//
// Parameter(s):
// - strs []string: array of strings to group
// Return type(s):
// [][]string: an array of grouped anagrams
func groupAnagrams(strs []string) [][]string {
	normalize := func(word string) [26]uint8 {
		store := [26]uint8{}

		for i := range word {
			store[word[i]-'a']++
		}

		return store
	}

	groupMap := make(map[[26]uint8][]string)

	for _, word := range strs {
		store := normalize(word)
		groupMap[store] = append(groupMap[store], word)
	}

	result := make([][]string, 0, len(groupMap))

	for i := range groupMap {
		result = append(result, groupMap[i])
	}

	return result
}
