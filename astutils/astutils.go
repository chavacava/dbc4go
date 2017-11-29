// Package astutils provides a set of utility functions to build ASTs
package astutils

import (
	"go/ast"
	"go/token"
)

// NewID yields an identifier node
func NewID(id string) *ast.Ident {
	return &ast.Ident{Name: id, NamePos: 0, Obj: nil}
}

// NewBinExp yields a binary expression AST
func NewBinExp(op token.Token, lhs, rhs ast.Expr) *ast.BinaryExpr {
	return &ast.BinaryExpr{Op: op, OpPos: 0, X: lhs, Y: rhs}
}

// NewCall yields a function call AST
func NewCall(pkg, funcName string, args []ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: pkg, NamePos: 0, Obj: nil},
		Sel: &ast.Ident{Name: funcName, NamePos: 0, Obj: nil}},
		Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
}

// NewCallArgs yields a slice of expressions that can be used as arguments in
// a function call
func NewCallArgs(args ...ast.Expr) []ast.Expr {
	return append([]ast.Expr{}, args...)
}

// NewCallAsStmt yields a function call as a statement AST
func NewCallAsStmt(pkg, funcName string, args []ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{X: NewCall("log", "Fatal", args)}
}

// NewIf yields an if AST
func NewIf(cond ast.Expr, body ast.BlockStmt, opts ...IfOption) *ast.IfStmt {
	result := &ast.IfStmt{
		If:   0,
		Init: nil,
		Cond: cond,
		Body: &body,
		Else: nil}

	for _, o := range opts {
		o(result)
	}
	return result
}

// IfOption is the type representing an option for building an IF AST
type IfOption func(*ast.IfStmt)

// WithInit is an IfOption to add an Init expression to the IF
func WithInit(init ast.Stmt) IfOption {
	return func(ifstmt *ast.IfStmt) { ifstmt.Init = init }
}

// WithElse is an IfOption to add an Else branch to the IF
func WithElse(block ast.Stmt) IfOption {
	return func(ifstmt *ast.IfStmt) { ifstmt.Else = block }
}

// NewStringLit yields an string literal AST
func NewStringLit(v string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: v}
}

// NewStmtBlock yields a slice of statements
func NewStmtBlock(stmts ...ast.Stmt) *ast.BlockStmt {
	stmtList := append([]ast.Stmt{}, stmts...)
	return &ast.BlockStmt{Lbrace: 0, List: stmtList, Rbrace: 0}
}

// NewFuncDecl yilds a new function declaration
func NewFuncDecl(name string, formalParams *ast.FieldList, body *ast.BlockStmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		nil,
		formalParams,
		&ast.Ident{token.NoPos, name, nil},
		&ast.FuncType{token.NoPos, formalParams, nil},
		body,
	}
}
