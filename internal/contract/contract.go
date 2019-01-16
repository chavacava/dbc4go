//go:generate dbc4go -i $GOFILE -o $GOFILE
package contract

import (
	"go/ast"
	"regexp"
	"strings"
)

// FuncContract represents a contract associated to a function
type FuncContract struct {
	requires []Requires
	ensures  []Ensures
	target   *ast.FuncDecl
}

// NewFuncContract creates a FuncContract
//@requires target != nil
//@ensures c.target == target
//@ensures len(c.requires) == 0
//@ensures len(c.ensures) == 0
func NewFuncContract(target *ast.FuncDecl) (c FuncContract) {
	return FuncContract{requires: []Requires{}, ensures: []Ensures{}, target: target}
}

//Target yields the function declaration to which this contract is attached to
//@ensures t == c.target
func (c *FuncContract) Target() (t *ast.FuncDecl) {
	return c.target
}

// AddRequires adds a requires to this contract
//@ensures len(c.requires) == len(@old(c.requires)) + 1
//@ensures c.requires[len(c.requires)-1] == r
func (c *FuncContract) AddRequires(r Requires) {
	c.requires = append(c.requires, r)
}

// Requires yields requires clauses of this contract
//@ensures len(r) == len(c.requires)
func (c *FuncContract) Requires() (r []Requires) {
	return c.requires
}

// AddEnsures adds a ensures to this contract
//@ensures len(c.ensures) == len(@old(c.ensures)) + 1
//@ensures c.ensures[len(c.ensures)-1] == e
func (c *FuncContract) AddEnsures(e Ensures) {
	c.ensures = append(c.ensures, e)
}

// Ensures yields ensures clauses of this contract
//@ensures len(r) == len(c.ensures)
func (c *FuncContract) Ensures() (r []Ensures) {
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
	return rewriteImpliesExpr(r.expr)
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

var re4old = regexp.MustCompile(`@old\(([^\)]+)\)`)

// ExpandedExpression yields the expanded ensures' expression
func (r Ensures) ExpandedExpression() string {
	expr := rewriteImpliesExpr(r.expr)

	// replace @old(id.otherId) by old_id.otherId
	return re4old.ReplaceAllString(expr, `old_$1`)
}

func (r Ensures) String() string {
	return "@ensures " + r.expr
}

var reImplies = regexp.MustCompile(`(.*)==>(.*)`)

// rewriteImpliesExpr transforms p ==> q into its canonical form !p || q
func rewriteImpliesExpr(expr string) string {
	impExp := reImplies.FindAllStringSubmatch(expr, -1)
	if impExp == nil {
		return expr // no ==> operator in the expression, nothing to do
	}

	p := strings.Trim(impExp[0][1], " ")
	q := strings.Trim(impExp[0][2], " ")

	return "!(" + p + ") || (" + q + ")"
}
