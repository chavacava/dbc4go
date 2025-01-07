package examples

// Borrowed from https://www.bytesizego.com/blog/set-in-golang

// Set is a collection of unique elements.
//
// @invariant Set.elements != nil
type Set struct {
	elements map[string]struct{}
}

// NewSet creates a new set.
//
// @ensures set is created and empty: s.elements != nil && len(s.elements) == 0
func NewSet() (s *Set) {
	return &Set{
		elements: map[string]struct{}{},
	}
}

// Add inserts an element into the set.
//
// @ensures element is present in the set: s.Contains(value)
// @let alreadyPresent := s.Contains(value)
// @ensures cardinality grows if new element: alreadyPresent ==> @old{len(s.elements)} == len(s.elements) - 1
// @ensures cardinality remains the same if no new element: alreadyPresent ==> @old{len(s.elements)} == len(s.elements)
func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set.
//
// @ensures element is not present in the set: !s.Contains(value)
// @let alreadyPresent := s.Contains(value)
// @ensures cardinality shirks if element was in the set: alreadyPresent ==> @old{len(s.elements)-1} == len(s.elements)
// @ensures cardinality remains the same if element was not in the set: alreadyPresent ==> @old{len(s.elements)} == len(s.elements)
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set.
//
// @ensures _, ok := s.elements[value]; ok ==> result == true
// @ensures _, ok := s.elements[value]; !ok ==> result != true
// @unmodified len(s.elements)
func (s *Set) Contains(value string) (result bool) {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set.
//
// @ensures result == @old{len(s.elements)}
// @unmodified len(s.elements)
func (s *Set) Size() (result int) {
	return len(s.elements)
}

// List returns all elements in the set as a slice.
//
// @ensures len(result) == @old{len(s.elements)}
// @unmodified len(s.elements)
func (s *Set) List() (result []string) {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
