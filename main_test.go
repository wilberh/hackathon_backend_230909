package main

import (
	"testing"

	"github.com/mbuchoff/hackathon_backend_230909/translate"
)

func TestTranslateTextFunction(t *testing.T) {
	textToTranslate := "Hello"

	translatedText, err := translate.TranslateText(textToTranslate)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expectedTranslation := "Olá"

	if translatedText != expectedTranslation {
		t.Errorf("Expected translation %q, but got %q", expectedTranslation, translatedText)
	}
}
