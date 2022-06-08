package main

import (
	"fmt"
	"html/template"
)

type Invitee struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

// template map
var templates = make(map[string]*template.Template, 3)

// responses slice
var responses = make([]*Invitee, 0, 10)

func main() {
	loadTemplates()
}

func loadTemplates() {
	// template names
	templateNames := [5]string{"welcome", "form", "sorry", "thanks", "list"}

	// parse templates
	for _, fileName := range templateNames {
		t, err := template.ParseFiles("layout.html", fileName+".html")
		if err != nil {
			fmt.Println("Unable to parse html files")
			panic(err)
		}
		templates[fileName] = t
	}
}
