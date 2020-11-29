package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
	lab2 "github.com/gogaeva/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Name of file containing input data")
	outputFile = flag.String("o", "", "Name of file for output")
)

func main() {
	var (
		input io.Reader
		output io.Writer
		err error
	)
	flag.Parse()

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression) 
	} else if *inputFile != "" {
		input, err = os.Open(*inputFile)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			os.Exit(1)
		}
	} else {
		input = os.Stdin
		fmt.Println("Enter your expressions. Press ctrl+D to end")
	}

	if *outputFile != "" {
		output, err = os.OpenFile(*outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		output = os.Stdout
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



