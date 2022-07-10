package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		fmt.Fprintf(w, "Error parsing form %v", err)
	}
	fmt.Fprintf(w, "Request Successful!\n")
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	// := is a short assignment operator which is used to declare and assign a variable at the same time.

	// handle the / endpoint
	http.Handle("/", fileServer)
	// handle the /hello endpoint 
	http.HandleFunc("/hello", helloHandler)
	// handle the /form endpoint
	http.HandleFunc("/form", formHandler)
	

	fmt.Println("Listening on port 8080")

	// start the server
	// ListenAndServe starts an HTTP server with a given address and handler.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
}