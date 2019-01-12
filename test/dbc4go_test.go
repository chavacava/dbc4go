package test

import (
	"io/ioutil"
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
	}

	tmpDir, err := ioutil.TempDir("/tmp", "testing-dbc4go")
	if err != nil {
		t.Fatalf("Unable to create temporary directory for tests: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	for i, test := range tests {
		outputFilename := filepath.Join(tmpDir, strconv.Itoa(i)+".go")
		err := generator.GenerateCode(test.input, outputFilename)
		if err != nil {
			t.Fatalf("Unable to generate code for test #%d: %v", i, err)
		}

		err = compareFiles(test.expOutput, outputFilename, t)
		if err != nil {
			t.Fatalf("Unable to compare generated file with reference: %v", err)
		}
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
