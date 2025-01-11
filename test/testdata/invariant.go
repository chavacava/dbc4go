package testdata

// @invariant a > 0
type foo struct {
	a int
}

func (f foo) f() {}

// Contract:
//   - invariant b != ""
type bar struct {
	b string
}

func (b bar) b() {}
