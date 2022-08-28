package main

import (
	"html/template"
	"net/http"
)

func testRoute(w http.ResponseWriter, r *http.Request) {
	_ = r
	tmpl := template.Must(template.ParseFiles("tmpl/template.html", "tmpl/header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
