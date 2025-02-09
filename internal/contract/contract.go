// Package contract provides contract related objects
package contract

// disabled until support for generics go:generate dbc4go -i $GOFILE -o $GOFILE

import (
	"fmt"
	"go/ast"
	"maps"
	"regexp"
	"strings"
)

// OldCounter the counter to use when creating old_ parameters
// TODO: use a more elegant approach
var OldCounter = 0

// TypeContract represents a contract associated to a type.
// Typically a @invariant contract.
//
// Contract:
//   - invariant TypeContract.ensures != nil
//   - invariant TypeContract.requires != nil
//   - invariant TypeContract.imports != nil
//   - invariant TypeContract.targetTypeName != nil
type TypeContract struct {
	ensures        []Ensures
	requires       []Requires
	imports        map[string]struct{}
	targetTypeName string
}

// NewTypeContract creates a TypeContract.
//
// Contract:
//   - requires target != ""
//   - ensures c.targetTypeName == target
//   - ensures c.ensures != nil && len(c.ensures) == 0
//   - ensures c.requires != nil && len(c.requires) == 0
//   - ensures c.imports != nil && len(c.imports) == 0
func NewTypeContract(target string) (c *TypeContract) {
	return &TypeContract{
		ensures:        []Ensures{},
		requires:       []Requires{},
		targetTypeName: target,
		imports:        map[string]struct{}{},
	}
}

// IsEmpty returns true if this contract doesn't have requires nor ensure clauses, false otherwise.
func (c *TypeContract) IsEmpty() bool {
	return len(c.ensures) == 0 && len(c.requires) == 0
}

// AddEnsures adds a ensures to this contract.
//
// Contract:
//   - requires e != nil
//   - ensures len(c.ensures) == @old{len(c.ensures})} + 1
//   - ensures c.ensures[len(c.ensures)-1] == e
func (c *TypeContract) AddEnsures(e Ensures) *TypeContract {
	c.ensures = append(c.ensures, e)
	return c
}

// Ensures yields ensures clauses of this contract.
//
// Contract:
//   - ensures len(r) == len(c.ensures)
func (c *TypeContract) Ensures() (r []Ensures) {
	return c.ensures
}

// AddRequires adds a requires to this contract.
//
// Contract:
//   - ensures len(c.requires) == @old{len(c.requires)} + 1
//   - ensures c.requires[len(c.requires)-1] == r
func (c *TypeContract) AddRequires(r Requires) *TypeContract {
	c.requires = append(c.requires, r)
	return c
}

// Requires yields requires clauses of this contract.
//
// Contract:
//   - ensures len(r) == len(c.requires)
func (c *TypeContract) Requires() (r []Requires) {
	return c.requires
}

// AddImport adds an import to this contract.
//
// Contract:
//   - requires path != ""
//   - ensures len(c.imports) >= @old{len(c.imports)}
func (c *TypeContract) AddImport(path string) *TypeContract {
	c.imports[strings.Trim(path, "\"")] = struct{}{}
	return c
}

// Imports returns imports required by this contract.
//
// Contract:
//   - ensures len(r) == len(c.imports)
func (c *TypeContract) Imports() (r map[string]struct{}) {
	return c.imports
}

// FuncContract represents a contract associated to a function.
//
// Contract:
//   - invariant FuncContract.requires != nil
//   - invariant FuncContract.ensures != nil
//   - invariant FuncContract.import != nil
//   - invariant FuncContract.lets != nil
//   - invariant FuncContract.target != nil
type FuncContract struct {
	requires []Requires
	ensures  []Ensures
	imports  map[string]struct{}
	lets     []Let
	target   *ast.FuncDecl
}

// NewFuncContract creates a FuncContract.
//
// Contract:
//   - ensures len(c.requires) == 0
//   - ensures len(c.ensures) == 0
//   - ensures len(c.imports) == 0
//   - ensures len(c.lets) == 0
func NewFuncContract() (c *FuncContract) {
	return &FuncContract{
		requires: []Requires{},
		ensures:  []Ensures{},
		lets:     []Let{},
		imports:  map[string]struct{}{},
	}
}

// IsEmpty returns true if this contract doesn't have requires nor ensure clauses, false otherwise.
func (c *FuncContract) IsEmpty() bool {
	return len(c.ensures) == 0 && len(c.requires) == 0
}

// AddRequires adds a requires to this contract.
// Contract:
//   - ensures len(c.requires) == @old{len(c.requires)} + 1
//   - ensures c.requires[len(c.requires)-1] == r
func (c *FuncContract) AddRequires(r Requires) *FuncContract {
	c.requires = append(c.requires, r)
	return c
}

// Requires yields requires clauses of this contract.
// Contract:
//   - ensures len(r) == len(c.requires)
//   - unmodified len(c.requires)
func (c *FuncContract) Requires() (r []Requires) {
	return c.requires
}

// AddEnsures adds a ensures to this contract.
//
// Contract:
//   - ensures len(c.ensures) == @old{len(c.ensures)} + 1
//   - ensures c.ensures[len(c.ensures)-1] == e
func (c *FuncContract) AddEnsures(e Ensures) *FuncContract {
	c.ensures = append(c.ensures, e)
	return c
}

// Ensures yields ensures clauses of this contract.
//
// Contract:
//   - ensures len(r) == len(c.ensures)
//   - unmodified len(c.ensures)
func (c *FuncContract) Ensures() (r []Ensures) {
	return c.ensures
}

// AddImport adds an import to this contract.
//
// Contract:
//   - requires path != ""
//   - ensures len(c.imports) >= @old{len(c.imports)}
func (c *FuncContract) AddImport(path string) *FuncContract {
	c.imports[strings.Trim(path, "\"")] = struct{}{}
	return c
}

// Imports returns imports required by this contract.
//
// Contract:
//   - ensures len(r) == len(c.imports)
func (c *FuncContract) Imports() map[string]struct{} {
	return c.imports
}

// AddLet adds a Let to this contract.
//
// Contract:
//   - ensures len(c.lets) == @old{len(c.lets)} + 1
//   - ensures c.lets[len(c.lets)-1] == l
func (c *FuncContract) AddLet(l Let) *FuncContract {
	c.lets = append(c.lets, l)
	return c
}

// Lets yields lets clauses of this contract.
// Contract:
//   - ensures len(r) == len(c.lets)
func (c *FuncContract) Lets() (r []Let) {
	return c.lets
}

// Let models a @Let clause.
type Let struct {
	expr        Expression
	description string
}

// NewLet creates a Let object
//
// Contract:
//   - requires expr != ""
func NewLet(expr Expression, description string) Let {
	return Let{expr: expr, description: description}
}

// Expression yields the let expression
func (l Let) Expression() Expression {
	return l.expr
}

// ExpandedExpression yields the expanded let expression
func (l Let) ExpandedExpression() Expression {
	return l.expr
}

// Description yields the let description
func (l Let) Description() string {
	return l.description
}

// Requires is a @requires clause of a contract
type Requires struct {
	expr        Expression
	description string
}

// NewRequires creates a Requires object.
//
// Contract:
//   - requires expr != ""
func NewRequires(expr Expression, description string) Requires {
	return Requires{expr: expr, description: description}
}

// Expression yields the expression of this Requires.
func (r Requires) Expression() (expr Expression) {
	return r.expr
}

// Description yields the description of this Requires.
func (r Requires) Description() (description string) {
	return r.description
}

// ExpandedExpression yields the expanded requires' expression.
//
// Contract:
//   - ensures result != ""
func (r Requires) ExpandedExpression() (result Expression) {
	switch r.expr.Kind {
	case ExprKindPlain:
		return Expression{
			Kind:     r.expr.Kind,
			SubExprs: r.expr.SubExprs,
			Raw:      rewriteImpliesExpr(r.expr.Raw),
		}
	default:
		panic(fmt.Sprintf("Unexpected expression kind %d", r.expr.Kind))
	}
}

// Ensures is a @ensures clause of a contract.
//
// Contract:
//   - invariant Ensures.expr != ""
type Ensures struct {
	expr        Expression
	description string
}

// NewEnsures creates a Ensures object.
//
// Contract:
//   - requires expr != ""
//   - ensures r.expr == expr
//   - ensures r.description == description
func NewEnsures(expr Expression, description string) (r Ensures) {
	return Ensures{expr: expr, description: description}
}

// Expression yields the expression of this Ensures.
func (r Ensures) Expression() (expr Expression) {
	return r.expr
}

// Description yields the description of this Ensures.
func (r Ensures) Description() (description string) {
	return r.description
}

// ExpandedExpression yields the expanded ensures' expression.
//
// Contract:
//   - ensures expr != ""
//   - ensures idToOldIdMap != nil
func (r Ensures) ExpandedExpression() (shortStmt, expr string, idToOldIDMap map[string]string) {
	expr = r.expr.Raw
	shortStmt = ""
	if strings.Contains(expr, ";") {
		parts := strings.SplitN(expr, ";", 2)
		shortStmt, expr = parts[0], parts[1]
	}
	expr = rewriteImpliesExpr(expr)

	// replace @old{id.otherId} by old_<number> in short-statement
	shortStmt, shortStmtMappings := expandOldExpressions(shortStmt)
	// replace @old{id.otherId} by old_<number> in expression
	expr, exprMappings := expandOldExpressions(expr)
	idToOldIDMap = map[string]string{}
	maps.Copy[map[string]string, map[string]string](idToOldIDMap, shortStmtMappings)
	maps.Copy[map[string]string, map[string]string](idToOldIDMap, exprMappings)

	return shortStmt, expr, idToOldIDMap
}

func ExpandEnsuresExpression(expression Expression) (shortStmt, expr string, idToOldIDMap map[string]string) {
	expr = expression.Raw
	shortStmt = ""
	if strings.Contains(expr, ";") {
		parts := strings.SplitN(expr, ";", 2)
		shortStmt, expr = parts[0], parts[1]
	}
	expr = rewriteImpliesExpr(expr)

	// replace @old{id.otherId} by old_<number> in short-statement
	shortStmt, shortStmtMappings := expandOldExpressions(shortStmt)
	// replace @old{id.otherId} by old_<number> in expression
	expr, exprMappings := expandOldExpressions(expr)
	idToOldIDMap = map[string]string{}
	maps.Copy[map[string]string, map[string]string](idToOldIDMap, shortStmtMappings)
	maps.Copy[map[string]string, map[string]string](idToOldIDMap, exprMappings)

	return shortStmt, expr, idToOldIDMap
}

// Re4old is a regular expression for @old expressions.
var Re4old = regexp.MustCompile(`@old\{(.+)\}`)

// Replace @old{<expression>} by old_<number> in the given string.
// It also returns the mapping between the <expression> and old_<number>.
//
// Contract:
//   - requires expr != ""
//   - ensures expanded != ""
//   - ensures mapping != nil
//   - ensures expanded == expr ==> len(mapping) == 0
func expandOldExpressions(expr string) (expanded string, mapping map[string]string) {
	exp2id := map[string]string{}
	matches := Re4old.FindAllStringSubmatch(expr, -1)
	for _, m := range matches {
		oldAsID := oldID()
		expr = strings.Replace(expr, m[0], oldAsID, 1)
		exp2id[m[1]] = oldAsID
	}
	return expr, exp2id
}

// Contract:
//   - ensures OldCounter == @old{OldCounter} + 1
//   - ensures r != ""
func oldID() (r string) {
	OldCounter++
	return fmt.Sprintf("old_%d", OldCounter)
}

var reImplies = regexp.MustCompile(`(.*)==>(.*)`)

// rewriteImpliesExpr transforms p ==> q into its canonical form !p || q.
//
// Contract:
//   - import strings
//   - let mustRewrite == strings.Contains(expr,"==>")
//   - ensures !mustRewrite ==> result == expr
//   - ensures mustRewrite ==> result != expr
func rewriteImpliesExpr(expr string) (result string) {
	impExp := reImplies.FindAllStringSubmatch(expr, -1)
	if impExp == nil {
		return expr // no ==> operator in the expression, nothing to do
	}

	p := strings.Trim(impExp[0][1], " ")
	q := strings.Trim(impExp[0][2], " ")

	return "!(" + p + ") || (" + q + ")"
}

type ExpressionKind byte

type Expression struct {
	Kind     ExpressionKind
	SubExprs map[string]Expression
	Raw      string
}

const (
	ExprKindPlain  ExpressionKind = 0
	ExprKindForall ExpressionKind = 1
)

const ExprKindForallFieldExpression = "expression"
const ExprKindForallFieldKind = "kind"
const ExprKindForallFieldSources = "sources"
const ExprKindForallFieldVariables = "variables"
const ForallKindIn = "in"
const ForallKindIndexof = "indexof"
