[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# dbc4go

[Design by Contract&trade;](https://en.wikipedia.org/wiki/Design_by_contract) for GO is a code generator that takes GO code annotated with contracts and generates code that enforces those contracts at runtime.  
Contracts are embedded into comments, therefore code annotated with contracts is still valid GO code.

<p align="center">
  <img src="./assets/mascots.jpg" alt="" width="300">
  <br>
  Logo by Eze
</p>


A simple example: imagine you have a `counter` type like the following

```go
type Counter struct {
	value int
}

// IncrementBy increments the counter value by n.
func (c *Counter) IncrementBy(n int) {
	c.value += n
}
```
You can add a contract to the `IncrementBy` method to enforce that the resulting value of the counter is actually incremented by `n` at the end of method's execution:

```go
// IncrementBy increments the counter value by n.
//
// Contract:
//   - ensures value is incremented by n: c.value == @old{c.value} + n
func (c *Counter) IncrementBy(n int) {
	c.value += n
}
```
As is, the contract you added has no effect on your code other than better documenting it.
But if you run `dbc4go` on the source file (let's say `counter.go`) 

```
$ dbc4go -i counter.go 
```
you will get the following code:

```go
type Counter struct {
	value int
}

// Increment increments the counter value by a n.
//
// Contract:
//   - ensures value is incremented by n: c.value == @old{c.value} + n
func (c *Counter) IncrementBy(n int) {
	{ // Open contract scope
		// Function's contracts
		old_1 := c.value
		defer func() {
			if !(c.value == old_1+n) {
				panic("function didn't ensure value is incremented by n")
			}
		}()
	} // Close contract scope

	c.value += n
}
```
As you can see, `dbc4go` instrumented the original code by adding a whole block of statements at the beginning of the method's body.
The new code block implements the checking of the contract you wrote in the documentation of the method.
Now, when the instrumented version of your code is executed, the contract will be enforced, and if the contract is not respected the method will panic.

For further information on how to write contracts you can read the section [Available contract clauses](#available-contract-clauses) below and check the [examples](./examples/).

You might still wonder how this can be useful for you.
Well, one thing you can do is execute your tests on the instrumented version of your code.
You could, for example, use `go generate` to obtain an instrumented version of your files and then run `go test`.
Like that, tests and contracts will combine to provide you more chances to find bugs. 

This project uses contracts itself! Check the source code and the `Makefile` to see how.

The article [Design by Contract](https://se.inf.ethz.ch/~meyer/publications/old/dbc_chapter.pdf) by Bertrand Meyer provides a complete and clear explanation of the idea. 

# Usage

```
Usage of dbc4go:
  -i string
        input source file (defaults to stdin)
  -o string
        output file (defaults to stdout)
```

## Available contract clauses

### `requires`

`requires` contract clauses describe **pre-conditions** imposed by functions/methods.     

Syntax:

`requires` _GO Boolean expression_

As you can see in the example below, the _GO Boolean expression_ must be a valid GO Boolean expression as it will be used as the condition in an `if-then` statement.

The expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions)

Contract clauses must be organized as a list under a contract declaration of the form `Contract:`.

Example:

```go
const maxAuthorizedSpeed = 350

// NewCar returns a Car struct.
//
// Contract:
//   - requires more than 2 wheels: wheels > 2
//   - requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0 
//   - requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//   - requires manufacturer != ""
func NewCar(wheels int, wheelsDrive int, maxSpeedKmh int, manufacturer string) Car { ... }
```

_Short-statements_ are supported as part of _GO Boolean expression_, therefore it's okay to write a contract like the following:

```go
// Accelerate the car.
//
// Contract:
//   - requires delta > 0
//   - requires targetSpeed := c.speed + delta; targetSpeed <= c.maxSpeedKmh
func (c *Car) Accelerate(delta int) { ... }
```
### `ensures`

`ensures` contract clauses describes **post-conditions** of a function/method.     

Syntax:

`ensures` [description:] _GO Boolean expression_

As for `ensure` clauses, the expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions).

Additionally, expressions in `ensure` clauses can use the `@old` operator to refer to the state of variables as it was just before the execution of the function.

Example:

```go
// Accelerate the car.
//
// Contract:
//   - requires delta > 0
//   - requires c.speed + delta <= c.maxSpeedKmh
//   - ensures c.speed == @old{c.speed}+delta
func (c *Car) Accelerate(delta int) { ... }
```

Limitations of `@old`: 
1. at most one `@old` expression in the short-statement and one in the Boolean expression.
2. `dbc4go` doesn't have type information of expressions inside an `@old` annotation therefore in some cases you need to add type casts to obtain compilable code. For example, to indicate that `@old{someExpression}` is of type `float` you must write `@old{someExpression}.(float)`

An alternative to `@old` annotations is the use of `let` clauses (see below) thus the previous example could be also written as:

```go
// Accelerate the car.
//
// Contract:
//   - requires delta > 0
//   - requires c.speed + delta <= c.maxSpeedKmh
//   - let initialSpeed := c.speed
//   - ensures c.speed == initialSpeed+delta
func (c *Car) Accelerate(delta int) { ... }
```
### `let`
Captures an expression's value into a variable **at the beginning of the function execution**.
The variable can be used in `ensures` and `requires`

Syntax:

`let` _id_ `:=` _expression_
 
_id_ must be a valid GO identifier and _expression_ a valid GO expression in the context of the annotated function.

### `invariant`

Defines an invariant property of a `struct`.
An invariant property is a property of the struct that should always hold.
Every method attached to the struct can assume the invariants hold and must preserve the invariants.
Invariants are enforced, materialized as `requires` and `ensures` clauses, on every method attached to the struct.

Syntax:

`invariant` _GO Boolean expression_

The expression can make reference to any identifier available in the scope of the struct declaration and any _private_ field identifier of the `struct` to which the invariant applies to.
References to fields must be a selector expression of the form `<struct-name>.<field-name>` 

Examples:

```go
// Car data-model.
//
// Contract:
//   - invariant Car.speed <= Car.maxSpeedKmh
//   - invariant Car.speed >= 0
type Car struct {
        maxSpeedKmh int
        speed       int
        // other fields ...
}
```

```go
// BankAccount data-model
//
// Contract:
//   - invariant BankAccount.balance >= 0
type BankAccount struct {
        balance float
        // other fields ...
}
```

### `unmodified`

Enforces the function keeps unmodified the given list of expressions.

Syntax:

`unmodified` _expressions list_

Example:

```go
// Accelerate the car.
//
// Contract:
//   - requires delta > 0
//   - requires c.speed + delta <= c.maxSpeedKmh
//   - let initialSpeed := c.speed
//   - ensures c.speed == initialSpeed+delta
//   - unmodified c.wheels, c.wheelsDrive, c.maxSpeedKmh, c.manufacturer
func (c *Car) Accelerate(delta int) { ... }
```
`unmodified` is just syntax sugar for simplify writing `ensures expr == @old{expr}`

### `import`

If in your contracts you need to use a package that is not imported by the original source code, then you can import the package with the `import` clause.

Syntax:

`import` _pakage name_

Example:

```go
// Add element e to container
//
// Contract:
//   - import strings
//   - requires strings.HasPrefix(e, "my")
func (c *Container) Add(e string) { ... }
```


### The ==> operator
The `==>` operator (implication) allows to write more precise and concise contracts like

```go
// Accelerate the car.
//
// Contract:
//   - requires delta > 0
//   - let initialSpeed := c.speed
//   - ensures initialSpeed + delta < c.maxSpeedKmh ==> c.speed == initialSpeed + delta
//   - ensures initialSpeed + delta >= c.maxSpeedKmh ==> c.speed == c.maxSpeedKmh 
func (c *Car) Accelerate(delta int) { ... }
```

### Contract Syntax 

`dbc4go` supports three contract syntaxes:
1. _standard_ syntax, the one introduced in the previous sections,
2. _raw_ syntax, and
2. _directive_ syntax 

All three syntaxes have equivalent expressiveness power. 

While the _raw_ syntax is easier/shorter to write, the _standard_ syntax lets GO tools and IDE to render function and types contracts in a nicer and readable form.

```go
// Contract in standard syntax

// NewCar returns a Car struct.
//
// Contract:
//  - requires wheels > 2
//  - requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
//  - requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//  - requires manufacturer != ""
func NewCar(...) {...}
```

```go
// Contract in raw syntax

// NewCar returns a Car struct.
//
// @requires wheels > 2
// @requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
// @requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
// @requires manufacturer != ""
func NewCar(...) {...}
```
_Raw_ syntax doesn't require a contract declaration, and contract clauses can be line-interleaved within non-contractual documentation.

_Directive_ syntax is useful in situations where you need to add a contract clause that will not render in the documentation.
For example, if a struct invariant refers to a private field and you don't want to leak the field's name in the documentation you can define the invariant using directive syntax:

```go
// Contract in directive syntax

// BankAccount represents a bank account.
//
//contract:invariant !BankAccount.closed ==> BankAccount.balance >= minBalance
//contract:invariant !BankAccount.closed ==> BankAccount.balance <= maxBalance
//contract:invariant BankAccount.closed ==> BankAccount.balance == 0
type BankAccount struct {
	balance int  // the balance of the account
	closed  bool // is the account closed?
}
```
(notice there is no blank space between the comment delimiter `//` and the contract clause)

#### Raw syntax summary

`@requires` [_description_:] _GO Boolean expression_

`@ensures` [_description_:] _GO Boolean expression_

`@let` _id_ `:=` _expression_

`@invariant` [_description_:] _GO Boolean expression_

`@unmodified` _identifiers list_

`@import` _pakage name_

You can **check [these examples](./examples/raw_contracts/)** of code annotated with the raw syntax.

#### Directive syntax summary

Contract clauses must respect the directive comment format as defined in the **Syntax** section of [Go Doc Comments](https://tip.golang.org/doc/comment) 

`contract:requires` [_description_:] _GO Boolean expression_

`contract:ensures` [_description_:] _GO Boolean expression_

`contract:let` _id_ `:=` _expression_

`contract:invariant` [_description_:] _GO Boolean expression_

`contract:unmodified` _identifiers list_

`contract:import` _pakage name_
