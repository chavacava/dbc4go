package test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/chavacava/dbc4go/internal/contract/generator"
	"github.com/stretchr/testify/assert"
)

func TestDBC4GO(t *testing.T) {
	tests := []struct {
		input     string
		expOutput string // expected output
	}{
		{
			input:     "./fixtures/1.input",
			expOutput: "./fixtures/1.output",
		},
		{
			input:     "./fixtures/2.input",
			expOutput: "./fixtures/2.output",
		},
		{
			input:     "./fixtures/3.input",
			expOutput: "./fixtures/3.output",
		},
	}

	tmpDir, err := ioutil.TempDir("", "testing-dbc4go")
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

		err = compareFiles(test.expOutput, outputFilename, t)
		if err != nil {
			t.Fatalf("Unable to compare generated file with reference: %v", err)
		}
		input.Close()
	}
}

func compareFiles(f1, f2 string, t *testing.T) error {
	content1, err := ioutil.ReadFile(f1)
	if err != nil {
		return err
	}

	content2, err := ioutil.ReadFile(f2)
	if err != nil {
		return err
	}

	assert.Equal(t, content1, content2)

	return nil
}
