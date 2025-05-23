package contract

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	want := &FuncContract{
		requires: []Requires{},
		ensures:  []Ensures{},
		lets:     []Let{},
		imports:  map[string]struct{}{},
	}

	got := NewFuncContract()
	require.NotNil(t, got)
	assert.Equal(t, got, want)
}

func TestAddRequires(t *testing.T) {
	sampleReq1 := NewRequires(Expression{Raw: "1"}, "")
	sampleReq2 := NewRequires(Expression{Raw: "2"}, "")
	tests := []struct {
		requires Requires
		want     []Requires
	}{
		{
			requires: sampleReq1,
			want:     []Requires{sampleReq1},
		},
		{
			requires: sampleReq2,
			want:     []Requires{sampleReq1, sampleReq2},
		},
		{
			requires: sampleReq1,
			want:     []Requires{sampleReq1, sampleReq2, sampleReq1},
		},
	}

	c := NewFuncContract()
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddRequires(tc.requires)
		assert.Equal(t, c.requires, tc.want)
	}
}

func TestAddImports(t *testing.T) {
	tests := []struct {
		newPath string
		want    []Requires
	}{
		{
			newPath: "/dir",
		},
		{
			newPath: "/dir2/",
		},
		{
			newPath: "rootdir/dir1/dir2",
		},
	}

	c := NewFuncContract()
	require.NotNil(t, c)

	for i, tc := range tests {
		c.AddImport(tc.newPath)

		if len(c.Imports()) != i+1 {
			t.Fatalf("Expected len of imports %d but got %d (%+v)", i+1, len(c.imports), c.Imports())
		}

		normalizedPath := strings.Trim(tc.newPath, "\"")
		_, ok := c.imports[normalizedPath]
		if !ok {
			t.Fatalf("Expected %s to be a key in the map but not %+v", normalizedPath, c.imports)
		}
	}
}

func TestAddLets(t *testing.T) {
	sampleLet1 := NewLet(Expression{Raw: "expr1"}, "descr1")
	sampleLet2 := NewLet(Expression{Raw: "expr2"}, "descr2")
	tests := []struct {
		let  Let
		want []Let
	}{
		{
			let:  sampleLet1,
			want: []Let{sampleLet1},
		},
		{
			let:  sampleLet2,
			want: []Let{sampleLet1, sampleLet2},
		},
		{
			let:  sampleLet1,
			want: []Let{sampleLet1, sampleLet2, sampleLet1},
		},
	}

	c := NewFuncContract()
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddLet(tc.let)
		assert.Equal(t, c.lets, tc.want)
		assert.Equal(t, c.Lets(), tc.want)
	}
}

func TestLetGetters(t *testing.T) {
	expr := "expr"
	descr := "descr"
	let := NewLet(Expression{Raw: expr}, descr)

	assert.Equal(t, expr, let.Expression().Raw)
	assert.Equal(t, expr, let.ExpandedExpression().Raw)
	assert.Equal(t, descr, let.Description())
}

func TestRequires(t *testing.T) {
	sampleReq1 := NewRequires(Expression{Raw: "1"}, "")
	sampleReq2 := NewRequires(Expression{Raw: "2"}, "")
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
	sampleEns1 := NewEnsures(Expression{Raw: "1"}, "")
	sampleEns2 := NewEnsures(Expression{Raw: "2"}, "")
	tests := []struct {
		ensures Ensures
		want    []Ensures
	}{
		{
			ensures: sampleEns1,
			want:    []Ensures{sampleEns1},
		},
		{
			ensures: sampleEns2,
			want:    []Ensures{sampleEns1, sampleEns2},
		},
		{
			ensures: sampleEns1,
			want:    []Ensures{sampleEns1, sampleEns2, sampleEns1},
		},
	}

	c := NewFuncContract()
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddEnsures(tc.ensures)
		assert.Equal(t, c.ensures, tc.want)
	}
}

func TestEnsures(t *testing.T) {
	sampleEns1 := NewEnsures(Expression{Raw: "1"}, "")
	sampleEns2 := NewEnsures(Expression{Raw: "2"}, "")
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
			expandedExpr: "p ==> q",
		},
		{
			expr:         "a == b ==> b == c",
			expandedExpr: "a == b ==> b == c",
		},
	}

	for _, tc := range tests {
		r := Requires{expr: Expression{Raw: tc.expr}}
		assert.Equal(t, tc.expandedExpr, r.ExpandedExpression().Raw)
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
			expandedExpr: "p ==> q",
		},
		{
			expr:         "a == b ==> b == c",
			expandedExpr: "a == b ==> b == c",
		},
		{
			expr:         "a == b ==> b == @old{c}",
			expandedExpr: "a == b ==> b == old_1",
		},
	}

	for _, tc := range tests {
		e := Ensures{expr: Expression{Raw: tc.expr}}
		_, got, _ := e.ExpandedExpression()
		assert.Equal(t, tc.expandedExpr, got)
	}
}
