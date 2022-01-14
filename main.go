package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// Functionallity
// 1.- Read a file and extract the contents.
// 2.- Match the content.
// 3.- Transform the content to HTML tags.
// 4.- Append all the HTML tags to an HTML file.

func Convert(filename *os.File) {
	// Start a new scanner to read
	// the markdown file.
	scanner := bufio.NewScanner(filename)

	for scanner.Scan() {
		var builder strings.Builder

		// Save the scanned content.
		content := scanner.Text()

		// Check if there is text that
		// contains the '#' symbol.
		if strings.Contains(content, "#") {
			line := bytes.TrimSpace(scanner.Bytes())

			if line[0] == '#' {
				count := bytes.Count(line, []byte(`#`))
				switch count {
				case 1:
					str := strings.Replace(content, "#", "", -1)

					builder.WriteString("<h1>")
					builder.WriteString(str)
					builder.WriteString("</h1>")
				case 2:
					str := strings.Replace(content, "##", "", -1)

					builder.WriteString("<h2>")
					builder.WriteString(str)
					builder.WriteString("</h2>")
				case 3:
					str := strings.Replace(content, "###", "", -1)

					builder.WriteString("<h3>")
					builder.WriteString(str)
					builder.WriteString("</h3>")
				}
			}

			// Output.
			fmt.Println(builder.String())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("example.md")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	Convert(file)
}
