// Package parser implements contract parsing
package parser

//go:generate dbc4go -i $GOFILE -o $GOFILE

import (
	"fmt"
	"regexp"

	"github.com/chavacava/dbc4go/internal/contract"
	"github.com/pkg/errors"
)

// Parser parses function contracts
type Parser struct {
}

// NewParser create a new contract parser
func NewParser() Parser {
	return Parser{}
}

var reContracts = regexp.MustCompile(`\s*//\s*@(?P<kind>[a-z]+)(?:[\t ]+(?P<description>\[[\w\s\d,]+\]))?[\t ]+(?P<expr>[^$]+)`)

func parseLine(line string) (kind, description, expr string, matched bool) {
	r2 := reContracts.FindAllStringSubmatch(line, -1)
	if r2 == nil {
		return kind, description, expr, false
	}

	kind = r2[0][1]
	expr = r2[0][2]
	description = ""
	if len(r2[0]) == 4 {
		description = expr
		expr = r2[0][3]
	}

	fmt.Printf(">>>> kind:%q\tdescription:%q\texpr:%q\n", kind, description, expr)

	return kind, description, expr, true
}

// Parse enrich the Contract with the clause if present in the given comment line
// @requires contract != nil
func (p Parser) Parse(contract *contract.FuncContract, line string) error {
	kind, description, expr, matched := parseLine(line)
	if !matched {
		return nil // nothing to do, there is no contract in this comment line
	}

	switch kind {
	case "requires":
		clause, err := p.parseRequires(expr, description)
		if err != nil {
			return errors.Wrap(err, "invalid @requires clause")
		}

		contract.AddRequires(clause)
	case "ensures":
		clause, err := p.parseEnsures(expr, description)
		if err != nil {
			return errors.Wrap(err, "invalid @ensures clause")
		}

		contract.AddEnsures(clause)
	case "import":
		clause, err := p.parseImport(expr)
		if err != nil {
			return errors.Wrap(err, "invalid @import clause")
		}

		contract.AddImport(clause)
	default:
		return errors.Errorf("unknown contract kind %s", kind)
	}

	return nil
}

// @requires path != ""
// @ensures r == "" ==> err != nil
func (p Parser) parseImport(path string) (r string, err error) {
	return path, nil
}

// @requires expr != ""
// @ensures r == contract.Requires{} ==> err != nil
func (Parser) parseRequires(expr, description string) (r contract.Requires, err error) {
	return contract.NewRequires(expr, description), nil
}

// @requires expr != ""
// @ensures r == contract.Ensures{} ==> err != nil
func (Parser) parseEnsures(expr, description string) (r contract.Ensures, err error) {
	return contract.NewEnsures(expr, description), nil
}
