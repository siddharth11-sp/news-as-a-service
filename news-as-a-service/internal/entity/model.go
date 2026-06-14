package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Entity struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name                     string
	Aliases                  pq.StringArray `gorm:"type:text[]"`
	Status                   string
	IngestionIntervalMinutes int
	LastIngestedAt           *time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
