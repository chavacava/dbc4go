package testdata

// @invariant dummy invariant: /
// true == true
type multilineType struct{}

func (m multilineType) foo() {}

// Contract:
//   - invariant dummy invariant: /
//     true == true
type multilineType2 struct{}

func (m multilineType2) foo() {}

// @requires dummy requirement: /
//
//	true == true
func multilineRaw() {

}

// Contract:
//   - ensures dummy clause: /
//     true == false
func multilineStandard() {

}
