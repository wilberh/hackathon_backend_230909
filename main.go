package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	// TODO: Phrases to be translated
	Phrase string `json:"phrase"`
}

type Response struct {
	TransLatedPhrase string `json:"translated_phrase"`
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
	if r.Method == "POST" {
		responseEncoded := []byte(Response{TransLatedPhrase: "Hello API"}.TransLatedPhrase)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// return the response in json format

		json.NewEncoder(w).Encode(responseEncoded)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
