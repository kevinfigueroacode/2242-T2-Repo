//Proper Example
package main

import (
	"log"
	"net/http"
	"strings"
)
//creating handler with two arguments
func enforceContentTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "" {//condition
			http.Error(w, "Missing Content-Type header", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(contentType, "application/json") {//condition if format
			http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))//body output
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceContentTypeHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)//port assignation
	if err != nil {
		log.Fatal(err)
	}
}
//curl -i localhost:3000
//curl -i -H "Content-Type: application/xml" localhost:3000
//curl -i -H "Content-Type: application/json; charset=UTF-8" localhost:3000