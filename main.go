package main

import (
	"fmt"
	"html/template"
	"net/http"
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

// welcome Handler
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

// list Handler
func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}

func main() {
	// load templates
	loadTemplates()

	// create and start http server
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe(":3000", nil)
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
