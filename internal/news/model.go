package news

import (
	"time"

	"github.com/google/uuid"
)

type NewsArticle struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EntityID uuid.UUID `gorm:"type:uuid;index"`

	Title       string
	Description string

	URL string

	URLHash string `gorm:"index"`

	DedupeKey string `gorm:"index"`

	Source string

	PublishedDate time.Time

	Sentiment      string
	SentimentScore float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
