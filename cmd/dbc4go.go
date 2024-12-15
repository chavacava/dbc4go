// Package main provides the dbc4go command line application
package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/chavacava/dbc4go/internal/contract/generator"
)

func main() {
	iFilename := flag.String("i", "", "input source file (defaults to stdin)")
	oFilename := flag.String("o", "", "output file (defaults to stdout)")
	flag.Parse()

	inputFilename := *iFilename
	outputFilename := *oFilename

	input := os.Stdin
	if inputFilename != "" {
		var err error
		input, err = os.Open(inputFilename)
		if err != nil {
			log.Fatalf("Unable to open input file: %v", err)
		}
	}
	output := os.Stdout

	if outputFilename != "" {
		if outputFilename == inputFilename {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			extension := strconv.Itoa(r.Intn(100))
			outputFilename = outputFilename + "." + extension

			defer func() {
				err := os.Rename(outputFilename, inputFilename)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		var err error
		output, err = os.Create(outputFilename)
		if err != nil {
			log.Fatalf("Unable to create output file '%s': %v", outputFilename, err)
		}
		defer output.Close()

		log.Printf("Generating file '%s'", *oFilename)
	}

	err := generator.GenerateCode(input, output)
	input.Close()
	if err != nil {
		log.Fatalf("Unable to generate code: %v", err)
	}
}
