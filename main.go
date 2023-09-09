package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	// TODO: Phrases to be translated
	Phrase string `json:"phrase"`
}

type Response struct {
	TransLatedPhrase string `json:"translated_phrase"`
}

type ResponseError struct {
	Message string `json:"message"`
}

func main() {
	// Start the web server using net/http

	http.HandleFunc("/", translateTexts)

	// Post endpoint to receive the phrase to be translated
	http.HandleFunc("/question", answerQuestion)

	// Start the web server
	fmt.Println("API is running on port 9999")
	http.ListenAndServe(":9999", nil)
}

// TODO: Implement the answerQuestion function to work with the /question endpoint in POST method
func answerQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Set the content type header
		w.Header().Set("Content-Type", "application/json")

		// Get the request body
		defer r.Body.Close() // Close the body when we're done
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Message: "Internal Server Error"})
			return
		}
		// Write the same body back as the response
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseError{Message: "Method not allowed"})
	}
}

func translateTexts(w http.ResponseWriter, r *http.Request) {
	// read the request body
	// decode the request body
	// translate the phrase
	// encode the response
	// write the response
	// response := Response{TransLatedPhrase: "Hello World"}
	responseEncoded := []byte(Response{TransLatedPhrase: "Hello API"}.TransLatedPhrase)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// return the response in json format

	w.Write(responseEncoded)
}
