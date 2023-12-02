package main

import (
	"fmt"
	"log"
	"net/http"
)

// HELLO HANDLER
func helloHandler(w http.ResponseWriter, r *http.Request) {

	/*Checking if the URL path is something else other than "/hello"*/
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Illegal Method", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "HELLO BAI! BALO ASAN?")

}

// FORM HANDLER
func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "NAME: %v\n", name)
	fmt.Fprintf(w, "ADDRESS: %v\n", address)
}

func main() {

	/*Telling the fileserver to check the static folder and it
	will automatically look for the index.html file*/
	fileServer := http.FileServer(http.Dir("./static"))

	/*Setting the route for "/" when "url/" is entered*/
	http.Handle("/", fileServer)

	/*Setting the route for "/form"*/
	http.HandleFunc("/form", formHandler)

	/*Setting the route for "/hello"*/
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting at port 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
