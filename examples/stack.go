package examples

// Stack implementation borrowed from https://dev.to/jpoly1219/stacks-in-go-54k
type Stack struct {
	items []int
}

// Push pushes data in the stack.
// @ensures [stack grows by one element] @old{len(s.items)} == len(s.items) - 1
// @ensures [new data is on top of the stack] value, err := s.Top(); err == nil && value == data
func (s *Stack) Push(data int) {
	// implementation
}

// Pop pops the top element from the stack; returns no nil error if call on empty stack.
// @ensures [if no empty then stack shrinks by one element] !s.IsEmpty() ==> @old{len(s.items)} == len(s.items) + 1
// @ensures [if call on empty then error] s.IsEmpty() ==> err != nil
func (s *Stack) Pop() (err error) {
	// implementation
}

// Top yields the data on the top of the stack; error if the stack is empty.
// @ensures [if not empty the resulting data is the top of the stack and no error] !s.IsEmpty() ==> data == s.items[len(s.items)-1] && err == nil
// @ensures [if empty then error] s.IsEmpty() ==> err != nil
// @ensures [stack top element is not modified] @old{s.items[len(s.items)-1]} == s.items[len(s.items)-1]
func (s *Stack) Top() (data int, err error) {
	// implementation
}

// IsEmpty returns true if the stack has no data, false otherwise.
// @ensures [if empty then true] @old{len(s.items)} == 0 ==> result == true
// @ensures [if not empty then false] @old{len(s.items)} != 0 ==> result == false
// @unmodified [stack size is unmodified] len(s.items)
func (s *Stack) IsEmpty() (result bool) {
	// implementation
}
