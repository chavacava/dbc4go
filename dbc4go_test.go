package main

import (
	"fmt"
	"testing"
)

func TestGetContractType(t *testing.T) {
	var tt = []struct {
		contract string
		expected int
	}{
		{"@\\requires nonil", requiresNonil},
		{"@\\requires bla bla", requiresExp},
		{"@\\ensures ", -1},
	}

	for _, test := range tt {
		if r, _ := getContractType(test.contract); r != test.expected {
			t.Errorf("for %s expexted %d but got %d", test.contract, test.expected, r)
		}
	}
}

func TestGetContractComments(t *testing.T) {
	var tt = []struct {
		comment  string
		expected int
	}{
		{"@\\requires nonil \n\n", 1},
		{"#@\\requires bla bla", 0},
		{"\\requires bla bla", 0},
		{"@\\requires bla bla\ninterline\n@\\requires nonil", 2},
	}

	for _, test := range tt {
		if r := getContractComments(test.comment); len(r) != test.expected {
			t.Errorf("for %s expexted %d but got %d", test.comment, test.expected, len(r))
		}
	}
}

func TestAnalyzeCodeBadFilename(t *testing.T) {
	var tt = []struct {
		file     string
		err      error
		buffsize int
	}{
		{"not-found", fmt.Errorf(""), 0},
		{"./test-fixtures/example1.go", nil, 0},
	}
	for _, test := range tt {
		b, err := analyzeCode([]byte{}, test.file)
		if test.err != nil && err == nil {
			t.Error("expected error, got no error")
		}
		if test.err == nil && err != nil {
			t.Errorf("no error expected, got: %v", err)
		}
		if test.buffsize != b.Len() {
			t.Errorf("expected result buffer of size %d bytes, got %d bytes", test.buffsize, b.Len())
		}
	}
}
