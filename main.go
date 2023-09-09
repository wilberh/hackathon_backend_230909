package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	// TODO: Phrases to be translated
	Text string `json:"Text"`
}

type Response struct {
	Message string `json:"translated_phrase"`
}

type ResponseError struct {
	Message string `json:"message"`
}

const apiTranslationURL = "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=en&to=pt"
const subscriptionKey = "138aff8d95b748468016c66926bb09c7"
const location = "eastus"

func main() {
	// Start the web server using net/http

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type header
		w.Header().Set("Content-Type", "application/json")

		// Write the response body
		json.NewEncoder(w).Encode(Response{Message: "OK"})
	})

	// Post endpoint to receive the phrase to be translated
	http.HandleFunc("/question", answerQuestion)

	// Start the web server
	fmt.Println("API is running on port 9999")
	http.ListenAndServe(":9999", nil)
}

func answerQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Set the content type header
		w.Header().Set("Content-Type", "application/json")
		// allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the request body
		defer r.Body.Close() // Close the body when we're done
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Message: "Internal Server Error"})
			return
		}
		fmt.Println(string(body))
		// extract the phrase from the request body and attach it to the payload and send to transtate the text
		textToBeTranslated := Payload{Text: string(body)}

		// translate the phrase using our function translateTexts and return the translated phrase
		translatedPhrase, err := translateTexts(textToBeTranslated.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Message: "Internal Server Error"})
			return
		}

		// Write the response body
		json.NewEncoder(w).Encode(Response{Message: translatedPhrase})

		// Write the same body back as the response
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseError{Message: "Method not allowed"})
	}
}

func translateTexts(phrase string) (string, error) {
	payload := Payload{Text: "I would really like to drive your car around the block a few times."}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// create the http request to translate the phrase
	req, err := http.NewRequest("POST", apiTranslationURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// set the request headers
	req.Header.Add("Accept", "*/*")
	// req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parse the response body
	var response []Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// return the translated phrase
	return response[0].Message, nil
}
