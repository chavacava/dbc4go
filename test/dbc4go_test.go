package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/chavacava/dbc4go/internal/contract/generator"
)

func TestDBC4GO(t *testing.T) {
	tests := []struct {
		input     string
		expOutput string // expected output
	}{
		{
			input:     "./testdata/1.input",
			expOutput: "./testdata/1.output",
		},
		{
			input:     "./testdata/2.input",
			expOutput: "./testdata/2.output",
		},
		{
			input:     "./testdata/3.input",
			expOutput: "./testdata/3.output",
		},
		{
			input:     "./testdata/4.input",
			expOutput: "./testdata/4.output",
		},
	}

	tmpDir, err := os.MkdirTemp("", "testing-dbc4go")
	if err != nil {
		t.Fatalf("Unable to create temporary directory for tests: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	for i, test := range tests {
		outputFilename := filepath.Join(tmpDir, strconv.Itoa(i)+".go")
		input, err := os.Open(test.input)
		if err != nil {
			t.Errorf("Unable to open input file %s: %v", test.input, err)
			continue
		}

		output, err := os.Create(outputFilename)
		if err != nil {
			log.Fatalf("Unable to create output file '%s': %v", outputFilename, err)
		}
		defer output.Close()

		err = generator.GenerateCode(input, output)
		if err != nil {
			t.Fatalf("Unable to generate code for test #%d: %v", i, err)
		}

		equal, err := areFilesEqual(test.expOutput, outputFilename)
		if err != nil {
			t.Fatalf("Error when comparing generated file with reference: %v", err)
		}

		if !equal {
			t.Fatalf("Files %q and %q are not equal", test.expOutput, outputFilename)
		}

		input.Close()
	}
}

func areFilesEqual(f1, f2 string) (bool, error) {
	content1, err := os.ReadFile(f1)
	if err != nil {
		return false, fmt.Errorf("error opening file %q: %w", f1, err)
	}

	content2, err := os.ReadFile(f2)
	if err != nil {
		return false, fmt.Errorf("error opening file %q: %w", f2, err)
	}

	if string(content1) != string(content2) {
		fmt.Printf("\n%s", string(content2))
		return false, nil
	}

	return true, nil
}
