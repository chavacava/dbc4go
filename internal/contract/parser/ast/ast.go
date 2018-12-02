//go:generate dbc4go -i $GOFILE -o $GOFILE
package ast

import (
	"go/ast"
)

// Contract is the root of the AST of contracts
type Contract struct {
	requires []Requires
	target   *ast.FuncDecl
}

// NewContract creates a Contract
//@requires target != nil
func NewContract(target *ast.FuncDecl) Contract {
	return Contract{requires: []Requires{}, target: target}
}

// AddRequires adds a requires to this contract
func (c *Contract) AddRequires(r Requires) {
	c.requires = append(c.requires, r)
}

// Requires yields requires clauses of this contract
func (c *Contract) Requires() []Requires {
	return c.requires
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
