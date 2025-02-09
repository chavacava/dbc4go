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
	tests := []string{
		"foreach_element_single",
		"import",
		"invariant",
		"let",
		"multiline",
		"old",
		"unmodified",
	}

	tmpDir, err := os.MkdirTemp("", "testing-dbc4go")
	if err != nil {
		t.Fatalf("Unable to create temporary directory for tests: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	for i, test := range tests {
		outputFilename := filepath.Join(tmpDir, strconv.Itoa(i)+".go")
		inputFilename := filepath.Join("./testdata", test+".go")
		input, err := os.Open(inputFilename)
		if err != nil {
			t.Errorf("Unable to open input file %s: %v", inputFilename, err)
			continue
		}

		output, err := os.Create(outputFilename)
		if err != nil {
			log.Fatalf("Unable to create output file '%s': %v", outputFilename, err)
		}

		err = generator.GenerateCode(input, output)
		if err != nil {
			t.Fatalf("Unable to generate code for test #%d: %v", i, err)
		}

		wantFilename := filepath.Join("./testdata", test+".want.go")
		equal, err := areFilesEqual(wantFilename, outputFilename)
		if err != nil {
			t.Fatalf("Error when comparing generated file with reference: %v", err)
		}

		if !equal {
			t.Fatalf("Files %q and %q are not equal", wantFilename, outputFilename)
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
