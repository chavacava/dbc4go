package testdata

// @ensures something: @exists e @in r: e > 0
func foo(a []int) (r []int) {
	// implementation
}

// @ensures @exists e @in r: e > 0 ==> e % 2 == 0
func bar(a []int) (r []int) {
	// implementation
}
