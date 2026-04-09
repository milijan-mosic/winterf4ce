package main

import (
	"html/template"
	"log"
	"os"
)

func readFile(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", filePath, err)
	}

	return splitLines(string(content))
}

func readTemplate(filePath string) *template.Template {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatalf("failed to parse template %s: %v", filePath, err)
	}

	return tmpl
}

func prepUrlsForTemplate(lines []string) []Link {
	links := make([]Link, 0)

	for i := 0; i < len(lines); i += 3 {
		if i+1 >= len(lines) {
			log.Printf("skipping incomplete record at index %d", i)
			break
		}

		name := lines[i]
		url := lines[i+1]

		if name == "" || url == "" {
			log.Printf("skipping invalid record at index %d", i)
			continue
		}

		links = append(links, Link{
			Name: name,
			URL:  url,
		})
	}

	return links
}

func splitLines(content string) []string {
	lines := make([]string, 0)
	start := 0

	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			line := content[start:i]
			if len(line) > 0 && line[len(line)-1] == '\r' {
				line = line[:len(line)-1]
			}
			lines = append(lines, line)
			start = i + 1
		}
	}

	if start <= len(content) {
		line := content[start:]
		if len(line) > 0 && line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
		lines = append(lines, line)
	}

	return lines
}
