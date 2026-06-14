package scheduler

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/siddharth11-sp/news-service/internal/entity"
	"github.com/siddharth11-sp/news-service/internal/ingestion"
)

type Scheduler struct {
	entityRepo       entity.Repository
	ingestionService *ingestion.Service
}

func NewScheduler(
	entityRepo entity.Repository,
	ingestionService *ingestion.Service,
) *Scheduler {

	return &Scheduler{
		entityRepo:       entityRepo,
		ingestionService: ingestionService,
	}
}

func (s *Scheduler) CheckEntities() {

	entities, err :=
		s.entityRepo.ListActiveEntities()

	if err != nil {

		log.Printf(
			"failed to load entities: %v",
			err,
		)

		return
	}

	now := time.Now().UTC()

	for _, entity := range entities {

		if shouldIngest(
			entity,
			now,
		) {

			log.Printf(
				"triggering ingestion for %s",
				entity.Name,
			)

			go func(entityID uuid.UUID) {

				if err :=
					s.ingestionService.
						IngestEntity(
							entityID,
						); err != nil {

					log.Printf(
						"ingestion failed: %v",
						err,
					)
				}

			}(entity.ID)
		}
	}
}

func shouldIngest(
	entity entity.Entity,
	now time.Time,
) bool {

	if entity.LastIngestedAt == nil {
		return true
	}

	nextRun :=
		entity.LastIngestedAt.Add(
			time.Duration(
				entity.IngestionIntervalMinutes,
			) * time.Minute,
		)

	return now.After(nextRun)
}

func (s *Scheduler) Start() {

	c := cron.New()

	_, err := c.AddFunc(
		"* * * * *",
		s.CheckEntities,
	)

	if err != nil {
		panic(err)
	}

	c.Start()

	log.Println(
		"background scheduler started",
	)
}
