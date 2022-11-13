package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	// if path doesnt match '/', return 404 error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox!"))
}

// view snippet
func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

// create snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// if method isnt POST, 405 error
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}