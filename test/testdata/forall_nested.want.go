// Code generated by dbc4go, DO NOT EDIT.
package testdata

// @ensures result is sorted: @forall i @indexof r: @forall j @indexof r: i <= j ==> r[i] <= r[j]
func sort(a []int) (r []int) {
	{ // Open contract scope
		// Function's contracts
		defer func() {
			cond := func() bool {
				for i := 0; i < len(r); i++ {
					cond := func() bool {
						cond := func() bool {
							for j := 0; j < len(r); j++ {
								cond := func() bool { return !(i <= j) || (r[i] <= r[j]) }
								if !cond() {
									return false
								}
							}
							return true
						}
						if !cond() {
							return false
						}
						return true
					}
					if !cond() {
						return false
					}
				}
				return true
			}
			if !cond() {
				panic("function didn't ensure @forall i @indexof r: @forall j @indexof r: i <= j ==> r[i] <= r[j]")
			}
		}()
	} // Close contract scope

	// implementation
}

// @ensures @forall i @indexof r: @forall j @indexof r: i != j ==> r[i] != r[j]
func deduplicate(a []int) (r []int) {
	{ // Open contract scope
		// Function's contracts
		defer func() {
			cond := func() bool {
				for i := 0; i < len(r); i++ {
					cond := func() bool {
						cond := func() bool {
							for j := 0; j < len(r); j++ {
								cond := func() bool { return !(i != j) || (r[i] != r[j]) }
								if !cond() {
									return false
								}
							}
							return true
						}
						if !cond() {
							return false
						}
						return true
					}
					if !cond() {
						return false
					}
				}
				return true
			}
			if !cond() {
				panic("function didn't ensure @forall i @indexof r: @forall j @indexof r: i != j ==> r[i] != r[j]")
			}
		}()
	} // Close contract scope

	// implementation
}
