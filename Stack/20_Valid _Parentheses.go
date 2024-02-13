package stack

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// link: https://leetcode.com/problems/valid-parentheses/
// isValid checks if a string is valid.
// A string is valid if:
// - Open brackets must be closed by the same type of brackets
// - Open brackets must be closed in the correct order.
// - Every close bracket has a corresponding open bracket of the same type.
//
// Parameter(s):
// s string - the string to be checked.
// Return type(s):
// bool - true if the string is valid, false otherwise.
func isValid(s string) bool {
	var stack Stack
	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else if c == ')' || c == ']' || c == '}' {
			if stack.IsEmpty() {
				return false
			}

			top := stack.Pop().(byte)
			if (c == ')' && top != '(') ||
				(c == ']' && top != '[') ||
				(c == '}' && top != '{') {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
