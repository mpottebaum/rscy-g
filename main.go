//go:generate npm run build
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

type RscyG struct {
	Name, Email, Busyness string
	Dopness               int // 0-100
}

var rscyGs = make([]*RscyG, 0, 10)

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [3]string{"welcome", "form", "list"}
	for _, name := range templateNames {
		t, err := template.ParseFiles("views/layout.html", "views/"+name+".html")
		if err == nil {
			templates[name] = t
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}
func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, rscyGs)
}

type formData struct {
	*RscyG
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			RscyG: &RscyG{}, Errors: []string{},
		})
	}
	if request.Method == http.MethodPost {
		request.ParseForm()
		rscyGData := RscyG{
			Name:     request.Form["name"][0],
			Email:    request.Form["email"][0],
			Busyness: request.Form["busyness"][0],
			Dopness:  rand.Intn(101),
		}
		rscyGs = append(rscyGs, &rscyGData)
		http.Redirect(writer, request, "/rscy", http.StatusPermanentRedirect)
	}
}

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/rscy", listHandler)
	http.HandleFunc("/rscy/new", formHandler)
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
