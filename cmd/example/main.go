package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Name of file containing input data")
	outputFile = flag.String("o", "", "Name of file for output")
)

func main() {
	var (
		input io.Reader = os.Stdin
		output io.Writer = os.Stdout
		err error
	)
	flag.Parse()
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression) 
	}
	if *inputFile != "" {
		input, err = os.Open(*inputFile)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			os.Exit(1)
		}
	}
	if *outputFile != "" {
		output, err = os.Open(*outputFile)
		if err != nil {
			output, err = os.Create(*outputFile)
			if err != nil {
				fmt.Println("Error creating file: ", err)
				os.Exit(1)
			}
		}
	}
	handler := &lab2.ComputeHandler{
		Input: input, 
		Output: output,
	}
	err = handler.Compute() 
	if err != nil {
		fmt.Println(err)
	}
}



