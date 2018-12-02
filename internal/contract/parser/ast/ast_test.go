package ast

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
		result Contract
	}{
		{
			target: fd,
			result: Contract{requires: []Requires{}, target: fd},
		},
	}

	for _, tc := range tests {
		c := NewContract(tc.target)
		require.NotNil(t, c)
		assert.Equal(t, c, tc.result)
	}
}

func TestAddRequires(t *testing.T) {
	sampleReq1 := NewRequires("1")
	sampleReq2 := NewRequires("2")
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

	c := NewContract(&ast.FuncDecl{})
	require.NotNil(t, c)

	for _, tc := range tests {
		c.AddRequires(tc.requires)
		assert.Equal(t, c.requires, tc.expectedRequires)
	}
}
