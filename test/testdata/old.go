package testdata

// @ensures @old{a} == a
func old1(a any) {

}

// @ensures @old{a.b} == a
func old2(a any) {

}

// @ensures @old{a.b + 1} == a
func old3(a any) {

}

// @ensures a == 0 ==> @old{a.b}
func old4(a any) {

}
