package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog/", http.StatusFound)
}

func Router() *mux.Router {
	r := mux.NewRouter()
	// Handle duplicate content, for SEO
	r.Path("/blog").HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			http.Redirect(writer, req, "/blog/", 301)
		},
	)
	r.StrictSlash(true)

	// Blog
	r.HandleFunc("/", redirect)
	s := http.StripPrefix("/blog", http.FileServer(http.Dir("hugo/public")))
	r.PathPrefix("/blog").Handler(s)

	return r
}

func main() {
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, Router()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
