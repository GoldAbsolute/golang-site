package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$POST must be set")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, "Heh")
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
