package main

import "testing"

func TestTranslateTextFunction(t *testing.T) {
	textToTranslate := "Hello"

	translatedText, err := translateText(textToTranslate)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expectedTranslation := "Olá"

	if translatedText != expectedTranslation {
		t.Errorf("Expected translation %q, but got %q", expectedTranslation, translatedText)
	}
}
