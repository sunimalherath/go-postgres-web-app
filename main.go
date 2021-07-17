package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tpl *template.Template

type Book struct {
	isbn  string
	title string
	price float32
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/books", http.StatusSeeOther)
}

func booksIndex(w http.ResponseWriter, req *http.Request) {
	// execute template books.gohtml
	tpl.ExecuteTemplate(w, "books.gohtml", nil)
}
