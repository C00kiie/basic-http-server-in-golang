package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		// http.StatusNotFound is a fancy way of typing 404
		// alternatively stating the code can be as useful
		// TODO: http library  conventions
		http.Error(w, "method isn't supported", http.StatusNotFound)
	}

	// handle response
	fmt.Fprintf(w, "Hello back!")
}
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Err %v", err)
	}
	// handle form vals
	// note that FormValue doesn't error when it doesn't find the value in the mapped values from
	// the resposne
	name := r.FormValue("name")
	age, err := strconv.ParseInt(r.FormValue("age"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid number")
		return
	}
	resposne := fmt.Sprintf("Name is %s, age is %d", name, age)

	fmt.Fprintf(w, resposne)
}

func main() {
	filesystem := http.FileServer(http.Dir("./static"))

	// handles
	http.Handle("/", filesystem)
	http.HandleFunc("/Form", FormHandler)
	http.HandleFunc("/Hello", helloHandler)

	// start server
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		// handle error
		log.Fatal(err)
	}

}
