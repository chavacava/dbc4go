package foo

var x int

type S struct {
	a string
}

// @unmodified x, s.a
func (s S) unmodified1(a int, b int) (r int) {
	return a
}

// Contract:
//   - unmodified x, s.a
func (s S) unmodified2(a int, b int) (r int) {
	return a
}
