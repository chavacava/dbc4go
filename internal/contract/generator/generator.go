//go:generate dbc4go -i $GOFILE -o $GOFILE
package generator

import (
	"go/ast"
	"go/parser"
	"strings"

	"github.com/chavacava/dbc4go/internal/astutils"
	"github.com/chavacava/dbc4go/internal/contract"
	"github.com/pkg/errors"
)

// GenerateCode yields the list of GO statements that enforce the given contract
// It also yields the list of errors that occurred while the generation
//@requires c != nil
// Must ensure
// 		len(ensures) > 0 => len(requires)+1 == len(stmts)+len(errs) and
// 		len(ensures) == 0 => len(requires) == len(stmts)+len(errs)
//@ensures !(len(c.Ensures()) > 0) || len(c.Requires())+1 == len(stmts)+len(errs)
//@ensures !(len(c.Ensures()) == 0) || len(c.Requires()) == len(stmts)+len(errs)
func GenerateCode(c *contract.FuncContract) (stmts []ast.Stmt, errs []error) {
	result := []ast.Stmt{}
	errs = []error{}
	for _, r := range c.Requires() {
		stmt, err := generateRequiresCode(r)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to generate code for the clause '%s'", r))
			continue
		}

		result = append(result, stmt)
	}

	if len(c.Ensures()) > 0 {
		stmt, err := generateEnsuresCode(c.Ensures(), c.Target())
		if err != nil {
			errs = append(errs, errors.Wrap(err, "unable to generate code for @ensure clause"))
		} else {
			result = append(result, stmt)
		}
	}

	return result, errs
}

func generateRequiresCode(r contract.Requires) (ast.Stmt, error) {
	exp := r.ExpandedExpression()
	expAST, err := parser.ParseExpr("!(" + exp + ")")
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse expression '%s'", exp)
	}

	msgAST := astutils.NewStringLit("\"precondition " + escapeDoubleQuotes(exp) + " not satisfied\"")
	panicArgs := astutils.NewCallArgs(msgAST)
	call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
	body := astutils.NewStmtBlock(call2panic)

	return astutils.NewIf(expAST, *body), nil
}

func generateEnsuresCode(clauses []contract.Ensures, fd *ast.FuncDecl) (ast.Stmt, error) {
	funcBody := []ast.Stmt{}

	for _, clause := range clauses {
		exp := clause.ExpandedExpression()
		expAST, err := parser.ParseExpr("!(" + exp + ")")
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse expression '%s'", exp)
		}

		msgAST := astutils.NewStringLit("\"postcondition " + escapeDoubleQuotes(exp) + " not satisfied\"")
		panicArgs := astutils.NewCallArgs(msgAST)
		call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
		body := astutils.NewStmtBlock(call2panic)
		funcBody = append(funcBody, astutils.NewIf(expAST, *body))
	}

	funcCall := astutils.NewCallAnonymous(astutils.CopyFields(fd.Type.Params.List, "old_"), astutils.NewStmtBlock(funcBody...), astutils.ArgsFromFields(fd.Type.Params.List))
	return astutils.NewDeferStmt(funcCall), nil
}

//@ensures len(str)<=len(r)
func escapeDoubleQuotes(str string) (r string) {
	return strings.Replace(str, "\"", "\\\"", -1)
}
