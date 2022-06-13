package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

type FormData struct {
	*Invitee
	Errors []string
}

// form Handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	// if the method is GET
	if r.Method == http.MethodGet {
		templates["form"].Execute(w, FormData{
			Invitee: &Invitee{},
			Errors:  []string{},
		})
	} else if r.Method == http.MethodPost {
		// parse the form from request
		r.ParseForm()

		// create new invitee
		newInvitee := Invitee{
			Name:       r.Form["name"][0],
			Email:      r.Form["email"][0],
			Phone:      r.Form["phone"][0],
			WillAttend: r.Form["willattend"][0] == "true",
		}

		// validate data
		errors := []string{}
		if strings.TrimSpace(newInvitee.Name) == "" {
			errors = append(errors, "Name can not be empty")
		}
		if strings.TrimSpace(newInvitee.Email) == "" {
			errors = append(errors, "Email can not be empty")
		}
		if strings.TrimSpace(newInvitee.Phone) == "" {
			errors = append(errors, "Phone number can not be empty")
		}

		if len(errors) > 0 {
			templates["form"].Execute(w, FormData{
				Invitee: &newInvitee,
				Errors:  errors,
			})
			return
		}

		// add new Invitee to responses
		responses = append(responses, &newInvitee)

		// execute template according attendance
		if newInvitee.WillAttend {
			templates["thanks"].Execute(w, newInvitee.Name)
		} else {
			templates["sorry"].Execute(w, newInvitee.Name)
		}
	}
}

func main() {
	// load templates
	loadTemplates()

	// create and start http server
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)
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
