// Package contract provides contract related objects
package contract

//go:generate dbc4go -i $GOFILE -o $GOFILE

import (
	"go/ast"
	"regexp"
	"strings"
)

// TypeContract represents a contract associated to a type.
// Typically a @invariant contract
type TypeContract struct {
	ensures        []Ensures
	imports        map[string]struct{}
	targetTypeName string
}

// NewTypeContract creates a TypeContract
// @requires target != ""
// @ensures c.targetTypeName == target
// @ensures len(c.ensures) == 0
// @ensures len(c.imports) == 0
func NewTypeContract(target string) (c *TypeContract) {
	return &TypeContract{
		ensures:        []Ensures{},
		targetTypeName: target,
		imports:        map[string]struct{}{},
	}
}

// AddEnsures adds a ensures to this contract
// ensures len(c.ensures) == len(@old(c.ensures)) + 1
// @ensures c.ensures[len(c.ensures)-1] == e
func (c *TypeContract) AddEnsures(e Ensures) {
	c.ensures = append(c.ensures, e)
}

// Ensures yields ensures clauses of this contract
// @ensures len(r) == len(c.ensures)
func (c *TypeContract) Ensures() (r []Ensures) {
	return c.ensures
}

// AddImport adds an import to this contract
func (c *TypeContract) AddImport(path string) {
	c.imports[strings.Trim(path, "\"")] = struct{}{}
}

// Imports returns imports required by this contract
func (c *TypeContract) Imports() map[string]struct{} {
	return c.imports
}

// FuncContract represents a contract associated to a function
type FuncContract struct {
	requires []Requires
	ensures  []Ensures
	imports  map[string]struct{}
	target   *ast.FuncDecl
}

// NewFuncContract creates a FuncContract
// @requires target != nil
// @ensures c.target == target
// @ensures len(c.requires) == 0
// @ensures len(c.ensures) == 0
// @ensures len(c.imports) == 0
func NewFuncContract(target *ast.FuncDecl) (c *FuncContract) {
	return &FuncContract{requires: []Requires{}, ensures: []Ensures{}, target: target, imports: map[string]struct{}{}}
}

// Target yields the function declaration to which this contract is attached to
// @ensures t == c.target
func (c *FuncContract) Target() (t *ast.FuncDecl) {
	return c.target
}

// AddRequires adds a requires to this contract
// ensures len(c.requires) == len(@old(c.requires)) + 1
// @ensures c.requires[len(c.requires)-1] == r
func (c *FuncContract) AddRequires(r Requires) {
	c.requires = append(c.requires, r)
}

// Requires yields requires clauses of this contract
// @ensures len(r) == len(c.requires)
func (c *FuncContract) Requires() (r []Requires) {
	return c.requires
}

// AddEnsures adds a ensures to this contract
// ensures len(c.ensures) == len(@old(c.ensures)) + 1
// @ensures c.ensures[len(c.ensures)-1] == e
func (c *FuncContract) AddEnsures(e Ensures) {
	c.ensures = append(c.ensures, e)
}

// Ensures yields ensures clauses of this contract
// @ensures len(r) == len(c.ensures)
func (c *FuncContract) Ensures() (r []Ensures) {
	return c.ensures
}

// AddImport adds an import to this contract
func (c *FuncContract) AddImport(path string) {
	c.imports[strings.Trim(path, "\"")] = struct{}{}
}

// Imports returns imports required by this contract
func (c *FuncContract) Imports() map[string]struct{} {
	return c.imports
}

// Requires is a @requires clause of a contract
type Requires struct {
	expr        string
	description string
}

// NewRequires creates a Requires object
// @requires expr != ""
func NewRequires(expr, description string) Requires {
	return Requires{expr: expr, description: description}
}

// ExpandedExpression yields the expanded requires' expression
func (r Requires) ExpandedExpression() string {
	return rewriteImpliesExpr(r.expr)
}

func (r Requires) String() string {
	if r.description != "" {
		r.description += " "
	}
	return "@requires " + r.description + r.expr
}

// Ensures is a @ensures clause of a contract
type Ensures struct {
	expr        string
	description string
}

// NewEnsures creates a Ensures object
// @ensures expr != ""
func NewEnsures(expr, description string) Ensures {
	return Ensures{expr: expr, description: description}
}

var Re4old = regexp.MustCompile(`@old\(([^\)]+)\)`)

// ExpandedExpression yields the expanded ensures' expression
func (r Ensures) ExpandedExpression() (string, map[string]string) {
	expr := rewriteImpliesExpr(r.expr)

	idToOldID := map[string]string{}
	// replace @old(id.otherId) by old_id_otherId
	matches := Re4old.FindAllStringSubmatch(expr, -1)
	for _, m := range matches {
		oldAsID := oldID(m[1])
		expr = strings.Replace(expr, m[0], oldAsID, 1)
		idToOldID[m[1]] = oldAsID
	}

	return expr, idToOldID
}

func oldID(id string) string {
	id = strings.ReplaceAll(id, `.`, "_")
	return "old_" + id
}

func (r Ensures) String() string {
	if r.description != "" {
		r.description += " "
	}
	return "@ensures " + r.description + r.expr
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
