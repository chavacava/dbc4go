// Package parser implements contract parsing
package parser

//go:generate dbc4go -i $GOFILE -o $GOFILE

import (
	"fmt"
	"regexp"
	"strings"

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

// Parse enrich the Contract with the clause if present in the given comment line
// @requires contract != nil
func (p Parser) Parse(funcContract *contract.FuncContract, line string) error {
	kind, description, expr, matched := parseLine(line)
	if !matched {
		return nil // nothing to do, there is no contract in this comment line
	}

	switch kind {
	case "requires":
		if contract.Re4old.MatchString(expr) {
			return fmt.Errorf("@old can not be used in @requires expressions: %s", expr)
		}

		clause, err := p.parseRequires(expr, description)
		if err != nil {
			return fmt.Errorf("invalid @requires clause: %w", err)
		}

		funcContract.AddRequires(clause)
	case "ensures":
		clause, err := p.parseEnsures(expr, description)
		if err != nil {
			return fmt.Errorf("invalid @ensures clause: %w", err)
		}

		funcContract.AddEnsures(clause)
	case "import":
		clause, err := p.parseImport(expr)
		if err != nil {
			return fmt.Errorf("invalid @import clause: %w", err)
		}

		funcContract.AddImport(clause)
	case "unmodified":
		clauses, err := p.parseUnmodified(expr)
		if err != nil {
			return fmt.Errorf("invalid @import clause: %w", err)
		}

		for _, clause := range clauses {
			funcContract.AddEnsures(clause)
		}
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

// @ensures r != nil
// @ensures err == nil
func (p Parser) parseUnmodified(expr string) (r []contract.Ensures, err error) {
	result := []contract.Ensures{}

	ids := strings.Split(expr, ",")
	for _, id := range ids {
		id := strings.TrimSpace(id)
		expr := fmt.Sprintf("@old(%s) == %s", id, id)
		description := fmt.Sprintf("[%s unmodified]", id)
		newEnsure := contract.NewEnsures(expr, description)
		result = append(result, newEnsure)
	}

	return result, nil
}

// parseLine extracts kind, description and expr from a given comment line
// If the line is a contract annotation it returns matched true, false otherwise.
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

	return kind, description, expr, true
}
