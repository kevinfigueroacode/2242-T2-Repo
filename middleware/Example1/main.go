//Illustrating the Flow of Control Example 	1
//Kevin Figueroa
package main

import (
	"fmt"
	"log"
	"net/http"
)
//creating middleware 1
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Starting middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Ending middlewareOne")
	})
}
//creating middleware 2
func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Starting middlewareTwo")
		if r.URL.Path == "/bar" {//condition if it has extension /bar
			fmt.Fprint(w, "Got /bar, exiting early")
			return
		}
		next.ServeHTTP(w, r)
		log.Print("Ending middlewareTwo")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling final request")//outputs body 
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Print("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
