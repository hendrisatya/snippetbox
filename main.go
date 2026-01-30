package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view module"))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create module"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
