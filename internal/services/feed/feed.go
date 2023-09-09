package feed

import (
	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mmcdole/gofeed"
)

// GetEnglishSentences returns the first four English sentences from the RSS feed
func GetEnglishSentences(rssURL string) (*dto.FeedResult, error) {
	// analyze the RSS feed
	parser := gofeed.NewParser()

	// parse the RSS feed
	feed, err := parser.ParseURL(rssURL)
	if err != nil {
		return nil, err
	}

	var englishSentences []string
	for _, item := range feed.Items {
		if isEnglishSentence(item.Title) {
			englishSentences = append(englishSentences, item.Title)
			if len(englishSentences) == 4 {
				break
			}
		}
	}

	result := &dto.FeedResult{Sentences: englishSentences}
	return result, nil
}

// isEnglishSentence verifies if a sentence is in English
func isEnglishSentence(sentence string) bool {
	for _, char := range sentence {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		}
	}
	return false
}
