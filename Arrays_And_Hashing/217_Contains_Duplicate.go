package arrays_and_hashing

// link: https://leetcode.com/problems/contains-duplicate/description/
// containsDuplicate checks if the given array contains any duplicates.
//
// Parameter(s):
// nums []int - the array of integers to be checked for duplicates.
//
// Return type(s):
// bool - true if the array contains duplicates, false otherwise.
func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	// Use a map to store the elements seen so far.
	// A bool typically requires at least one byte of memory to store its value.
	// On the other hand, an empty struct {} in Go doesn't consume any additional memory
	// beyond the memory required to represent the map keys themselves.
	// It serves as a placeholder without adding any overhead per key.
	seen := make(map[int]struct{})

	for _, num := range nums {
		// If the element is already seen, return true.
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
