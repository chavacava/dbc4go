package parser

import (
	"testing"

	"github.com/chavacava/dbc4go/internal/contract"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p := NewParser()
	assert.Equal(t, p, Parser{})
}

func TestParseImport(t *testing.T) {
	tests := []struct {
		path     string
		expected string
		err      bool
	}{
		{
			path:     "path",
			expected: "path",
			err:      false,
		},
	}

	p := NewParser()
	for _, tc := range tests {
		path, err := p.parseImport(tc.path)
		assert.Equal(t, path, tc.expected)
		if tc.err {
			assert.NotEqual(t, err, nil)
		}
	}
}

func TestParseRequires(t *testing.T) {
	tests := []struct {
		expr     string
		expected contract.Requires
		err      bool
	}{
		{
			expr:     "a == b",
			expected: contract.NewRequires("a == b"),
			err:      false,
		},
	}

	p := NewParser()
	for _, tc := range tests {
		r, err := p.parseRequires(tc.expr)
		assert.Equal(t, r, tc.expected)
		if tc.err {
			assert.NotEqual(t, err, nil)
		}
	}
}

func TestParseEnsures(t *testing.T) {
	tests := []struct {
		expr     string
		expected contract.Ensures
		err      bool
	}{
		{
			expr:     "a == b",
			expected: contract.NewEnsures("a == b"),
			err:      false,
		},
	}

	p := NewParser()
	for _, tc := range tests {
		r, err := p.parseEnsures(tc.expr)
		assert.Equal(t, r, tc.expected)
		if tc.err {
			assert.NotEqual(t, err, nil)
		}
	}
}

func TestParse(t *testing.T) {
	c := contract.NewFuncContract(nil)
	tests := []struct {
		line     string
		contract contract.FuncContract
		err      bool
	}{
		{
			line:     "//@ensures",
			contract: c,
			err:      false,
		},
		{
			line:     "//@ensures a == b",
			contract: c,
			err:      false,
		},
		{
			line:     "//@requires",
			contract: c,
			err:      false,
		},
		{
			line:     "//@requires (a > b)",
			contract: c,
			err:      false,
		},
		{
			line:     "//@import",
			contract: c,
			err:      false,
		},
		{
			line:     "//@import strings",
			contract: c,
			err:      false,
		},
		{
			line:     "//@unknown a",
			contract: c,
			err:      true,
		},
	}
	p := NewParser()
	for _, tc := range tests {
		err := p.Parse(&tc.contract, tc.line)
		if tc.err {
			assert.NotEqual(t, err, nil, "line %s", tc.line)
		}
	}
}
