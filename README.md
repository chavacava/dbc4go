# dbc4go

Design by Contract&trade; for GO is a code generator that takes GO code annotated with contracts and generates code that enforces those contracts.  

Contracts are embedded into comments, therefore code annotated with contracts is valid GO code.

This project uses contracts itself! Check the source code and the `Makefile` to see how.

# Current State

This project is in a pre-ALPHA state.
Syntax of contracts will evolve in future versions.

## Available directives to write contracts

### `@requires`

Describes **preconditions** imposed by functions/methods.     

Syntax:

`@requires` _GO Boolean expression_

As you can see in the example below, the _GO Boolean expression_ must be a valid GO boolean expression as it will be used as the condition in a `if-then` statement.

The expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions)

Example:

```go
const maxAuthorizedSpeed = 350

// NewCar returns a Car struct
//@requires wheels > 2
//@requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
//@requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//@requires manufacturer != ""
func NewCar(wheels int, wheelsDrive int, maxSpeedKmh int, manufacturer string) Car { ... }
```

`dbc4go` will generate the following code

```go
// NewCar returns a Car struct
//@requires wheels > 2
//@requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
//@requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//@requires manufacturer != ""
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

	...
}
```

### `@ensures`

Describes **postconditions** of a function/method.     

Syntax:

`@ensures` _GO Boolean expression_

The expression can make reference to any identifier available in the scope at the beginning of the annotated function (for example: function parameters, method receiver, global variables, other functions)

Example:

```go
// accelerate the car
//@requires delta > 0
//@ensures c.speed > 0
//@ensures c.speed <= c.maxSpeedKmh
func (c *Car) accelerate(delta int) { ... }
```

## Planned directives to write contracts

### ==>
The future implementation of the operator `==>` will allow to write more precise contracts like

```go
// accelerate the car
//@requires delta > 0
//@ensures c.speed + delta >= c.maxSpeedKmh ==> c.speed' == c.maxSpeedKmh 
//@ensures c.speed + delta < c.maxSpeedKmh ==> c.speed' == c.speed + delta
func (c *Car) accelerate(delta int) { ... }
```

where `c.speed'` refers to the value of `c.speed` at the beginning of the method execution.

Notice that until the implementation of the `==>` operator, it is possible to express the same contracts by using canonical forms: `p ==> q` can be written as `!p || q`.

### `@invariant`

Defines an invariant property of a `struct`ure.

Syntax:

`@invariant` _GO Boolean expression_

The expression can make reference to any _private_ field identifier of the `struct` to which the invariant applies to.

Examples:

```go
// Car data-model
//@invariant speed <= maxSpeedKmh
type Car struct {
        maxSpeedKmh int
        speed       int
        // other fields ...
}
```

```go
// BankAccount data-model
//@invariant balance >= 0
type BankAccount struct {
        balance float
        Owner   string
        // other fields ...
}
```
