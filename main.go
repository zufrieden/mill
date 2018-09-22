package main

import (
	"fmt"
	"net/http"
  "html/template"

  "github.com/gorilla/mux"
)

func main() {

  r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})




  tmpl := template.Must(template.ParseFiles("layout.html"))

  r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    title := "Let's go!"
    tmpl.Execute(w, title)
	})


  r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))


	http.ListenAndServe(":3000", r)

}
