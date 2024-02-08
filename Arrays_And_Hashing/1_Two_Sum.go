package arrays_and_hashing

// link: https://leetcode.com/problems/two-sum/description/
// twoSum finds the indices of two numbers that add up to a given target.
//
// Parameter(s):
// nums []int - the array of integers to search for the target.
// target int - the target sum to search for in the array.
//
// Return type(s):
// []int - the indices of the two numbers that add up to the target.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		if val, ok := seen[target-num]; ok {
			return []int{val, i}
		}
		seen[num] = i
	}

	return []int{-1, -1}
}
