//go:generate dbc4go -i $GOFILE -o $GOFILE
package generator

import (
	"go/ast"
	"go/parser"
	"strings"

	"github.com/chavacava/dbc4go/internal/astutils"
	cast "github.com/chavacava/dbc4go/internal/contract/parser/ast"
	"github.com/pkg/errors"
)

// GenerateCode yields the list of GO statements that enforce the given contract
// It also yields the list of errors that occurred while the generation
//@requires contract != nil
//@ensures len(contract.Ensures())+len(contract.Requires())==len(stmts)+len(errs)
func GenerateCode(contract *cast.Contract) (stmts []ast.Stmt, errs []error) {
	result := []ast.Stmt{}
	errs = []error{}
	for _, r := range contract.Requires() {
		stmt, err := generateRequiresCode(r)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to generate code for the clause '%s'", r))
			continue
		}

		result = append(result, stmt)
	}

	for _, e := range contract.Ensures() {
		stmt, err := generateEnsuresCode(e)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to generate code for the clause '%s'", e))
			continue
		}

		result = append(result, stmt)
	}

	return result, errs
}

func generateRequiresCode(r cast.Requires) (ast.Stmt, error) {
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

func generateEnsuresCode(e cast.Ensures) (ast.Stmt, error) {
	exp := e.ExpandedExpression()
	expAST, err := parser.ParseExpr("!(" + exp + ")")
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse expression '%s'", exp)
	}

	msgAST := astutils.NewStringLit("\"postcondition " + escapeDoubleQuotes(exp) + " not satisfied\"")
	panicArgs := astutils.NewCallArgs(msgAST)
	call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
	body := astutils.NewStmtBlock(call2panic)

	funcBody := astutils.NewStmtBlock(astutils.NewIf(expAST, *body))
	funcCall := astutils.NewCallAnonymous(funcBody, []ast.Expr{})
	return astutils.NewDeferStmt(funcCall), nil
}

//@ensures len(str)<=len(r)
func escapeDoubleQuotes(str string) (r string) {
	return strings.Replace(str, "\"", "\\\"", -1)
}
