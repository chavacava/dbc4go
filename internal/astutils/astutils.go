// Package astutils provides a set of utility functions to build ASTs
package astutils

import (
	"go/ast"
	"go/token"
)

// NewID yields an identifier node
//@requires id != ""
//@ensures r != nil && r.Name == id
func NewID(id string) (r *ast.Ident) {
	return &ast.Ident{Name: id, NamePos: 0, Obj: nil}
}

// NewBinExp yields a binary expression AST
//@ensures r != nil
func NewBinExp(op token.Token, lhs, rhs ast.Expr) (r *ast.BinaryExpr) {
	return &ast.BinaryExpr{Op: op, OpPos: 0, X: lhs, Y: rhs}
}

// NewCall yields a function call AST
//@ensures r != nil
func NewCall(pkg, funcName string, args []ast.Expr) (r *ast.CallExpr) {
	if pkg != "" {
		return &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: pkg, NamePos: 0, Obj: nil},
			Sel: &ast.Ident{Name: funcName, NamePos: 0, Obj: nil}},
			Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
	}

	return &ast.CallExpr{Fun: &ast.Ident{Name: funcName, NamePos: 0, Obj: nil},
		Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
}

// NewCallAnonymous yields an anonymous call
//@ensures r != nil
func NewCallAnonymous(params []*ast.Field, body *ast.BlockStmt, args []ast.Expr) (r *ast.CallExpr) {
	fun := &ast.FuncLit{
		Type: &ast.FuncType{
			Func: 0,
			Params: &ast.FieldList{
				Opening: 0,
				Closing: 0,
				List:    params,
			},
		},
		Body: body,
	}

	return &ast.CallExpr{Fun: fun, Args: args, Ellipsis: 0, Lparen: 0, Rparen: 0}
}

// NewCallArgs yields a slice of expressions that can be used as arguments in
// a function call
//@ensures len(args) == len(r)
func NewCallArgs(args ...ast.Expr) (r []ast.Expr) {
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
//@ensures len(r.List) == len(stmts)
func NewStmtBlock(stmts ...ast.Stmt) (r *ast.BlockStmt) {
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

// NewDeferStmt yields a new defer statement
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

//@ensures len(r.Names) == len(names)
//@ensures r.Type == t
func newField(names []string, t ast.Expr) (r ast.Field) {
	r = ast.Field{}
	r.Names = []*ast.Ident{}
	for _, n := range names {
		id := NewID(n)
		r.Names = append(r.Names, id)
	}
	r.Type = t
	return r
}

// CopyFields creates a copy of the given field slice.
// New field names are prefixed by the givin prefix
//@ensures len(l) == len(result)
func CopyFields(l []*ast.Field, prefix string) (result []*ast.Field) {
	result = []*ast.Field{}
	for _, f := range l {
		names := []string{}
		for _, n := range f.Names {
			names = append(names, prefix+n.Name)
		}

		fType := f.Type
		e, ok := f.Type.(*ast.StarExpr)
		if ok {
			fType = e.X
		}

		ellip, ok := f.Type.(*ast.Ellipsis)
		if ok {
			fType = &ast.ArrayType{Elt: ellip.Elt}
		}

		nf := newField(names, fType)
		result = append(result, &nf)
	}

	return result
}

// ArgsFromFields creates a list of arguments ([]ast.Expr) from a list of fields
//@ensures len(l)<=len(result)
func ArgsFromFields(l []*ast.Field) (result []ast.Expr) {
	result = []ast.Expr{}
	for _, f := range l {
		for _, n := range f.Names {
			arg := ast.Expr(n)
			_, ok := f.Type.(*ast.StarExpr)
			if ok {
				arg = &ast.StarExpr{X: n}
			}

			result = append(result, arg)
		}
	}

	return result
}
