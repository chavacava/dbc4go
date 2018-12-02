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
	if pkg != "" {
		return &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: pkg, NamePos: 0, Obj: nil},
			Sel: &ast.Ident{Name: funcName, NamePos: 0, Obj: nil}},
			Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
	}

	return &ast.CallExpr{Fun: &ast.Ident{Name: funcName, NamePos: 0, Obj: nil},
		Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
}

func NewCallAnonymous(body *ast.BlockStmt, args []ast.Expr) *ast.CallExpr {
	fun := &ast.FuncLit{
		Type: &ast.FuncType{
			Func: 0,
			Params: &ast.FieldList{
				Opening: 0,
				Closing: 0,
				List:    []*ast.Field{},
			},
		},
		Body: body,
	}

	return &ast.CallExpr{Fun: fun, Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
}

// NewCallArgs yields a slice of expressions that can be used as arguments in
// a function call
func NewCallArgs(args ...ast.Expr) []ast.Expr {
	return append([]ast.Expr{}, args...)
}

// NewCallAsStmt yields a function call as a statement AST
func NewCallAsStmt(pkg, funcName string, args []ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{X: NewCall(pkg, funcName, args)}
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
		Doc:  nil,
		Recv: formalParams,
		Name: &ast.Ident{NamePos: 0, Name: name, Obj: nil},
		Type: &ast.FuncType{Func: 0, Params: formalParams, Results: nil},
		Body: body,
	}
}

// NewDeferStmt yilds a new defer statement
func NewDeferStmt(functionCall *ast.CallExpr) *ast.DeferStmt {
	return &ast.DeferStmt{
		Defer: 0,
		Call:  functionCall,
	}
}

// NewCallExpr yields a new function call expression
func NewCallExpr(function *ast.Expr, params []ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:      *function,
		Lparen:   0,
		Args:     params,
		Ellipsis: token.NoPos,
		Rparen:   0,
	}
}
