//go:generate dbc4go -i $GOFILE -o $GOFILE
package main

import (
	"flag"
	"log"

	"github.com/chavacava/dbc4go/internal/contract/generator"
)

func main() {
	inputFilename := flag.String("i", "", "input source file")
	outputFilename := flag.String("o", "", "output file (defaults to stdout")
	flag.Parse()

	if *inputFilename == "" {
		log.Fatal("Undefined input file, please set the flag -i")
	}

	err := generator.GenerateCode(*inputFilename, *outputFilename)
	if err != nil {
		log.Fatalf("Unable to generate code: %v", err)
	}
}
