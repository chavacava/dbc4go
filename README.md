[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# dbc4go

[Design by Contract&trade;](https://en.wikipedia.org/wiki/Design_by_contract) for GO is a code generator that takes GO code annotated with contracts and generates code that enforces those contracts at runtime.  

<p align="center">
  <img src="./assets/mascots.jpg" alt="" width="300">
  <br>
  Logo by Eze
</p>


Contracts are embedded into comments, therefore code annotated with contracts is still valid GO code.

This project uses contracts itself! Check the source code and the `Makefile` to see how.

You can also **check the [examples](./examples/)**.

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
References to field must be a selector expressions of the form `<struct-name>.<field-name>` 

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

Enforces the function keeps unmodified the given list of identifiers.

Syntax:

`unmodified` _identifiers list_

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
`unmodified` is just syntax sugar for simplify writing `ensures id == @old{id}`

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

`dbc4go` supports two contract syntaxes:
1. _standard_ syntax, the one introduced in the previous sections, and
2. _raw_ syntax 

Both syntaxes have equivalent expressiveness power. 
While the _raw_ syntax is easier/shorter to write, the _standard_ syntax lets `go` tools to render function and types contracts in a nicer and readable form.

<table>
<tr>
<td> doc friendly contract </td> <td> Raw contract </td>
</tr>
<tr>
<td>

```go
// NewCar returns a Car struct.
//
// Contract:
//  - requires wheels > 2
//  - requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
//  - requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//  - requires manufacturer != ""
func NewCar(...) {...}
```
</td>
<td>

```go
// NewCar returns a Car struct.
//
// @requires wheels > 2
// @requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
// @requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
// @requires manufacturer != ""
func NewCar(...) {...}
```
</td>
</tr>
</table>

Raw syntax does not require a contract declaration, and contract clauses can be interleaved within non-contractual documentation.

#### Raw syntax summary

`@requires` [_description_:] _GO Boolean expression_

`@ensures` [_description_:] _GO Boolean expression_

`@let` _id_ `:=` _expression_

`@invariant` _GO Boolean expression_

`@unmodified` _identifiers list_

`@import` _pakage name_

You can **check [these examples](./examples/raw_contracts/)** of code annotated with the raw syntax.
