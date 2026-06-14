package rss

import "time"

type Article struct {
	Title         string
	Description   string
	URL           string
	Source        string
	PublishedDate time.Time
}
