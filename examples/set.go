package examples

// Borrowed from https://www.bytesizego.com/blog/set-in-golang

// Set is a collection of unique elements
type Set struct {
	elements map[string]struct{}
}

// NewSet creates a new set
// @ensures [set is empty] len(elements) == 0
func NewSet() (s *Set) {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

// Add inserts an element into the set
// @ensure [element is present in the set] s.Contains(value)
// @ensure [cardinality grows if new element] @old{s.Contains(value)} ==> @old{len(s.elements)+1} == len(s.elements)
// @ensure [cardinality remains the same if no new element] @old{!s.Contains(value)} ==> @old{len(s.elements)} == len(s.elements)
func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
// @ensure [element is not present in the set] !s.Contains(value)
// @ensure [cardinality shirks if element was in the set] @old{s.Contains(value)} ==> @old{len(s.elements)-1} == len(s.elements)
// @ensure [cardinality remains the same if element was not in the set] @old{!s.Contains(value)} ==> @old{len(s.elements)} == len(s.elements)
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
// @ensure _, ok := s.elements[value]; ok ==> result == true
// @ensure _, ok := s.elements[value]; !ok ==> result != true
// @unmodified len(s.elements)
func (s *Set) Contains(value string) (result bool) {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
// @ensures result == @old{len(s.elements)}
// @unmodified len(s.elements)
func (s *Set) Size() (result int) {
	return len(s.elements)
}

// List returns all elements in the set as a slice
// @ensure len(result) == @old{len(s.elements)}
// @unmodified len(s.elements)
func (s *Set) List() (result []string) {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
