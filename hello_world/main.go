package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// http.HandleFunc("/", redirect)
	// http.ListenAndServe(":8080", nil)

	http.Handle("/build/web/", http.StripPrefix("/build/web/", http.FileServer(http.Dir("build/web"))))
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "index.html", 301)
}
