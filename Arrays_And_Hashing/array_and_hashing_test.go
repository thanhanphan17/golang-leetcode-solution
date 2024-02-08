package arrays_and_hashing

import (
	"reflect"
	"sort"
	"testing"
)

func Test_217ContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{input: []int{}, expected: false},
		{input: []int{1, 2, 3, 4, 5}, expected: false},
		{input: []int{1, 2, 3, 1}, expected: true},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true},
	}

	for _, test := range testCases {
		got := containsDuplicate(test.input)
		if got != test.expected {
			t.Errorf("For input %v, expected %t, but got %t", test.input, test.expected, got)
		}
	}
}

func Test_242ValidAnagram(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "", t: "", expected: true},
		{s: "anagram", t: "nagaram", expected: true},
		{s: "rat", t: "car", expected: false},
	}

	for _, test := range testCases {
		got := isAnagram(test.s, test.t)
		if got != test.expected {
			t.Errorf("For input %s and %s, expected %t, but got %t", test.s, test.t, test.expected, got)
		}
	}
}

func Test_1TwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, test := range testCases {
		got := twoSum(test.nums, test.target)
		if len(got) != len(test.expected) {
			t.Errorf("For input %v and %d, expected %v, but got %v", test.nums, test.target, test.expected, got)
		}
		for i := 0; i < len(got); i++ {
			if got[i] != test.expected[i] {
				t.Errorf("For input %v and %d, expected %v, but got %v", test.nums, test.target, test.expected, got)
			}
		}
	}
}

func Test_49GroupAnagrams(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output [][]string
	}{
		{"Empty input", []string{}, [][]string{}},
		{"Single string", []string{"a"}, [][]string{{"a"}}},
		{"Multiple anagram groups", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := groupAnagrams(test.input)
			sort.Slice(result, func(i, j int) bool {
				return len(result[i]) < len(result[j])
			})

			sort.Slice(test.output, func(i, j int) bool {
				return len(test.output[i]) < len(test.output[j])
			})

			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %v, but got %v", test.output, result)
			}
		})
	}
}
