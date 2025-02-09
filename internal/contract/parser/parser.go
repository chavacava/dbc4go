// Package parser implements contract parsing
package parser

//go:generate dbc4go -i $GOFILE -o $GOFILE

import (
	"fmt"
	"go/ast"
	"iter"
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

func (p *Parser) ParseTypeContract(typeName string, comments []*ast.Comment) (result *contract.TypeContract, err error) {
	result = contract.NewTypeContract(typeName)
	for canonicalLine := range p.canonicalLinesFromComments(comments) {
		kind, description, exprStr, matched := p.parseLine(canonicalLine)
		if !matched {
			continue // nothing to do, there is no contract in this comment line
		}

		expr := parseExpression(exprStr)

		switch kind {
		case "invariant":
			if contract.Re4old.MatchString(exprStr) {
				return result, fmt.Errorf("@old can not be used in 'invariant' expressions: %s", exprStr)
			}

			ensuresClause := contract.NewEnsures(expr, description) // invariants are ensures that apply to all methods of the type

			result.AddEnsures(ensuresClause)

			requiresClause := contract.NewRequires(expr, description) // invariants are, also, requires that apply to all methods of the type

			result.AddRequires(requiresClause)
		case "import":
			result.AddImport(expr.Raw)
		case "ensures", "requires", "unmodified":
			return result, fmt.Errorf("'%s' can not be used in type contracts: %s %s", kind, kind, expr.Raw)
		default:
			return result, errors.Errorf("unknown contract kind %s", kind)
		}
	}

	return result, nil
}

func (p *Parser) ParseFuncContract(comments []*ast.Comment) (result *contract.FuncContract, err error) {
	result = contract.NewFuncContract()
	for canonicalLine := range p.canonicalLinesFromComments(comments) {
		kind, description, exprStr, matched := p.parseLine(canonicalLine)
		if !matched {
			continue // nothing to do, there is no contract in this comment line
		}

		expr := parseExpression(exprStr)
		switch kind {
		case "requires":
			if contract.Re4old.MatchString(exprStr) {
				return result, fmt.Errorf("@old can not be used in 'requires' expressions: %s", exprStr)
			}

			clause := contract.NewRequires(expr, description)
			result.AddRequires(clause)
		case "ensures":
			clause := contract.NewEnsures(expr, description)
			result.AddEnsures(clause)
		case "import":
			result.AddImport(expr.Raw)
		case "invariant":
			return result, errors.New("can not define invariants for functions/methods")
		case "let":
			if contract.Re4old.MatchString(exprStr) {
				return result, fmt.Errorf("@old can not be used in 'let' expressions: %s", exprStr)
			}

			clause := contract.NewLet(expr, description)

			result.AddLet(clause)
		case "unmodified":
			clauses := p.parseUnmodified(expr)

			for _, clause := range clauses {
				result.AddEnsures(clause)
			}
		default:
			return result, errors.Errorf("unknown contract kind %s", kind)
		}
	}

	return result, nil
}

// Contract:
//   - ensures r != nil
func (p *Parser) parseUnmodified(expr contract.Expression) (r []contract.Ensures) {
	result := []contract.Ensures{}

	ids := strings.Split(expr.Raw, ",")
	for _, id := range ids {
		id := strings.TrimSpace(id)
		exprStr := fmt.Sprintf("@old{%s} == %s", id, id)
		expr.Raw = exprStr
		description := fmt.Sprintf("%s unmodified", id)
		newEnsure := contract.NewEnsures(expr, description)
		result = append(result, newEnsure)
	}

	return result
}

var reRawFormatContractClause = regexp.MustCompile(`\s*@(?P<kind>[a-z]+)(?:\s+(?P<description>[\w\s\d,]+): )?\s?(?P<expr>[^$]+)`)
var reDirectiveFormatContractClause = regexp.MustCompile(`\s*contract\:(?P<kind>[a-z]+)(?:\s+(?P<description>[\w\s\d,]+): )?\s?(?P<expr>[^$]+)`)
var reStandardFormatContractClause = regexp.MustCompile(`\s*-\s+(?P<kind>[a-z]+)(?:\s+(?P<description>[\w\s\d,]+): )?\s?(?P<expr>[^$]+)`)
var reStandardFormatContractBlockStarter = regexp.MustCompile(`\s*[Cc]ontract:\s*$`)

// parseLine extracts kind, description and expr from a given comment line
// If the line is a contract annotation it returns matched true, false otherwise.
//
// Contract:
//   - ensures matched ==> kind != ""
//   - ensures matched ==> expr != ""
func (p *Parser) parseLine(line string) (kind, description, expr string, matched bool) {
	r2 := reDirectiveFormatContractClause.FindAllStringSubmatch(line, -1)
	if r2 != nil {
		return extractContractPartsFromMatch(r2)
	}

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

	r2 = reRawFormatContractClause.FindAllStringSubmatch(line, -1)
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
	return kind, description, strings.TrimSpace(expr), true
}

func (*Parser) canonicalLinesFromComments(comments []*ast.Comment) iter.Seq[string] {
	return func(yield func(string) bool) {
		acc := ""
		mustAcc := false
		for _, commentLine := range comments {
			line := strings.TrimLeft(commentLine.Text, "/")
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			if strings.HasSuffix(line, "/") {
				line = strings.TrimRight(line, "/") + " "
				mustAcc = true
			}

			acc += line

			if mustAcc {
				mustAcc = false
				continue
			}

			if !yield(acc) {
				return
			}
			acc = ""
		}
	}
}

var reForeach = regexp.MustCompile(`^@forall (?P<variables>[^@]+) @(?P<kind>indexof|in) (?P<sources>[^:]+):\s+(?P<expression>.+)$`)

func parseExpression(raw string) contract.Expression {
	matches := reForeach.FindStringSubmatch(raw)
	if matches != nil {
		subExprs := make(map[string]contract.Expression)
		for i, name := range reForeach.SubexpNames() {
			if i != 0 && name != "" {
				subExprs[name] = parseExpression(matches[i])
			}
		}
		return contract.Expression{
			Kind:     contract.ExprKindForeach,
			SubExprs: subExprs,
			Raw:      raw,
		}
	}

	return contract.Expression{
		Kind: contract.ExprKindPlain,
		Raw:  raw,
	}
}
