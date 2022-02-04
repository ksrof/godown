package godown

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Read reads a given file and returns its contents.
func Read(filename string) {
	// Open the file.
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	defer file.Close()

	// Start a new scanner to read the .md file.
	scanner := bufio.NewScanner(file)

	// Scan the content line by line.
	for scanner.Scan() {
		// Save the current line.
		content := scanner.Text()

		// Check if the content contains a hash sign.
		if strings.Contains(content, "#") {
			// Convert the contents of the file to HTML.
			Convert(content, scanner)
		}
	}

	// Check if the scanner errored.
	err = scanner.Err()
	if err != nil {
		log.Fatalf("unable to scan the file: %v", err)
	}
}

// Convert replaces the hash sign with an HTML tag.
func Convert(content string, scanner *bufio.Scanner) {
	// Builder to write a new string.
	var builder strings.Builder

	// Remove whitespaces.
	line := bytes.TrimSpace(scanner.Bytes())

	if line[0] == '#' {
		// Count how many hash signs there are.
		count := bytes.Count(line, []byte("#"))

		// Depending on how many hash signs there are
		// a different html tag will be generated.
		switch count {
		case 1:
			// Remove the hash sign.
			str := strings.Replace(content, "#", "", -1)

			// Build a new string with the resulting content
			// surrounded by <h1> html tags.
			builder.WriteString(fmt.Sprintf("<h1>%s</h1>", strings.TrimSpace(str)))
		}

		// Output the resutling string.
		fmt.Println(builder.String())
	}
}
