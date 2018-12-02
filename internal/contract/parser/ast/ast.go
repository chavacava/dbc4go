//go:generate dbc4go -i $GOFILE -o $GOFILE
package ast

import (
	"go/ast"
)

// Contract is the root of the AST of contracts
type Contract struct {
	requires []Requires
	ensures  []Ensures
	target   *ast.FuncDecl
}

// NewContract creates a Contract
//@requires target != nil
//@ensures c.target == target
func NewContract(target *ast.FuncDecl) (c Contract) {
	return Contract{requires: []Requires{}, ensures: []Ensures{}, target: target}
}

// AddRequires adds a requires to this contract
func (c *Contract) AddRequires(r Requires) {
	c.requires = append(c.requires, r)
}

// Requires yields requires clauses of this contract
func (c *Contract) Requires() []Requires {
	return c.requires
}

// AddEnsures adds a ensures to this contract
func (c *Contract) AddEnsures(r Ensures) {
	c.ensures = append(c.ensures, r)
}

// Ensures yields ensures clauses of this contract
func (c *Contract) Ensures() []Ensures {
	return c.ensures
}

// Requires is a @requires clause of a contract
type Requires struct {
	expr string
}

// NewRequires creates a Requires object
//@requires expr != ""
func NewRequires(expr string) Requires {
	return Requires{expr: expr}
}

// ExpandedExpression yields the expanded requires' expression
func (r Requires) ExpandedExpression() string {
	return r.expr
}

func (r Requires) String() string {
	return "@requires " + r.expr
}

// Ensures is a @ensures clause of a contract
type Ensures struct {
	expr string
}

// NewEnsures creates a Ensures object
//@ensures expr != ""
func NewEnsures(expr string) Ensures {
	return Ensures{expr: expr}
}

// ExpandedExpression yields the expanded ensures' expression
func (r Ensures) ExpandedExpression() string {
	return r.expr
}

func (r Ensures) String() string {
	return "@ensures " + r.expr
}
