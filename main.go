package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

type test struct {
	Ap1   string
	Ap2   string
	Port1 string
}

func main() {
	// port env
	port := os.Getenv("PORT")
	var APP_IP string
	_ = APP_IP
	APP_IP = os.Getenv("APP_IP")
	var APP_IP2 string
	_ = APP_IP2
	APP_IP2 = os.Getenv("IP")
	TestData := test{
		Ap1:   APP_IP,
		Ap2:   APP_IP2,
		Port1: port,
	}
	log.Println(TestData)
	fmt.Println(TestData)

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

	err := http.ListenAndServe(":"+port, MainRouter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server has been started")
}
