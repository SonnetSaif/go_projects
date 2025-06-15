package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm error: %v\n", err)
			return
		}
		fmt.Fprintf(w, "Form values: %v\n", r.Form)
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Fprintf(w, "Name: %s, Email: %s\n", name, email)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/hello" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, "hello")
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
