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
	currentContractStyle contractStyleType
}

type contractStyleType byte

const (
	contractStyleUnknown  contractStyleType = 0
	contractStyleRaw      contractStyleType = 1
	contractStyleStandard contractStyleType = 2
)

// NewParser create a new contract parser
func NewParser() Parser {
	return Parser{
		currentContractStyle: contractStyleUnknown,
	}
}

// ParseTypeContract enrich the contract with the clause if present in the given comment line.
//
// Contract:
//   - requires typeContract != nil
func (p *Parser) ParseTypeContract(typeContract *contract.TypeContract, line string) error {
	kind, description, expr, matched := p.parseLine(line)
	if !matched {
		return nil // nothing to do, there is no contract in this comment line
	}

	switch kind {
	case "invariant":
		if contract.Re4old.MatchString(expr) {
			return fmt.Errorf("@old can not be used in @invariant expressions: %s", expr)
		}

		ensuresClause, err := p.parseEnsures(expr, description) // invariants are ensures that apply to all methods of the type
		if err != nil {
			return fmt.Errorf("invalid @invariant clause: %w", err)
		}

		typeContract.AddEnsures(ensuresClause)

		requiresClause, err := p.parseRequires(expr, description) // invariants are, also, requires that apply to all methods of the type
		if err != nil {
			return fmt.Errorf("invalid @invariant clause: %w", err)
		}

		typeContract.AddRequires(requiresClause)
	case "import":
		clause, err := p.parseImport(expr)
		if err != nil {
			return fmt.Errorf("invalid @import clause: %w", err)
		}

		typeContract.AddImport(clause)
	case "ensures", "requires", "unmodified":
		return fmt.Errorf("@%s can not be used in type contracts: %s", kind, expr)
	default:
		return errors.Errorf("unknown contract kind %s", kind)
	}

	return nil
}

// ParseFuncContract enrich the Contract with the clause if present in the given comment line.
//
// Contract:
//   - requires funcContract != nil
func (p *Parser) ParseFuncContract(funcContract *contract.FuncContract, line string) error {
	kind, description, expr, matched := p.parseLine(line)
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
	case "let":
		if contract.Re4old.MatchString(expr) {
			return fmt.Errorf("@old can not be used in @let expressions: %s", expr)
		}

		clause, err := p.parseLet(expr, description)
		if err != nil {
			return fmt.Errorf("invalid @let clause: %w", err)
		}

		funcContract.AddLet(clause)
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

// Contract:
//   - requires expr != ""
//   - ensures err != nil ==> r.Expression() == expr
//   - ensures err != nil ==> r.Description() == description
func (p *Parser) parseLet(expr string, description string) (r contract.Let, err error) {
	return contract.NewLet(expr, description), nil
}

// Contract:
//   - requires path != ""
//   - ensures r == "" ==> err != nil
func (p *Parser) parseImport(path string) (r string, err error) {
	return path, nil
}

// Contract:
//   - requires expr != ""
//   - ensures err != nil ==> r.Expression() == expr
//   - ensures err != nil ==> r.Description() == description
func (*Parser) parseRequires(expr, description string) (r contract.Requires, err error) {
	return contract.NewRequires(expr, description), nil
}

// Contract:
//   - requires expr != ""
//   - ensures err != nil ==> r.Expression() == expr
//   - ensures err != nil ==> r.Description() == description
func (*Parser) parseEnsures(expr, description string) (r contract.Ensures, err error) {
	return contract.NewEnsures(expr, description), nil
}

// Contract:
//   - ensures r != nil
//   - ensures err == nil
func (p *Parser) parseUnmodified(expr string) (r []contract.Ensures, err error) {
	result := []contract.Ensures{}

	ids := strings.Split(expr, ",")
	for _, id := range ids {
		id := strings.TrimSpace(id)
		expr := fmt.Sprintf("@old{%s} == %s", id, id)
		description := fmt.Sprintf("%s unmodified", id)
		newEnsure := contract.NewEnsures(expr, description)
		result = append(result, newEnsure)
	}

	return result, nil
}

var reRawFormatContractClause = regexp.MustCompile(`\s*@(?P<kind>[a-z]+)(?:\s+(?P<description>[\w\s\d,]+): )?\s?(?P<expr>[^$]+)`)
var reStandardFormatContractClause = regexp.MustCompile(`\s*-\s+(?P<kind>[a-z]+)(?:\s+(?P<description>[\w\s\d,]+): )?\s?(?P<expr>[^$]+)`)
var reStandardFormatContractBlockStarter = regexp.MustCompile(`\s*Contract:\s*$`)

// parseLine extracts kind, description and expr from a given comment line
// If the line is a contract annotation it returns matched true, false otherwise.
//
// Contract:
//   - ensures matched ==> kind != ""
//   - ensures matched ==> expr != ""
func (p *Parser) parseLine(line string) (kind, description, expr string, matched bool) {
	if p.currentContractStyle == contractStyleUnknown {
		r2 := reStandardFormatContractBlockStarter.FindAllStringSubmatch(line, -1)
		if r2 != nil {
			p.currentContractStyle = contractStyleStandard
			return kind, description, expr, false
		}
	}

	if p.currentContractStyle == contractStyleStandard {
		r2 := reStandardFormatContractClause.FindAllStringSubmatch(line, -1)
		if r2 == nil {
			p.currentContractStyle = contractStyleUnknown
			return kind, description, expr, false
		}
		return extractContractPartsFromMatch(r2)
	}

	r2 := reRawFormatContractClause.FindAllStringSubmatch(line, -1)
	if r2 == nil {
		return kind, description, expr, false
	}

	p.currentContractStyle = contractStyleRaw
	return extractContractPartsFromMatch(r2)
}

func extractContractPartsFromMatch(match [][]string) (kind, description, expr string, ok bool) {
	if len(match[0]) < 3 {
		return kind, description, expr, false
	}

	kind = match[0][1]
	expr = match[0][2]
	description = ""
	if len(match[0]) == 4 {
		description = expr
		expr = match[0][3]
	}
	return kind, description, expr, true
}
