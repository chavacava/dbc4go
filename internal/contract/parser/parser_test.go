package parser

import (
	"errors"
	"go/ast"
	"testing"

	"github.com/chavacava/dbc4go/internal/contract"
	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p := NewParser()
	assert.Equal(t, p, Parser{})
}

func TestParseTypeContract(t *testing.T) {
	target := "fooTest"
	tests := []struct {
		name         string
		comments     []string
		wantContract *contract.TypeContract
		wantErr      error
	}{
		{
			name:         "no comment",
			comments:     []string{},
			wantContract: contract.NewTypeContract(target),
			wantErr:      nil,
		},
		{
			name: "comment without contract",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Please use constructor",
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      nil,
		},
		{
			name: "comment with raw full invariant",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// @invariant an invariant: a == a",
			},
			wantContract: contract.NewTypeContract(target).AddEnsures(
				contract.NewEnsures("a == a", "an invariant"),
			).AddRequires(
				contract.NewRequires("a == a", "an invariant"),
			),
			wantErr: nil,
		},
		{
			name: "comment with raw invariant",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// @invariant a == a",
			},
			wantContract: contract.NewTypeContract(target).AddEnsures(
				contract.NewEnsures("a == a", ""),
			).AddRequires(
				contract.NewRequires("a == a", ""),
			),
			wantErr: nil,
		},
		{
			name: "comment with standard full invariant",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - invariant an invariant: a == a",
			},
			wantContract: contract.NewTypeContract(target).AddEnsures(
				contract.NewEnsures("a == a", "an invariant"),
			).AddRequires(
				contract.NewRequires("a == a", "an invariant"),
			),
			wantErr: nil,
		},
		{
			name: "comment with standard invariant",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - invariant a == a",
			},
			wantContract: contract.NewTypeContract(target).AddEnsures(
				contract.NewEnsures("a == a", ""),
			).AddRequires(
				contract.NewRequires("a == a", ""),
			),
			wantErr: nil,
		},
		{
			name: "comment with @old in standard invariant",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				`//   - invariant a == @old{a}`,
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      errors.New("@old can not be used in 'invariant' expressions: a == @old{a}"),
		},
		{
			name: "comment with in standard invariant + import",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - import math",
				"//   - import string",
				"//   - invariant a == a",
			},
			wantContract: contract.NewTypeContract(target).AddEnsures(
				contract.NewEnsures("a == a", ""),
			).AddRequires(
				contract.NewRequires("a == a", ""),
			).AddImport("string").AddImport("math"),
			wantErr: nil,
		},
		{
			name: "comment with invalid @ensure in type contract",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - ensures a == true",
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      errors.New("'ensures' can not be used in type contracts: ensures a == true"),
		},
		{
			name: "comment with invalid @requires in type contract",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - requires a == true",
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      errors.New("'requires' can not be used in type contracts: requires a == true"),
		},
		{
			name: "comment with invalid @unmodified in type contract",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - unmodified a",
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      errors.New("'unmodified' can not be used in type contracts: unmodified a"),
		},
		{
			name: "comment with unknown clause 'unknown' in type contract",
			comments: []string{
				"// TypeFoo is a model for foo",
				"// Contract:",
				"//   - unknown a",
			},
			wantContract: contract.NewTypeContract(target),
			wantErr:      errors.New("unknown contract kind unknown"),
		},
	}

	for _, tc := range tests {
		p := NewParser()
		got, gotErr := p.ParseTypeContract(target, commentASTFromStrings(tc.comments))

		switch {
		case tc.wantErr != nil && gotErr != nil:
			if tc.wantErr.Error() != gotErr.Error() {
				t.Fatalf("test %q: expected error %q but got %q", tc.name, tc.wantErr, gotErr)
			}
			continue
		case tc.wantErr != nil && gotErr == nil:
			t.Fatalf("test %q: expected error %q got nil", tc.name, tc.wantErr)
			continue
		case tc.wantErr == nil && gotErr != nil:
			t.Fatalf("test %q: expected no error ut got %q", tc.name, gotErr)
			continue
		}

		assert.Equal(t, tc.wantContract, got, "test %q: contract was not the expexted", tc.name)
	}
}

func TestParseFuncContract(t *testing.T) {
	tests := []struct {
		name         string
		comments     []string
		wantContract *contract.FuncContract
		wantErr      error
	}{
		{
			name:         "no comment",
			comments:     []string{},
			wantContract: contract.NewFuncContract(),
			wantErr:      nil,
		},
		{
			name: "comment without contract",
			comments: []string{
				"// Do does",
				"// what a doer",
			},
			wantContract: contract.NewFuncContract(),
			wantErr:      nil,
		},
		{
			name: "comment with raw full invariant",
			comments: []string{
				"// Do does",
				"// @invariant an invariant: a == a",
			},
			wantContract: contract.NewFuncContract(),
			wantErr:      errors.New("can not define invariants for functions/methods"),
		},
		{
			name: "comment with raw let",
			comments: []string{
				"// @let dummy: a := false",
			},
			wantContract: contract.NewFuncContract().AddLet(contract.NewLet("a := false", "dummy")),
			wantErr:      nil,
		},
		{
			name: "comment with raw let with invalid @old",
			comments: []string{
				"// @let dummy: @old{a} := false",
			},
			wantContract: contract.NewFuncContract(),
			wantErr:      errors.New("@old can not be used in 'let' expressions: @old{a} := false"),
		},
		{
			name: "comment with standard requires",
			comments: []string{
				"// Contract:",
				"//",
				"//   - requires a == a",
			},
			wantContract: contract.NewFuncContract().AddRequires(
				contract.NewRequires("a == a", "")),
			wantErr: nil,
		},
		{
			name: "comment with standard requires but using an invalid @old{}",
			comments: []string{
				"// Contract:",
				"//",
				"//   - requires a == @old{a}",
			},
			wantContract: contract.NewFuncContract(),
			wantErr:      errors.New("@old can not be used in 'requires' expressions: a == @old{a}"),
		},
		{
			name: "comment with standard ensures",
			comments: []string{
				"// Contract:",
				"//",
				"//   - ensures a == @old{a}",
				"//   - ensures with description: b == @old{b}",
			},
			wantContract: contract.NewFuncContract().AddEnsures(
				contract.NewEnsures("a == @old{a}", "")).AddEnsures(contract.NewEnsures("b == @old{b}", "with description")),
			wantErr: nil,
		},
		{
			name: "comment with raw ensures",
			comments: []string{
				"// @ensures a == @old{a}",
				"// @ensures with description: b == @old{b}",
			},
			wantContract: contract.NewFuncContract().AddEnsures(
				contract.NewEnsures("a == @old{a}", "")).AddEnsures(contract.NewEnsures("b == @old{b}", "with description")),
			wantErr: nil,
		},
		{
			name: "comment with raw unmodified",
			comments: []string{
				"// @unmodified a, b",
			},
			wantContract: contract.NewFuncContract().AddEnsures(
				contract.NewEnsures("@old{a} == a", "a unmodified")).AddEnsures(contract.NewEnsures("@old{b} == b", "b unmodified")),
			wantErr: nil,
		},
		{
			name: "comment with raw imports",
			comments: []string{
				"// @import a",
				"// @import b",
			},
			wantContract: contract.NewFuncContract().AddImport("a").AddImport("b"),
			wantErr:      nil,
		},
		{
			name: "comment with standard imports",
			comments: []string{
				"//",
				"// Contract:",
				"//    - import a",
				"//    - import b",
			},
			wantContract: contract.NewFuncContract().AddImport("a").AddImport("b"),
			wantErr:      nil,
		},
		{
			name: "comment with unknown clause",
			comments: []string{
				"//",
				"// Contract:",
				"//    - import a",
				"//    - imports b",
			},
			wantContract: contract.NewFuncContract(),
			wantErr:      errors.New("unknown contract kind imports"),
		},
		{
			name: "all in one standard",
			comments: []string{
				"//",
				"// Contract:",
				"//    - import a",
				"//    - let b := true",
				"//    - requires dummy req: c != true",
				"//    - ensures dummy ensures: @old(d) == d ==> false",
				"//    - unmodified a",
			},
			wantContract: contract.NewFuncContract().AddImport("a").
				AddLet(contract.NewLet("b := true", "")).
				AddRequires(contract.NewRequires("c != true", "dummy req")).
				AddEnsures(contract.NewEnsures("@old(d) == d ==> false", "dummy ensures")).
				AddEnsures(contract.NewEnsures("@old{a} == a", "a unmodified")),
			wantErr: nil,
		},
		{
			name: "all in one raw",
			comments: []string{
				"//",
				"// @import a",
				"// @let b := true",
				"// @requires dummy req: c != true",
				"// @ensures dummy ensures: @old(d) == d ==> false",
				"// @unmodified a",
			},
			wantContract: contract.NewFuncContract().AddImport("a").
				AddLet(contract.NewLet("b := true", "")).
				AddRequires(contract.NewRequires("c != true", "dummy req")).
				AddEnsures(contract.NewEnsures("@old(d) == d ==> false", "dummy ensures")).
				AddEnsures(contract.NewEnsures("@old{a} == a", "a unmodified")),
			wantErr: nil,
		},
		{
			name: "raw with multilines",
			comments: []string{
				"//",
				"// @import /",
				"//  a",
				"// @let b := /",
				"//   true",
				"// @requires/",
				"//   dummy req: /",
				"//    c != true",
				"// @ensures /",
				"//    dummy ensures: /",
				"//    @old(d) == d /",
				"//        ==> false",
				"// @unmodified /",
				"// a",
			},
			wantContract: contract.NewFuncContract().AddImport("a").
				AddLet(contract.NewLet("b :=  true", "")).
				AddRequires(contract.NewRequires("c != true", "dummy req")).
				AddEnsures(contract.NewEnsures("@old(d) == d  ==> false", "dummy ensures")).
				AddEnsures(contract.NewEnsures("@old{a} == a", "a unmodified")),
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		p := NewParser()
		got, gotErr := p.ParseFuncContract(commentASTFromStrings(tc.comments))

		switch {
		case tc.wantErr != nil && gotErr != nil:
			if tc.wantErr.Error() != gotErr.Error() {
				t.Fatalf("test %q: expected error %q but got %q", tc.name, tc.wantErr, gotErr)
			}
			continue
		case tc.wantErr != nil && gotErr == nil:
			t.Fatalf("test %q: expected error %q got nil", tc.name, tc.wantErr)
			continue
		case tc.wantErr == nil && gotErr != nil:
			t.Fatalf("test %q: expected no error ut got %q", tc.name, gotErr)
			continue
		}

		assert.Equal(t, tc.wantContract, got, "test %q: contract was not the expected", tc.name)
	}
}

func commentASTFromStrings(lines []string) []*ast.Comment {
	result := []*ast.Comment{}
	for _, line := range lines {
		result = append(result, &ast.Comment{Text: line})
	}
	return result
}
