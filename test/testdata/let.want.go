// Code generated by dbc4go, DO NOT EDIT.
package testdata

// @let foo := 1
// @ensures foo == 1
func letRaw() {
	{ // Open contract scope
		// Function's contracts
		foo := 1 //  defined with @let
		defer func() {
			if !(foo == 1) {
				panic("function didn't ensure foo == 1")
			}
		}()
	} // Close contract scope
}

// Contract:
//   - let foo := 1
//   - ensures foo == 1
func letStandard() {
	{ // Open contract scope
		// Function's contracts
		foo := 1 //  defined with @let
		defer func() {
			if !(foo == 1) {
				panic("function didn't ensure foo == 1")
			}
		}()
	} // Close contract scope
}