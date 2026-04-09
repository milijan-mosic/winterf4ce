package main

import (
	"fmt"
	"log"
	"net/http"
)

type Link struct {
	Name string
	URL  string
}

type Section struct {
	Title string
	Links []Link
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := "data"
	indexTemplate := readTemplate("templates/index.html")

	files := []struct {
		Title    string
		Filename string
	}{
		{Title: "Frontend", Filename: "frontend.txt"},
		{Title: "Documentation", Filename: "documentation.txt"},
		{Title: "Backend", Filename: "backend.txt"},
		{Title: "Job", Filename: "job.txt"},
		{Title: "Useful", Filename: "useful.txt"},
		{Title: "Social", Filename: "social.txt"},
	}

	sections := make([]Section, 0, len(files))

	for _, file := range files {
		urls := readFile(path + "/" + file.Filename)

		sections = append(sections, Section{
			Title: file.Title,
			Links: prepUrlsForTemplate(urls),
		})
	}

	if err := indexTemplate.Execute(w, sections); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/robots.txt")
}

func humansHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/humans.txt")
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/sitemap.xml")
}

func main() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/robots.txt", robotsHandler)
	router.HandleFunc("/humans.txt", humansHandler)
	router.HandleFunc("/sitemap.xml", sitemapHandler)

	fmt.Println("Server is running on port 20000...")
	err := http.ListenAndServe(":20000", router)
	if err != nil {
		log.Fatal("Server error -> ", err)
	}
}
