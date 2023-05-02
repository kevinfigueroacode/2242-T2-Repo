// in class lesson example
package main

import (
	"log"
	"net/http"
)

// Define middleware functions
func middlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// executed before the handler
		log.Println("Executing middleware A")
		next.ServeHTTP(w, r)
		// executed after the handler
		log.Println("Executing middleware A again")
	})
}

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// executed before the handler
		log.Println("Executing middleware B")
		if r.URL.Path == "/banana" {
			http.Error(w, "Access denied to /banana", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
		// executed after the handler
		log.Println("Executing middleware B again")
	})
}

func middlewareC(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// executed before the handler
		log.Println("Executing middleware C")
		// Add a custom header
		w.Header().Set("X-Server", "MyServer")
		next.ServeHTTP(w, r)
		// executed after the handler
		log.Println("Executing middleware C again")
	})
}

// Define the final handler function
func finalHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing final handler")
	w.Write([]byte("FRUITS\n"))
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Attach middleware to the final handler
	mux.Handle("/", middlewareA(middlewareB(middlewareC(http.HandlerFunc(finalHandler)))))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
