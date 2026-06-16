package rss

import (
	"log"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/siddharth11-sp/news-service/internal/entity"
)

type Fetcher interface {
	Fetch(entity entity.Entity) ([]Article, error)
}

type RSSFetcher struct {
	parser *gofeed.Parser
}

func NewRSSFetcher() *RSSFetcher {
	return &RSSFetcher{
		parser: gofeed.NewParser(),
	}
}

type FeedConfig struct {
	Name string
	URL  string
}

var rssFeeds = []FeedConfig{
	// Global Business
	{
		Name: "Reuters Business",
		URL:  "https://feeds.reuters.com/reuters/businessNews",
	},
	{
		Name: "CNBC",
		URL:  "https://www.cnbc.com/id/10001147/device/rss/rss.html",
	},

	// Indian Business (HIGH PRIORITY)
	{
		Name: "Economic Times",
		URL:  "https://economictimes.indiatimes.com/rssfeedsdefault.cms",
	},
	{
		Name: "Times of India Business",
		URL:  "https://timesofindia.indiatimes.com/rssfeeds/1898055.cms",
	},
	{
		Name: "Moneycontrol",
		URL:  "https://www.moneycontrol.com/rss/business.xml",
	},
	{
		Name: "Business Standard",
		URL:  "https://www.business-standard.com/rss/home_page_top_stories.rss",
	},
	{
		Name: "Financial Express",
		URL:  "https://www.financialexpress.com/feed/",
	},

	// Tech
	{
		Name: "TechCrunch",
		URL:  "https://techcrunch.com/feed/",
	},
	{
		Name: "The Verge",
		URL:  "https://www.theverge.com/rss/index.xml",
	},
	{
		Name: "Ars Technica",
		URL:  "https://feeds.arstechnica.com/arstechnica/index",
	},
	{
		Name: "Engadget",
		URL:  "https://www.engadget.com/rss.xml",
	},

	// Apple Specific
	{
		Name: "MacRumors",
		URL:  "https://www.macrumors.com/macrumors.xml",
	},
	{
		Name: "9to5Mac",
		URL:  "https://9to5mac.com/feed/",
	},
	{
		Name: "AppleInsider",
		URL:  "https://www.appleinsider.com/rss/news/",
	},
}

func (r *RSSFetcher) Fetch(
	entity entity.Entity,
) ([]Article, error) {

	var articles []Article

	for _, feedCfg := range rssFeeds {

		log.Printf("Fetching feed: %s", feedCfg.Name)

		feed, err := r.parser.ParseURL(feedCfg.URL)

		if err != nil {
			continue
		}
		log.Printf(
			"%s returned %d items",
			feedCfg.Name,
			len(feed.Items),
		)
		filtered := r.filterArticles(
			feed.Items,
			entity,
			feedCfg.Name,
		)

		log.Printf(
			"%s matched %d articles for entity %s",
			feedCfg.Name,
			len(filtered),
			entity.Name,
		)

		articles = append(
			articles,
			filtered...,
		)
	}
	return articles, nil
}

func (r *RSSFetcher) filterArticles(
	items []*gofeed.Item,
	entity entity.Entity,
	source string,
) []Article {

	var result []Article

	searchTerms := buildSearchTerms(entity)

	for _, item := range items {

		log.Println("RSS ITEM:", item.Title)
		content := strings.ToLower(
			item.Title + " " +
				item.Description + " " +
				item.Content,
		)

		if !containsEntity(content, searchTerms) {
			continue
		}

		article := normalizeItem(item, source)

		if article.URL == "" {
			continue
		}

		result = append(result, article)
	}

	return result
}

func buildSearchTerms(
	entity entity.Entity,
) []string {

	var terms []string

	terms =
		append(
			terms,
			strings.ToLower(entity.Name),
		)

	for _, alias := range entity.Aliases {

		terms =
			append(
				terms,
				strings.ToLower(alias),
			)
	}

	return terms
}

func containsEntity(
	content string,
	terms []string,
) bool {

	for _, term := range terms {

		if strings.Contains(
			content,
			term,
		) {
			return true
		}
	}

	return false
}

// func mapItem(
// 	item *gofeed.Item,
// 	source string,
// ) Article {

// 	var published time.Time

// 	if item.PublishedParsed != nil {
// 		published = *item.PublishedParsed
// 	}

// 	return Article{
// 		Title:         item.Title,
// 		Description:   item.Description,
// 		URL:           item.Link,
// 		Source:        source,
// 		PublishedDate: published,
// 	}
// }
