package examples

// This is an example on how to enforce the use of a struct constructor

type MyObject struct {

	// ... struct fields ...

	usedConstructor bool // created with the constructor?
}

// New creates a new MyObject with the given parameters
//@ensures result.usedConstructor
func New( /* parametres */ ) (result MyObject) {
	return MyObject{
		/* ... set fields ...*/
		usedConstructor: true,
	}
}

// Methods of MyObject require usedConstructor to be true,
// and fail if not the case.

// DoSomething does something
//@require o.usedConstructor
func (o *MyObject) DoSomething() {
	// ... implementation ...
}
