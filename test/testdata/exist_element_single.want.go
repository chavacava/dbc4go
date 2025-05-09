// Code generated by dbc4go, DO NOT EDIT.
package testdata

// @ensures something: @exists e @in r: e > 0
func foo(a []int) (r []int) {
	{ // Open contract scope
		// Function's contracts
		defer func() {
			{
				cond := func() bool {
					for _, e := range r {
						cond := func() bool { return e > 0 }
						if cond() {
							return true
						}
					}
					return false
				}
				if !cond() {
					panic("function didn't satisfy something")
				}
			}
		}()
	} // Close contract scope

	// implementation
}

// @ensures @exists e @in r: e > 0 ==> e % 2 == 0
func bar(a []int) (r []int) {
	{ // Open contract scope
		// Function's contracts
		defer func() {
			{
				cond := func() bool {
					for _, e := range r {
						cond := func() bool {
							cond1 := func() bool { return e > 0 }
							cond2 := func() bool { return e%2 == 0 }
							return !cond1() || cond2()
						}
						if cond() {
							return true
						}
					}
					return false
				}
				if !cond() {
					panic("function didn't satisfy @exists e @in r: e > 0 ==> e % 2 == 0")
				}
			}
		}()
	} // Close contract scope

	// implementation
}
