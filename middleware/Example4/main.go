//using logginghandler middleware to record request logs
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	// Open the log file for writing
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	// Create a logging handler that combines Apache-style logging with request/response logging
	loggingHandler := handlers.CombinedLoggingHandler(logFile, http.HandlerFunc(final))

	// Create a new HTTP server with the logging handler
	server := &http.Server{
		Addr:    ":3000",
		Handler: loggingHandler,
	}

	// Start the server
	log.Print("Listening on :3000...")
	err = server.ListenAndServe()
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK: World"))
}
