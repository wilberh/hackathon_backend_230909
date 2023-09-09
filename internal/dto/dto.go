package dto

type Response struct {
	Message string `json:"translated_phrase"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type Payload struct {
	// TODO: Phrases to be translated
	Text string `json:"Text"`
}
