[![Build Status](https://travis-ci.com/chavacava/dbc4go.svg?branch=master)](https://travis-ci.com/chavacava/dbc4go)
[![Go Report Card](https://goreportcard.com/badge/github.com/chavacava/dbc4go)](https://goreportcard.com/report/github.com/chavacava/dbc4go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# dbc4go

[Design by Contract&trade;](https://en.wikipedia.org/wiki/Design_by_contract) for GO is a code generator that takes GO code annotated with contracts and generates code that enforces those contracts.  

<p align="center">
  <img src="./assets/mascots.jpg" alt="" width="300">
  <br>
  Logo by Eze
</p>


Contracts are embedded into comments, therefore code annotated with contracts is valid GO code.

This project uses contracts itself! Check the source code and the `Makefile` to see how.

# Usage

```
Usage of dbc4go:
  -i string
        input source file (defaults to stdin)
  -o string
        output file (defaults to stdout)
```

# Current State

This project is in a pre-ALPHA state.
Syntax of contracts might evolve in future versions.

## Available directives to write contracts

### `@requires`

Describes **pre-conditions** imposed by functions/methods.     

Syntax:

`@requires` _GO Boolean expression_

As you can see in the example below, the _GO Boolean expression_ must be a valid GO boolean expression as it will be used as the condition in an `if-then` statement.

The expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions)

Example:

```go
const maxAuthorizedSpeed = 350

// NewCar returns a Car struct
// @requires wheels > 2
// @requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
// @requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
// @requires manufacturer != ""
func NewCar(wheels int, wheelsDrive int, maxSpeedKmh int, manufacturer string) Car { ... }
```

`dbc4go` will generate the following code

```go
// NewCar returns a Car struct
// @requires wheels > 2
// @requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
// @requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
// @requires manufacturer != ""
func NewCar(wheels int, wheelsDrive int, maxSpeedKmh int, manufacturer string) Car {
        if !(manufacturer != "") {
                panic("precondition manufacturer != \"\" not satisfied")
        }
        if !(maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed) {
                panic("precondition maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed not satisfied")
        }
        if !(wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0) {
                panic("precondition wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0 not satisfied")
        }
        if !(wheels > 2) {
                panic("precondition wheels > 2 not satisfied")
        }

	// ...
}
```

### `@ensures`

Describes **post-conditions** of a function/method.     

Syntax:

`@ensures` _GO Boolean expression_

The expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions)

Expressions in `@ensure` clauses can use the `@old` operator to refer to the state of variables as it was just before the execution of the function.

Example:

```go
// accelerate the car
// @requires delta > 0
// @requires c.speed + delta <= c.maxSpeedKmh
// @ensures c.speed == @old{c.speed}+delta
func (c *Car) Accelerate(delta int) { ... }
```

where `@old{c.speed}` refers to the value of `c.speed` at the beginning of the method execution.

Limitation: only one `@old` expression per directive.

### The ==> operator
The `==>` operator (implication) allows to write more precise and concise contracts like

```go
// accelerate the car
// @requires delta > 0
// @ensures @old{c.speed} + delta >= c.maxSpeedKmh ==> c.speed == c.maxSpeedKmh 
func (c *Car) accelerate(delta int) { ... }
```

### `@invariant`

Defines an invariant property of a `struct`.
Invariants are enforced, materialized as `@ensure` clauses, on every method attached to the struct.

Syntax:

`@invariant` _GO Boolean expression_

The expression can make reference to any _private_ field identifier of the `struct` to which the invariant applies to.
References to field must be a selector expressions of the form `<struct-name>.<field-name>` 

Examples:

```go
// Car data-model
// @invariant Car.speed <= Car.maxSpeedKmh
type Car struct {
        maxSpeedKmh int
        speed       int
        // other fields ...
}
```

```go
// BankAccount data-model
// @invariant BankAccount.balance >= 0
type BankAccount struct {
        balance float
        // other fields ...
}
```

### `@unmodified`

Enforces the function keeps unmodified the given list of identifiers.

Syntax:

`@unmodified` _identifiers list_

Example:

```go
// accelerate the car
// @requires delta > 0
// @ensures @old{c.speed} + delta >= c.maxSpeedKmh ==> c.speed == c.maxSpeedKmh 
// @unmodified c.wheels, c.wheelsDrive, c.maxSpeedKmh, c.manufacturer
func (c *Car) Accelerate(delta int) { ... }
```
`@unmodified` is just syntax sugar for simplify writing `@ensures id == @old{id}`

### `@import`

If in your contracts you need to use a package that is not imported by the original source code, then you can import the package with the `@import` clause.

Syntax:

`@import` _pakage name_

Example:

```go
// Add element e to container
// @import strings
// @requires strings.HasPrefix(e, "my")
func (c *Container) Add(e string) { ... }
```
