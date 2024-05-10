package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the port on which the server should listen.
	const port = "8080"

	// Define a simple HTTP handler function.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// If an error occurs while reading the body, log the error and send a response.
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		// Be sure to close the request body when done.
		defer r.Body.Close()

		// Write request body to standard output.
		fmt.Printf("Received request with body: %s\n", string(body))

		// Write a response back to the client.
		fmt.Fprintf(w, "Request body logged successfully")
	})

	// Log that we are starting the server.
	log.Printf("Server starting on port %s\n", port)

	// Listen and serve HTTP requests on the specified port.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
