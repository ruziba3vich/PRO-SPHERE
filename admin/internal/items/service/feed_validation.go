package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

// ValidateFeedLink validates if a provided link is a valid feed URL.
// It returns an error if the link is not valid, and a boolean indicating success or failure.
func ValidateFeedLink(ctx context.Context, link string) (bool, error) {
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new feed parser
	parser := gofeed.NewParser()
	parser.Client = client

	// Attempt to parse the feed from the link
	_, err := parser.ParseURLWithContext(link, ctx)
	if err != nil {
		return false, errors.New("invalid feed link or unable to fetch feed")
	}

	// If parsing succeeded, the link is valid
	return true, nil
}
