package twopointer

// link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// twoSum finds the indices of two numbers that add up to a given target.
//
// Parameter(s):
// numbers []int - the array of integers to search for the target.
// target int - the target sum to search for in the array.
//
// Return type(s):
// []int - the indices of the two numbers that add up to the target.
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}
