package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func readFile(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("File reading error -> ", err)
	}

	return strings.Split(string(content), "\n")
}

func readTemplate(filePath string) *template.Template {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	template, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	return template
}

func prepUrlsForTemplate(urls []string) map[string]string {
	var mappedWebsites = make(map[string]string)

	for index := 0; index < len(urls); index += 3 {
		mappedWebsites[urls[index]] = urls[index+1]
	}

	return mappedWebsites
}
