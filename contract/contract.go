package contract

import (
	"fmt"
	"github.com/chavacava/dbc4go/astutils"
	"go/ast"
	"go/parser"
	"regexp"
)

// Parser provides access to functions for parsing contracts
type Parser struct{}

// Parse a text line and produces the GO code to enforce the contract in the line
// If the line has not a contract it returns nil, nil
func (p Parser) Parse(line string) (ast.Stmt, error) {
	re := regexp.MustCompile("@(?P<kind>[a-z]+)[\t ]+(?P<prop>[^$]+)")
	r2 := re.FindAllStringSubmatch(line, -1)

	if r2 == nil {
		return nil, nil // nothing to do, there is no contract in this comment line
	}

	kind := r2[0][1]
	expr := r2[0][2]

	switch kind {
	case "requires":
		return p.parseRequires(expr)
	default:
		return nil, fmt.Errorf("unknown contract kind %s", kind)
	}
}

func (p Parser) parseRequires(prop string) (ast.Stmt, error) {
	expAST, err := parser.ParseExpr("!(" + prop + ")")
	if err != nil {
		return nil, fmt.Errorf("unable to parse @require proposition '%s': %s", prop, err.Error())
	}

	return p.generateRequires(expAST, prop), nil
}

func (p Parser) generateRequires(cond ast.Expr, prop string) ast.Stmt {
	msgAST := astutils.NewStringLit("\"precondition " + prop + " not satisfied\"")
	panicArgs := astutils.NewCallArgs(msgAST)
	call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
	body := astutils.NewStmtBlock(call2panic)

	return astutils.NewIf(cond, *body)
}
