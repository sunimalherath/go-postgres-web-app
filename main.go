package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tpl *template.Template
var db *sql.DB

type Book struct {
	Isbn  string
	Title string
	Price float32
}

func init() {
	var err error
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	db, err = sql.Open("postgres", "postgres://johndoe:pass123@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	// display all books in books.gohtml
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {
		book := Book{}
		rows.Scan(&book.Isbn, &book.Title, &book.Price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// execute template books.gohtml
	tpl.ExecuteTemplate(w, "books.gohtml", books)
}
