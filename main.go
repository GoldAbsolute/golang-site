package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// port env
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}
	if port == "" {
		log.Fatal("$POST must be set")
	}
	// main router
	MainRouter := mux.NewRouter()

	// static files start

	MainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// static files end

	// routes
	MainRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.tmpl.html", "templates/header.tmpl.html", "templates/nav.tmpl.html"))
		errT := tmpl.ExecuteTemplate(writer, "index.tmpl.html", nil)
		if errT != nil {
			panic(errT)
		}
	})

	testSubRouter := MainRouter.PathPrefix("/test").Subrouter()
	testSubRouter.HandleFunc("", testRoute)
	testSubRouter.HandleFunc("/", testRoute)

	err := http.ListenAndServe(":"+port, MainRouter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server has been started")
}
