package testdata

// @ensures result is sorted: @forall i @indexof r: @forall j @indexof r: i <= j ==> r[i] <= r[j]
func sort(a []int) (r []int) {
	// implementation
}

// @ensures @forall i @indexof r: @forall j @indexof r: i != j ==> r[i] != r[j]
func deduplicate(a []int) (r []int) {
	// implementation
}
