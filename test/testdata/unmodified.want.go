// Code generated by dbc4go, DO NOT EDIT.
package foo

var x int

type S struct {
	a string
}

// @unmodified x, s.a
func (s S) unmodified1(a int, b int) (r int) {
	{ // Open contract scope
		// Function's contracts
		old_1 := x
		old_2 := s.a
		defer func() {
			if !(old_1 == x) {
				panic("function didn't ensure x unmodified")
			}
			if !(old_2 == s.a) {
				panic("function didn't ensure s.a unmodified")
			}
		}()
	} // Close contract scope

	return a
}

// Contract:
//   - unmodified x, s.a
func (s S) unmodified2(a int, b int) (r int) {
	{ // Open contract scope
		// Function's contracts
		old_1 := x
		old_2 := s.a
		defer func() {
			if !(old_1 == x) {
				panic("function didn't ensure x unmodified")
			}
			if !(old_2 == s.a) {
				panic("function didn't ensure s.a unmodified")
			}
		}()
	} // Close contract scope

	return a
}
