package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) == 0 {
		path = "index.html"
	}
	if !strings.Contains(path, ".") {
		path = path + ".html"
	}
	p, err := loadPage(path)
	if err != nil {
		http.Redirect(w, r, "/edit/"+path, http.StatusFound)
	}
	fmt.Fprintf(w, "%s", p.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) save() error {
	filename := p.Title
	return ioutil.WriteFile(filename, p.Body, 0600)
}

type Page struct {
	Title string
	Body  []byte
}