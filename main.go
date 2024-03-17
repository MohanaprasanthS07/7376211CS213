package main

import (
	"fmt" 
	"log"
	"net/http"  						//responsible for create basic http server
)

 func main() {                             									
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
                            									
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Mohanaprasanth S")
		})

		http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "7376211CS213")
		})

		http.HandleFunc("/dep", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "CSE")
		})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {          
		log.Fatal(err)
		
	}
}