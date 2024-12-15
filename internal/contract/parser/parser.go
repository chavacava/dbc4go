// Package parser implements contract parsing
package parser

//go:generate dbc4go -i $GOFILE -o $GOFILE

import (
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

var reContracts = regexp.MustCompile("\\s*//\\s*@(?P<kind>[a-z]+)[\t ]+(?P<expr>[^$]+)")

// Parse enrich the Contract with the clause if present in the given comment line
// @requires contract != nil
func (p Parser) Parse(contract *contract.FuncContract, line string) error {
	r2 := reContracts.FindAllStringSubmatch(line, -1)

	if r2 == nil {
		return nil // nothing to do, there is no contract in this comment line
	}

	kind := r2[0][1]
	expr := r2[0][2]

	switch kind {
	case "requires":
		clause, err := p.parseRequires(expr)
		if err != nil {
			return errors.Wrap(err, "invalid @requires clause")
		}

		contract.AddRequires(clause)
	case "ensures":
		clause, err := p.parseEnsures(expr)
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
func (p Parser) parseRequires(expr string) (r contract.Requires, err error) {
	return contract.NewRequires(expr), nil
}

// @requires expr != ""
// @ensures r == contract.Ensures{} ==> err != nil
func (p Parser) parseEnsures(expr string) (r contract.Ensures, err error) {
	return contract.NewEnsures(expr), nil
}
