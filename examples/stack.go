package examples

// Stack implementation borrowed from https://dev.to/jpoly1219/stacks-in-go-54k
type Stack struct {
	items []int
}

// Push pushes data in the stack.
//
// Contract:
//   - ensures stack grows by one element: @old{len(s.items)} == len(s.items) - 1
//   - ensures new data is on top of the stack: s.Top() == data
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop pops the top element from the stack.
//
// Contract:
//   - requires non empty stack: !s.IsEmpty()
//   - ensures stack shrinks by one element: @old{len(s.items)} == len(s.items) + 1
func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

// Top yields the data on the top of the stack.
//
// Contract:
//   - ensure non empty stack: !s.IsEmpty()
//   - ensure resulting data is the top of the stack: data == s.items[len(s.items)-1]
//   - unmodified stack top element is not modified: @old{s.items[len(s.items)-1]} == s.items[len(s.items)-1]
func (s *Stack) Top() (data int) {
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack has no data, false otherwise.
//
// Contract:
//   - let size := len(s.items)
//   - ensures if empty then true: size == 0 ==> result == true
//   - ensures if not empty then false: size != 0 ==> result == false
//   - unmodified stack size: len(s.items)
func (s *Stack) IsEmpty() (result bool) {
	return len(s.items) == 0
}
