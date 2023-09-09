package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/handlers"
)

func main() {
	// Start the web server using net/http

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type header
		w.Header().Set("Content-Type", "application/json")

		// Write the response body
		json.NewEncoder(w).Encode(dto.Response{Message: "OK"})
	})

	// Post endpoint to receive the phrase to be translated
	http.HandleFunc("/question", handlers.AnswerQuestion)

	// Start the web server
	fmt.Println("API is running on port 9999")
	http.ListenAndServe(":9999", nil)
}
