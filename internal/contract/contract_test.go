package contract

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	fd := &ast.FuncDecl{}
	tests := []struct {
		target *ast.FuncDecl
		result *FuncContract
	}{
		{
			target: fd,
			result: &FuncContract{requires: []Requires{}, ensures: []Ensures{}, target: fd, imports: map[string]struct{}{}},
		},
	}

	for _, tc := range tests {
		c := NewFuncContract(tc.target)
		require.NotNil(t, c)
		assert.Equal(t, c, tc.result)
	}
}

func TestAddRequires(t *testing.T) {
	sampleReq1 := NewRequires("1", "")
	sampleReq2 := NewRequires("2", "")
	tests := []struct {
		requires         Requires
		expectedRequires []Requires
	}{
		{
			requires:         sampleReq1,
			expectedRequires: []Requires{sampleReq1},
		},
		{
			requires:         sampleReq2,
			expectedRequires: []Requires{sampleReq1, sampleReq2},
		},
		{
			requires:         sampleReq1,
			expectedRequires: []Requires{sampleReq1, sampleReq2, sampleReq1},
		},
	}

	c := NewFuncContract(&ast.FuncDecl{})
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddRequires(tc.requires)
		assert.Equal(t, c.requires, tc.expectedRequires)
	}
}

func TestTarget(t *testing.T) {
	fd := &ast.FuncDecl{}
	c := FuncContract{target: fd}

	tgt := c.Target()
	assert.Equal(t, tgt, fd)
}

func TestRequires(t *testing.T) {
	sampleReq1 := NewRequires("1", "")
	sampleReq2 := NewRequires("2", "")
	tests := []struct {
		requires []Requires
	}{
		{
			requires: []Requires{},
		},
		{
			requires: []Requires{sampleReq1},
		},
		{
			requires: []Requires{sampleReq1, sampleReq2},
		},
		{
			requires: []Requires{sampleReq1, sampleReq2, sampleReq1},
		},
	}

	for _, tc := range tests {
		c := FuncContract{requires: tc.requires}
		assert.Equal(t, c.Requires(), tc.requires)
	}
}

func TestAddEnsures(t *testing.T) {
	sampleEns1 := NewEnsures("1", "")
	sampleEns2 := NewEnsures("2", "")
	tests := []struct {
		ensures         Ensures
		expectedEnsures []Ensures
	}{
		{
			ensures:         sampleEns1,
			expectedEnsures: []Ensures{sampleEns1},
		},
		{
			ensures:         sampleEns2,
			expectedEnsures: []Ensures{sampleEns1, sampleEns2},
		},
		{
			ensures:         sampleEns1,
			expectedEnsures: []Ensures{sampleEns1, sampleEns2, sampleEns1},
		},
	}

	c := NewFuncContract(&ast.FuncDecl{})
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddEnsures(tc.ensures)
		assert.Equal(t, c.ensures, tc.expectedEnsures)
	}
}

func TestEnsures(t *testing.T) {
	sampleEns1 := NewEnsures("1", "")
	sampleEns2 := NewEnsures("2", "")
	tests := []struct {
		ensures []Ensures
	}{
		{
			ensures: []Ensures{},
		},
		{
			ensures: []Ensures{sampleEns1},
		},
		{
			ensures: []Ensures{sampleEns1, sampleEns2},
		},
		{
			ensures: []Ensures{sampleEns1, sampleEns2, sampleEns1},
		},
	}

	for _, tc := range tests {
		c := FuncContract{ensures: tc.ensures}
		assert.Equal(t, c.Ensures(), tc.ensures)
	}
}

func TestRequiresExpandedExpr(t *testing.T) {
	tests := []struct {
		expr         string
		expandedExpr string
	}{
		{
			expr:         "",
			expandedExpr: "",
		},
		{
			expr:         "true",
			expandedExpr: "true",
		},
		{
			expr:         "a == b.c",
			expandedExpr: "a == b.c",
		},
		{
			expr:         "a != b.c",
			expandedExpr: "a != b.c",
		},
		{
			expr:         "p ==> q",
			expandedExpr: "!(p) || (q)",
		},
		{
			expr:         "a == b ==> b == c",
			expandedExpr: "!(a == b) || (b == c)",
		},
	}

	for _, tc := range tests {
		r := Requires{expr: tc.expr}
		assert.Equal(t, tc.expandedExpr, r.ExpandedExpression())
	}
}

func TestRequiresString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{
			expr: "",
		},
		{
			expr: "true",
		},
		{
			expr: "a == b.c",
		},
		{
			expr: "a != b.c",
		},
		{
			expr: "p ==> q",
		},
		{
			expr: "a == b ==> b == c",
		},
	}

	for _, tc := range tests {
		r := Requires{expr: tc.expr}
		assert.Equal(t, "@requires "+tc.expr, r.String())
	}
}

func TestEnsuresExpandedExpr(t *testing.T) {
	tests := []struct {
		expr         string
		expandedExpr string
	}{
		{
			expr:         "",
			expandedExpr: "",
		},
		{
			expr:         "true",
			expandedExpr: "true",
		},
		{
			expr:         "a == b.c",
			expandedExpr: "a == b.c",
		},
		{
			expr:         "a != b.c",
			expandedExpr: "a != b.c",
		},
		{
			expr:         "p ==> q",
			expandedExpr: "!(p) || (q)",
		},
		{
			expr:         "a == b ==> b == c",
			expandedExpr: "!(a == b) || (b == c)",
		},
		{
			expr:         "a == b ==> b == @old{c}",
			expandedExpr: "!(a == b) || (b == old_1)",
		},
	}

	for _, tc := range tests {
		e := Ensures{expr: tc.expr}
		_, got, _ := e.ExpandedExpression()
		assert.Equal(t, tc.expandedExpr, got)
	}
}

func TestEnsuresString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{
			expr: "",
		},
		{
			expr: "true",
		},
		{
			expr: "a == b.c",
		},
		{
			expr: "a != b.c",
		},
		{
			expr: "p ==> q",
		},
		{
			expr: "a == b ==> b == c",
		},
	}

	for _, tc := range tests {
		r := Ensures{expr: tc.expr}
		assert.Equal(t, "@ensures "+tc.expr, r.String())
	}
}
