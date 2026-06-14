package ingestion

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/siddharth11-sp/news-service/internal/entity"
	dedupe "github.com/siddharth11-sp/news-service/internal/ingestion/dedupe"
	"github.com/siddharth11-sp/news-service/internal/ingestion/rss"
	"github.com/siddharth11-sp/news-service/internal/ingestion/sentiment"
	"github.com/siddharth11-sp/news-service/internal/news"
)

type Service struct {
	entityRepo entity.Repository
	newsRepo   news.Repository
	fetcher    rss.Fetcher
	analyzer   sentiment.Analyzer
}

func NewService(
	entityRepo entity.Repository,
	newsRepo news.Repository,
	fetcher rss.Fetcher,
	analyzer sentiment.Analyzer,
) *Service {

	return &Service{
		entityRepo: entityRepo,
		newsRepo:   newsRepo,
		fetcher:    fetcher,
		analyzer:   analyzer,
	}
}

func (s *Service) IngestEntity(
	entityID uuid.UUID,
) error {

	entityObj, err := s.entityRepo.GetEntity(entityID)
	if err != nil {
		return err
	}

	log.Printf(
		"Starting ingestion for entity=%s",
		entityObj.Name,
	)

	articles, err := s.fetcher.Fetch(*entityObj)
	if err != nil {
		return err
	}

	log.Printf(
		"Fetched %d articles",
		len(articles),
	)

	savedCount := 0

	for _, article := range articles {

		//-----------------------------------
		// Global URL Dedupe
		//-----------------------------------

		urlHash :=
			dedupe.URLHash(
				article.URL,
			)

		exists, err :=
			s.newsRepo.
				ExistsByURLHash(
					urlHash,
				)

		if err != nil {
			return err
		}

		if exists {

			log.Printf(
				"Skipping duplicate URL=%s",
				article.URL,
			)

			continue
		}

		//-----------------------------------
		// Entity Dedupe
		//-----------------------------------

		dedupeKey :=
			dedupe.EntityDedupeKey(
				article.Source,
				article.Title,
				article.PublishedDate.Format(
					time.RFC3339,
				),
			)

		exists, err =
			s.newsRepo.
				ExistsByEntityDedupeKey(
					entityObj.ID,
					dedupeKey,
				)

		if err != nil {
			return err
		}

		if exists {

			log.Printf(
				"Skipping entity duplicate=%s",
				article.Title,
			)

			continue
		}

		//-----------------------------------
		// Sentiment
		//-----------------------------------

		sentimentResult,
			err :=
			s.analyzer.Analyze(
				article.Title,
				article.Description,
			)

		if err != nil {

			log.Printf(
				"Sentiment failed: %v",
				err,
			)

			sentimentResult = "NEUTRAL"
			// sentimentScore = 0.50
		}

		//-----------------------------------
		// Save
		//-----------------------------------

		newsArticle :=
			&news.NewsArticle{
				ID: uuid.New(),

				EntityID: entityObj.ID,

				Title: article.Title,

				Description: article.Description,

				URL: article.URL,

				URLHash: urlHash,

				DedupeKey: dedupeKey,

				Source: article.Source,

				PublishedDate: article.PublishedDate,

				Sentiment: sentimentResult,

				// SentimentScore: sentimentScore,
			}

		if err :=
			s.newsRepo.Create(
				newsArticle,
			); err != nil {

			return err
		}

		savedCount++

		log.Printf(
			"Saved article=%s",
			article.Title,
		)
	}

	//-----------------------------------
	// Update Last Ingested
	//-----------------------------------

	now := time.Now().UTC()

	if err :=
		s.entityRepo.
			UpdateLastIngestedAt(
				entityObj.ID,
				now,
			); err != nil {

		return err
	}

	log.Printf(
		"Ingestion completed entity=%s saved=%d",
		entityObj.Name,
		savedCount,
	)

	return nil
}
