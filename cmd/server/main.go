package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/siddharth11-sp/news-service/internal/config"
	"github.com/siddharth11-sp/news-service/internal/database"
	"github.com/siddharth11-sp/news-service/internal/entity"
	"github.com/siddharth11-sp/news-service/internal/ingestion"
	"github.com/siddharth11-sp/news-service/internal/ingestion/rss"
	"github.com/siddharth11-sp/news-service/internal/ingestion/sentiment"
	"github.com/siddharth11-sp/news-service/internal/news"
	"github.com/siddharth11-sp/news-service/internal/scheduler"
)

func main() {

	cfg := config.Load()

	db, err := database.NewPostgres(cfg)

	if err != nil {
		log.Fatal(err)
	}

	_ = db

	log.Println("application started")

	repo := entity.NewRepository(db)

	service := entity.NewService(repo)

	handler := entity.NewHandler(service)

	rssFetcher := rss.NewRSSFetcher()
	entityRepo := entity.NewRepository(db)
	newsRepo := news.NewRepository(db)

	analyzer := sentiment.NewOllamaAnalyzer()

	ingestionService := ingestion.NewService(
		entityRepo,
		newsRepo,
		rssFetcher,
		analyzer,
	)

	ingestionHandler :=
		ingestion.NewHandler(
			ingestionService,
		)

	scheduler := scheduler.NewScheduler(entityRepo, ingestionService)

	scheduler.Start()

	newsService :=
		news.NewService(
			newsRepo,
		)

	newsHandler :=
		news.NewHandler(
			newsService,
		)

	router := gin.Default()

	router.Use(
		cors.Default(),
	)

	router.POST(
		"/entities",
		handler.CreateEntity,
	)

	router.GET(
		"/entities",
		handler.ListEntities,
	)

	router.GET(
		"/entities/:id",
		handler.GetEntity,
	)

	// temproary endpoint to trigger ingestion for an entity
	router.POST(
		"/entities/:id/fetch",
		func(c *gin.Context) {

			id := c.Param("id")

			entityID, err := uuid.Parse(id)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "invalid entity id",
				})
				return
			}

			if err := ingestionService.IngestEntity(entityID); err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "fetch complete",
			})
		},
	)

	// news endpoints
	router.GET(
		"/entities/:id/news",
		newsHandler.ListNews,
	)

	// ingestion endpoints
	router.POST(
		"/entities/:id/refresh",
		ingestionHandler.Refresh,
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
