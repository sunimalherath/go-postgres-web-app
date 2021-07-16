package main

import "net/http"

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
}
