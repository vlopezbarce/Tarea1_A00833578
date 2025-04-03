package main

// Stack represents a simple integer stack
type Stack struct {
	items []int
}

// Push adds an element to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Stack is empty
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Remove the last element
	return item, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Stack is empty
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
