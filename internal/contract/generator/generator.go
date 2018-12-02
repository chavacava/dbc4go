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
func GenerateCode(contract *cast.Contract) ([]ast.Stmt, []error) {
	result := []ast.Stmt{}
	errs := []error{}
	for _, r := range contract.Requires() {
		stmt, err := generateRequiresCode(r)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to generate code for the clause '%s'", r))
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

func escapeDoubleQuotes(str string) string {
	return strings.Replace(str, "\"", "\\\"", -1)
}
