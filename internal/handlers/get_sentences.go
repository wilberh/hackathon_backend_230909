package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/feed"
)

func GetEnglishSentencesHandler(w http.ResponseWriter, r *http.Request) {
	rssURL := "https://www.ted.com/feeds/talks.rss"

	result, err := feed.GetEnglishSentences(rssURL)
	if err != nil {
		fmt.Println("Error getting English sentences:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.FeedResult{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
