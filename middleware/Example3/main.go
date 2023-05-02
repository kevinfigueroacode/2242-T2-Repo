// Using Third-Party Middleware
// Authentication
package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	// Define the username and password for the user that can authenticate
	username := "Jane"
	password := "P@ssw0rd"

	// Create an authentication handler that uses basic authentication
	authHandler := httpauth.SimpleBasicAuth(username, password)

	// Create a final handler that just returns "OK"
	finalHandler := http.HandlerFunc(final)

	// Combine the authentication and final handlers
	http.Handle("/", authHandler(finalHandler))

	// Start the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK: Logged In"))
}
