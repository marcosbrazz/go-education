package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Great success!"))
}

var templates *template.Template

func usingTemplate(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

type user struct {
	Name string
}

func usingTemplateWithParams(w http.ResponseWriter, r *http.Request) {
	u := user{"Marcos"}
	templates.ExecuteTemplate(w, "home2.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)
	http.HandleFunc("/template", usingTemplate)
	http.HandleFunc("/template2", usingTemplateWithParams)

	fmt.Println("Listening to 5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
