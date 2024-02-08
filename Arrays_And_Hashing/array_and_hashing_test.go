package arrays_and_hashing

import (
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
