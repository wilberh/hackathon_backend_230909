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

type TranslationResponse struct {
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
}

const (
	apiTranslationURL = "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=en&to=pt"
	subscriptionKey   = "138aff8d95b748468016c66926bb09c7"
	location          = "eastus"
)

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
		translatedPhrase, err := translateText(textToBeTranslated.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Message: "Internal Server Error"})
			return
		}
		fmt.Println("Translated phrase:", translatedPhrase)

		// Write the response body
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: translatedPhrase})

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseError{Message: "Method not allowed"})
	}
}

func translateText(text string) (string, error) {

	payload := []struct {
		Text string `json:"Text"`
	}{
		{Text: text},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Error marshaling JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", apiTranslationURL, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("Error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	var translationResponse []TranslationResponse
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		return "", fmt.Errorf("Error unmarshaling JSON response: %v", err)
	}

	if len(translationResponse) > 0 && len(translationResponse[0].Translations) > 0 {
		fmt.Println("Translation:", translationResponse[0].Translations[0].Text)
		return translationResponse[0].Translations[0].Text, nil
	}

	return "", fmt.Errorf("Translation not found in the response")
}
