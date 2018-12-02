//go:generate dbc4go -i $GOFILE -o $GOFILE
package parser

import (
	"regexp"

	cast "github.com/chavacava/dbc4go/internal/contract/parser/ast"
	"github.com/pkg/errors"
)

// Parser parses function contracts
type Parser struct {
}

// NewParser create a new contract parser
func NewParser() Parser {
	return Parser{}
}

// Parse enrich the Contract with the clause if present in the given comment line
//@requires contract != nil
func (p Parser) Parse(contract *cast.Contract, line string) error {
	re := regexp.MustCompile("//@(?P<kind>[a-z]+)[\t ]+(?P<expr>[^$]+)")
	r2 := re.FindAllStringSubmatch(line, -1)

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
	default:
		return errors.Errorf("unknown contract kind %s", kind)
	}

	return nil
}

//@requires expr != ""
func (p Parser) parseRequires(expr string) (cast.Requires, error) {
	return cast.NewRequires(expr), nil
}
