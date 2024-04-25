package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/russross/blackfriday/v2"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input Markdown file")
	outputFile := flag.String("output", "", "Output HTML file")
	flag.Parse()

	// Check if input file is provided
	if *inputFile == "" {
		fmt.Println("Please provide an input Markdown file using the -input flag.")
		return
	}

	// Read the input Markdown file
	inputContent, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		return
	}

	// Convert Markdown to HTML
	htmlContent := blackfriday.Run(inputContent)

	// Determine the output file
	output := *outputFile
	if output == "" {
		// If output file is not provided, use the input file name with .html extension
		output = *inputFile + ".html"
	}

	// Write HTML content to the output file
	err = os.WriteFile(output, htmlContent, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %s\n", err)
		return
	}

	fmt.Printf("Conversion successful. HTML file saved at: %s\n", output)
}

// go run main.go -input README.md -output output.html
