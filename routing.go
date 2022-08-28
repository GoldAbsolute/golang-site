package main

import (
	"html/template"
	"net/http"
)

func testRoute(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/template.html", "tmpl/header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
