package rss

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func normalizeItem(
	item *gofeed.Item,
	source string,
) Article {

	description := extractDescription(item)

	var published time.Time

	if item.PublishedParsed != nil {
		published = *item.PublishedParsed
	}

	return Article{
		Title:         cleanText(item.Title),
		Description:   description,
		URL:           cleanText(item.Link),
		Source:        source,
		PublishedDate: published,
	}
}

func cleanText(s string) string {
	return strings.TrimSpace(s)
}

func extractDescription(
	item *gofeed.Item,
) string {

	if item.Description != "" {
		return strings.TrimSpace(item.Description)
	}

	if item.Content != "" {
		return strings.TrimSpace(item.Content)
	}

	return ""
}
