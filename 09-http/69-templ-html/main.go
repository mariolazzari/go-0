package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title  string
	Header string
	Body   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := &Page{
			Title:  "Benvenuto",
			Header: "Ciao, mondo!",
			Body:   "Benvenuto nella mia applicazione!",
		}

		t, _ := template.ParseFiles("layout.html", "content.html")

		t.ExecuteTemplate(w, "content", p)
	})

	http.ListenAndServe(":8080", nil)
}
