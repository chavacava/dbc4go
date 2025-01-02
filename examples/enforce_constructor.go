package examples

// This is an example on how to enforce the use of a struct constructor

// MyObject is a dummy struct.
//
// @invariant MyObject.usedConstructor == true
type MyObject struct {

	// ... struct fields ...

	usedConstructor bool // created with the constructor?
}

// New creates a new MyObject with the given parameters.
//
// @ensures result.usedConstructor == true
func New( /* parameters */ ) (result MyObject) {
	return MyObject{
		/* ... set fields ...*/
		usedConstructor: true,
	}
}

// Methods of MyObject will enforce usedConstructor to be true,
// and fail if not the case.

// DoSomething does something
func (o *MyObject) DoSomething() {
	// ... implementation ...
}
